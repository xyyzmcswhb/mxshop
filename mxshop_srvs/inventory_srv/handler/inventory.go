package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"mxshop_srvs/inventory_srv/global"
	"mxshop_srvs/inventory_srv/model"
	"mxshop_srvs/inventory_srv/proto"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"
)

var m sync.Mutex

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (*InventoryServer) SetStocks(c context.Context, req *proto.GoodsInvInfo) (*emptypb.Empty, error) {
	//设置库存（新增库存及更新库存
	var inv model.Inventory
	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv)
	//global.DB.First(&inv, req.GoodsId)//查询当前商品库存,当且仅当goodsid为主键时这样写正确
	if inv.Goods == 0 {
		//没查到库存
		inv.Goods = req.GoodsId
	}
	inv.Stocks = req.Num
	global.DB.Save(&inv)
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) InvDetail(c context.Context, req *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	var inv model.InventoryNew
	if res := global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv); res.RowsAffected == 0 {
		//商品不存在
		return nil, status.Error(codes.NotFound, "库存信息不存在")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num:     inv.Stocks - inv.Freeze,
	}, nil
}

func (*InventoryServer) Sell(c context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	//client := goredislib.NewClient(&goredislib.Options{
	//	Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	//})
	//pool := goredis.NewPool(client)
	//rs := redsync.New(pool)
	//扣减库存
	//需满足数据库的一致性，原子性,因此需要开启事务
	//并发情况下可能会出现超卖
	tx := global.DB.Begin()
	//m.Lock() //获取锁
	sellDetail := model.StockSellDetail{
		OrderSn: req.OrderSn,
		Status:  1, //表示已经扣减
	}

	var details []model.GoodsDetail
	for _, goods := range req.GoodInfo {
		details = append(details, model.GoodsDetail{
			Goods: goods.GoodsId,
			Num:   goods.Num,
		})
		var inv model.Inventory
		//查每个商品的库存,利用悲观锁保持并发
		//if res := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: goods.GoodsId}).First(&inv); res.RowsAffected == 0 {
		//	tx.Rollback()
		//	return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
		//}
		//for {
		mutex := global.Rs.NewMutex(fmt.Sprintf("goods_%d", goods.GoodsId)) //获取分布式锁
		if err := mutex.Lock(); err != nil {
			return nil, status.Error(codes.Internal, "获取分布式锁错误")
		}
		if res := global.DB.Where(&model.Inventory{Goods: goods.GoodsId}).First(&inv); res.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
		}
		//如果现有库存数量小于要扣减的库存数量，事务回滚
		if inv.Stocks < goods.Num {
			tx.Rollback()
			return nil, status.Error(codes.ResourceExhausted, "剩余库存不足")
		}

		//扣减，并发场景下会出现数据不一致的问题-分布式锁
		inv.Stocks -= goods.Num
		tx.Save(&inv)

		if ok, err := mutex.Unlock(); !ok || err != nil {
			return nil, status.Error(codes.Internal, "释放分布式锁异常")
		}
		//用乐观锁解决数据不一致的问题,更新多个字段要使用updates
		//零值会被gorm忽略掉，需强制更新：select
		//if res := tx.Model(&model.Inventory{}).Select("Stocks", "Version").Where("goods = ? and version = ?", goods.GoodsId, inv.Version).Updates(model.Inventory{Stocks: inv.Stocks, Version: inv.Version + 1}); res.RowsAffected == 0 {
		//	zap.S().Infof("库存扣减失败")
		//} else {
		//	break //扣减成功跳出循环
		//}
	}
	sellDetail.Detail = details
	//将这条扣减记录插入数据库
	if res := tx.Create(&sellDetail); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "库存扣减记录保存失败")
	}
	//手动提交事务，使写入的数据持久化入数据库
	tx.Commit()
	//m.Unlock() //释放锁
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) ReBack(c context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	//库存归还：1.订单超时归还 2. 订单创建失败，归还之前扣减的库存 3.手动归还
	tx := global.DB.Begin()
	m.Lock() //获取锁
	for _, goods := range req.GoodInfo {
		var inv model.Inventory
		//查每个商品的库存
		if res := global.DB.Where(&model.Inventory{Goods: goods.GoodsId}).First(&inv); res.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
		}

		//归还库存，并发场景下会出现数据不一致的问题-分布式锁
		inv.Stocks += goods.Num
		tx.Save(&inv)
	}
	//手动提交事务，使写入的数据持久化入数据库
	tx.Commit()
	m.Unlock() //释放锁
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) TrySell(c context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	tx := global.DB.Begin()
	//m.Lock() //获取锁
	for _, goods := range req.GoodInfo {
		var inv model.InventoryNew
		//查每个商品的库存,利用悲观锁保持并发
		//if res := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: goods.GoodsId}).First(&inv); res.RowsAffected == 0 {
		//	tx.Rollback()
		//	return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
		//}
		//for {
		mutex := global.Rs.NewMutex(fmt.Sprintf("goods_%d", goods.GoodsId)) //获取锁
		if err := mutex.Lock(); err != nil {
			return nil, status.Error(codes.Internal, "获取分布式锁错误")
		}
		if res := global.DB.Where(&model.Inventory{Goods: goods.GoodsId}).First(&inv); res.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
		}
		//如果现有库存数量小于要扣减的库存数量，事务回滚
		if inv.Stocks < goods.Num {
			tx.Rollback()
			return nil, status.Error(codes.ResourceExhausted, "剩余库存不足")
		}

		//扣减，并发场景下会出现数据不一致的问题-分布式锁
		//inv.Stocks -= goods.Num
		inv.Freeze += goods.Num
		tx.Save(&inv)
		if ok, err := mutex.Unlock(); !ok || err != nil {
			return nil, status.Error(codes.Internal, "释放分布式锁异常")
		}
		//用乐观锁解决数据不一致的问题,更新多个字段要使用updates
		//零值会被gorm忽略掉，需强制更新：select
		//if res := tx.Model(&model.Inventory{}).Select("Stocks", "Version").Where("goods = ? and version = ?", goods.GoodsId, inv.Version).Updates(model.Inventory{Stocks: inv.Stocks, Version: inv.Version + 1}); res.RowsAffected == 0 {
		//	zap.S().Infof("库存扣减失败")
		//} else {
		//	break //扣减成功跳出循环
		//}
	}

	//}
	//手动提交事务，使写入的数据持久化入数据库
	tx.Commit()
	//m.Unlock() //释放锁
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) ConfirmSell(c context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	tx := global.DB.Begin()
	//m.Lock() //获取锁
	for _, goods := range req.GoodInfo {
		var inv model.InventoryNew
		mutex := global.Rs.NewMutex(fmt.Sprintf("goods_%d", goods.GoodsId)) //获取锁
		if err := mutex.Lock(); err != nil {
			return nil, status.Error(codes.Internal, "获取分布式锁错误")
		}
		if res := global.DB.Where(&model.Inventory{Goods: goods.GoodsId}).First(&inv); res.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
		}
		//如果现有库存数量小于要扣减的库存数量，事务回滚
		if inv.Stocks < goods.Num {
			tx.Rollback()
			return nil, status.Error(codes.ResourceExhausted, "剩余库存不足")
		}

		//扣减，并发场景下会出现数据不一致的问题-分布式锁
		inv.Stocks -= goods.Num
		inv.Freeze -= goods.Num //确认扣减时，预扣减时加上的数量也得减掉

		tx.Save(&inv)
		if ok, err := mutex.Unlock(); !ok || err != nil {
			return nil, status.Error(codes.Internal, "释放分布式锁异常")
		}
	}

	//}
	//手动提交事务，使写入的数据持久化入数据库
	tx.Commit()
	//m.Unlock() //释放锁
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) CancelSell(c context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	tx := global.DB.Begin()
	//m.Lock() //获取锁
	for _, goods := range req.GoodInfo {
		var inv model.InventoryNew
		mutex := global.Rs.NewMutex(fmt.Sprintf("goods_%d", goods.GoodsId)) //获取锁
		if err := mutex.Lock(); err != nil {
			return nil, status.Error(codes.Internal, "获取分布式锁错误")
		}
		if res := global.DB.Where(&model.Inventory{Goods: goods.GoodsId}).First(&inv); res.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Error(codes.InvalidArgument, "库存信息不存在")
		}
		//如果现有库存数量小于要扣减的库存数量，事务回滚
		if inv.Stocks < goods.Num {
			tx.Rollback()
			return nil, status.Error(codes.ResourceExhausted, "剩余库存不足")
		}

		//扣减，并发场景下会出现数据不一致的问题-分布式锁
		inv.Freeze -= goods.Num // 取消扣减时，预扣减时的冻结库存数量要减掉

		tx.Save(&inv)
		if ok, err := mutex.Unlock(); !ok || err != nil {
			return nil, status.Error(codes.Internal, "释放分布式锁异常")
		}
	}

	//}
	//手动提交事务，使写入的数据持久化入数据库
	tx.Commit()
	//m.Unlock() //释放锁
	return &emptypb.Empty{}, nil
}

func AutoInvReturn(c context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	//响应逻辑，类似于回调
	//归还库存，应具体知道是哪个订单，重复归还？需确保幂等性，不能因为消息的重复发送导致一个订单的库存归还多次，没有扣减的库存不能归还
	//解决方案：新建一张表记录详细订单库存扣减细节以及归还细节
	type OrderInfo struct {
		OrderSn string
	}
	for i := range msgs {
		var orderInfo OrderInfo
		if err := json.Unmarshal(msgs[i].Body, &orderInfo); err != nil {
			zap.S().Errorf("Unmarshal err:%v\n", msgs[i].Body)
			return consumer.ConsumeSuccess, nil //丢弃消息
		}

		//将inv的库存加回去，将selldetail status设置为2
		tx := global.DB.Begin()
		var sellDetail model.StockSellDetail
		//查询已经扣减但是还未归还的数据
		if res := tx.Model(&model.StockSellDetail{}).Where(&model.StockSellDetail{OrderSn: orderInfo.OrderSn, Status: 1}).First(&sellDetail); res.RowsAffected == 0 {
			tx.Rollback()
			return consumer.ConsumeRetryLater, nil
		}

		//如果查询到，那么逐个归还库存
		for _, orderGood := range sellDetail.Detail {
			//update，归还库存
			if res := tx.Model(&model.Inventory{}).Where(&model.Inventory{Goods: orderGood.Goods}).Update("stocks", gorm.Expr("stocks+?", orderGood.Num)); res.RowsAffected == 0 {
				//更新失败
				tx.Rollback()
				return consumer.ConsumeRetryLater, nil
			}
		}

		if res := tx.Model(&model.StockSellDetail{}).Where(&model.StockSellDetail{OrderSn: orderInfo.OrderSn}).Update("status", 2); res.RowsAffected == 0 {
			tx.Rollback()
			return consumer.ConsumeSuccess, nil
		}
		tx.Commit()
		return consumer.ConsumeSuccess, nil
	}
	//直接返回success,此条消息被标记为已经消费过，不会再被消费
	return consumer.ConsumeSuccess, nil
}
