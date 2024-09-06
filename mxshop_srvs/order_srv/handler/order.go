package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"mxshop_srvs/order_srv/global"
	"mxshop_srvs/order_srv/model"
	"mxshop_srvs/order_srv/proto"
	"sync"
	"time"

	"github.com/apache/rocketmq-client-go/v2/consumer"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"
)

var m sync.Mutex

type OrderServer struct {
	proto.UnimplementedOrderServer
}

// 生成订单编号
func GenerateOrderSn(userId int32) string {
	//订单号的生成规则
	/*
		年月日时分秒+用户ID+2位随机数
	*/
	now := time.Now()
	rand.Seed(now.UnixNano())
	OrderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Nanosecond(), userId, rand.Intn(90)+10)
	return OrderSn
}

func (*OrderServer) CartItemList(c context.Context, req *proto.UserInfo) (*proto.CartItemListResponse, error) {
	//获取用户的购物车列表
	var shopCarts []model.ShoppingCart
	var rsp proto.CartItemListResponse

	if result := global.DB.Where(&model.ShoppingCart{User: req.Id}).Find(&shopCarts); result.Error != nil {
		return nil, result.Error
	} else {
		rsp.Total = int32(result.RowsAffected)
	}

	for _, shopCart := range shopCarts {
		rsp.Data = append(rsp.Data, &proto.ShopCartInfoResponse{
			Id:      shopCart.ID,
			UserId:  shopCart.User,
			GoodsId: shopCart.Goods,
			Nums:    shopCart.Nums,
			Checked: shopCart.Checked,
		})
	}
	return &rsp, nil

}

// 商品添加到购物车：1.新建 2.合并
func (*OrderServer) CreateCartItem(c context.Context, req *proto.CartItemReq) (*proto.ShopCartInfoResponse, error) {
	var shopCart model.ShoppingCart
	//筛选条件为商品ID和用户ID
	if res := global.DB.Where(&model.ShoppingCart{Goods: req.GoodsId, User: req.UserId}).First(&shopCart); res.RowsAffected == 1 {
		//如果记录已经存在，则合并购物车数量
		shopCart.Nums += req.Nums
	} else {
		shopCart.User = req.UserId
		shopCart.Goods = req.GoodsId
		shopCart.Nums = req.Nums
		shopCart.Checked = false
	}
	global.DB.Save(&shopCart)
	return &proto.ShopCartInfoResponse{Id: shopCart.ID}, nil
}

func (*OrderServer) UpdateCartItem(c context.Context, req *proto.CartItemReq) (*emptypb.Empty, error) {
	var shopCart model.ShoppingCart
	//筛选条件为商品ID和用户ID
	if res := global.DB.Where("goods = ? AND user = ?", req.GoodsId, req.UserId).First(&shopCart); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "记录不存在")
	}

	shopCart.Checked = req.Checked
	if req.Nums > 0 {
		//nums默认值为0
		shopCart.Nums = req.Nums
	}
	global.DB.Save(&shopCart)

	return &emptypb.Empty{}, nil
}

func (*OrderServer) DeleteCartItem(c context.Context, req *proto.CartItemReq) (*emptypb.Empty, error) {
	if res := global.DB.Where("goods = ? AND user = ?", req.GoodsId, req.UserId).Delete(&model.ShoppingCart{}); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "购物车记录不存在")
	}
	return &emptypb.Empty{}, nil
}

func (*OrderServer) OrderList(c context.Context, req *proto.OrderFilterRequest) (*proto.OrderListRsp, error) {
	var orders []model.OrderInfo
	//是后台管理系统查询还是电商系统查询,如果userid的值为默认值0，会自动被gorm忽略
	var total int64
	global.DB.Where(&model.OrderInfo{User: req.UserId}).Count(&total)

	rsp := &proto.OrderListRsp{
		Total: int32(total),
	}

	//分页
	global.DB.Scopes(Paginate(int(req.Page), int(req.PerNums))).Where(&model.OrderInfo{User: req.UserId}).Find(&orders)
	for _, order := range orders {
		rsp.Data = append(rsp.Data, &proto.OrderInfoRsp{
			Id:      order.ID,
			UserId:  order.User,
			Status:  order.Status,
			OrderSn: order.OrderSn,
			PayType: order.PayType,
			Total:   order.OrderMount,
			Post:    order.Post,
			Address: order.Address,
			Name:    order.SignerName,
			Mobile:  order.SignerMobile,
			AddTime: order.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return rsp, nil
}

func (*OrderServer) OrderDetail(c context.Context, req *proto.OrderReq) (*proto.OrderInfoDetalRsp, error) {
	var (
		order      model.OrderInfo
		orderGoods []model.OrderGoods
	)

	//web层应该检查当前订单是否是当前用户的，电商系统：用户ID和订单ID，后台管理系统：订单ID
	if result := global.DB.Where(&model.OrderInfo{User: req.UserId, BaseModel: model.BaseModel{ID: req.Id}}).First(&order); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "订单不存在")
	}

	global.DB.Where(&model.OrderGoods{Order: order.ID}).Find(&orderGoods)

	rsp := &proto.OrderInfoDetalRsp{
		OrderInfo: &proto.OrderInfoRsp{
			Id:      order.ID,
			UserId:  order.User,
			Status:  order.Status,
			OrderSn: order.OrderSn,
			PayType: order.PayType,
			Total:   order.OrderMount,
			Post:    order.Post,
			Address: order.Address,
			Name:    order.SignerName,
			Mobile:  order.SignerMobile,
		},
	}

	for _, orderGood := range orderGoods {
		rsp.Goods = append(rsp.Goods, &proto.OrderItemRsp{
			GoodsId:    orderGood.Goods,
			GoodName:   orderGood.GoodsName,
			GoodsPrice: orderGood.GoodsPrice,
			Nums:       orderGood.Nums,
		})
	}
	return rsp, nil
}

type OrderListener struct {
	Code        codes.Code
	Detail      string
	ID          int32
	OrderAmount float32
}

func (o *OrderListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	//执行本地事务
	//跨服务调用-商品
	var orderInfo model.OrderInfo
	_ = json.Unmarshal(msg.Body, &orderInfo)

	var goodsIds []int32
	var shopCarts []model.ShoppingCart
	goodsNumsMap := make(map[int32]int32)
	if result := global.DB.Where(&model.ShoppingCart{User: orderInfo.User, Checked: true}).Find(&shopCarts); result.RowsAffected == 0 {
		o.Code = codes.InvalidArgument
		o.Detail = "没有选中结算的商品"
		return primitive.RollbackMessageState
	}

	for _, shopCart := range shopCarts {
		goodsIds = append(goodsIds, shopCart.Goods)
		goodsNumsMap[shopCart.Goods] = shopCart.Nums
	}

	//跨服务调用商品微服务
	goods, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{Id: goodsIds})
	if err != nil {
		o.Code = codes.Internal
		o.Detail = "批量查询商品信息失败"
		return primitive.RollbackMessageState
	}

	var orderAmount float32
	var orderGoods []*model.OrderGoods
	var goodsInvInfo []*proto.GoodsInvInfo
	for _, good := range goods.Data {
		orderAmount += good.ShopPrice * float32(goodsNumsMap[good.Id])
		orderGoods = append(orderGoods, &model.OrderGoods{
			Goods:      good.Id,
			GoodsName:  good.Name,
			GoodsImage: good.GoodsFrontImage,
			GoodsPrice: good.ShopPrice,
			Nums:       goodsNumsMap[good.Id],
		})

		goodsInvInfo = append(goodsInvInfo, &proto.GoodsInvInfo{
			GoodsId: good.Id,
			Num:     goodsNumsMap[good.Id],
		})
	}

	//跨服务调用库存微服务进行库存扣减
	/*
		1. 调用库存服务的trysell
		2. 调用仓库服务的trysell
		3. 调用积分服务的tryAdd
		任何一个服务出现了异常，那么你得调用对应的所有的微服务的cancel接口
		如果所有的微服务都正常，那么你得调用所有的微服务的confirm
	*/
	if _, err = global.InventorySrvClient.Sell(context.Background(), &proto.SellInfo{OrderSn: orderInfo.OrderSn, GoodInfo: goodsInvInfo}); err != nil {
		//如果是因为网络问题， 这种如何避免误判， 大家自己改写一下sell的返回逻辑
		o.Code = codes.ResourceExhausted
		o.Detail = "扣减库存失败"
		return primitive.RollbackMessageState
	}

	//o.Code = codes.Internal
	//o.Detail = "本地事务执行失败"
	//return primitive.UnknowState
	//生成订单表
	tx := global.DB.Begin()
	orderInfo.OrderMount = orderAmount
	if result := tx.Save(&orderInfo); result.RowsAffected == 0 {
		tx.Rollback()
		o.Code = codes.Internal
		o.Detail = "创建订单失败"
		return primitive.CommitMessageState
	}

	o.OrderAmount = orderAmount
	o.ID = orderInfo.ID
	for _, orderGood := range orderGoods {
		orderGood.Order = orderInfo.ID
	}

	//批量插入orderGoods
	if result := tx.CreateInBatches(orderGoods, 100); result.RowsAffected == 0 {
		tx.Rollback()
		o.Code = codes.Internal
		o.Detail = "批量插入订单商品失败"
		return primitive.CommitMessageState
	}

	if result := tx.Where(&model.ShoppingCart{User: orderInfo.User, Checked: true}).Delete(&model.ShoppingCart{}); result.RowsAffected == 0 {
		tx.Rollback()
		o.Code = codes.Internal
		o.Detail = "删除购物车记录失败"
		return primitive.CommitMessageState
	}

	//发送延时消息
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithGroupName("hb_producer_order"))
	if err != nil {
		panic("生成producer失败")
	}

	//不要在一个进程中使用多个producer， 如果使用的话，不要随便调用shutdown因为会影响其他的producer
	if err = p.Start(); err != nil {
		panic("启动producer失败")
	}

	msg = primitive.NewMessage("order_timeout", msg.Body)
	msg.WithDelayTimeLevel(3)
	_, err = p.SendSync(context.Background(), msg)
	if err != nil {
		zap.S().Errorf("发送延时消息失败: %v\n", err)
		tx.Rollback()
		o.Code = codes.Internal
		o.Detail = "发送延时消息失败"
		return primitive.CommitMessageState
	}

	//if err = p.Shutdown(); err != nil {panic("关闭producer失败")}

	//提交事务
	tx.Commit()
	o.Code = codes.OK
	return primitive.RollbackMessageState
}

func (o *OrderListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	var orderInfo model.OrderInfo
	_ = json.Unmarshal(msg.Body, &orderInfo)
	//回查逻辑,如何检查之前的逻辑是否完成呢
	//查询此订单是否存在
	if res := global.DB.Where(model.OrderInfo{OrderSn: orderInfo.OrderSn}).First(&orderInfo); res.RowsAffected == 0 {
		//订单不存在，订单并没有创建成功
		return primitive.CommitMessageState //此处不能说明订单已扣减，需要幂等性保证
	}
	return primitive.RollbackMessageState
}

func (*OrderServer) CreateOrder(c context.Context, req *proto.OrderReq) (*proto.OrderInfoRsp, error) {
	/*新建订单
	1.商品金额：访问商品（跨服务调用）
	2.库存扣减（跨服务调用）
	3.从购物车中获取选中的商品
	4.订单基本信息表-商品信息表
	5.从购物车中删除已购买的记录
	*/
	//分布式事务
	orderListener := &OrderListener{}
	p, err := rocketmq.NewTransactionProducer(
		orderListener,
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2),
	)
	if err != nil {
		//panic(err)
		zap.S().Errorf("生成producer失败：%s", err.Error())
		return nil, err
	}
	err = p.Start()
	if err != nil {
		zap.S().Errorf("启动producer失败：%s", err.Error())
		return nil, err
	}

	order := &model.OrderInfo{
		OrderSn:      GenerateOrderSn(req.UserId),
		Address:      req.Address,
		SignerMobile: req.Mobile,
		SignerName:   req.Name,
		Post:         req.Post,
		User:         req.UserId,
	}

	jsonString, err := json.Marshal(order)

	//topic为库存归还,发送消息
	transaction, err := p.SendMessageInTransaction(context.Background(), primitive.NewMessage("inventory_return", jsonString))
	if err != nil {
		zap.S().Errorf("发送失败")
		return nil, status.Errorf(codes.Internal, "rocketmq发送消息失败")
	} else {
		fmt.Println("发送成功:%s\n", transaction.String())
	}
	if orderListener.Code != codes.OK {
		//新建订单失败，需要归还库存
		return nil, status.Errorf(orderListener.Code, orderListener.Detail)
	}
	//time.Sleep(1 * time.Hour)
	//var (
	//	shoppingCarts []model.ShoppingCart
	//	orderGoods    []*model.OrderGoods
	//	goodsInvInfo  []*proto.GoodsInvInfo
	//	goodsIds      []int32
	//	orderAmount   float32
	//)
	//goodsNumMap := make(map[int32]int32) //goodsid为key,商品数量为value
	//if res := global.DB.Where(&model.ShoppingCart{User: req.UserId, Checked: true}).Find(&shoppingCarts); res.RowsAffected == 0 {
	//	return nil, status.Errorf(codes.InvalidArgument, "没有选中结算的商品")
	//}
	//fmt.Println(shoppingCarts)

	//for _, shoppingCart := range shoppingCarts {
	//	goodsIds = append(goodsIds, shoppingCart.Goods)
	//	goodsNumMap[shoppingCart.Goods] = shoppingCart.Nums
	//}
	return &proto.OrderInfoRsp{
		Id:      orderListener.ID,
		OrderSn: order.OrderSn,
		Total:   orderListener.OrderAmount,
	}, nil
}

func (*OrderServer) UpdateOrderStatus(c context.Context, req *proto.OrderStatus) (*emptypb.Empty, error) {
	//根据订单编号来生成
	if res := global.DB.Model(&model.OrderInfo{}).Where("order_sn = ?", req.OrderSn).Update("status", req.Status); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "订单不存在")
	}
	return &emptypb.Empty{}, nil
}

func OrderTimeout(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

	for i := range msgs {
		var orderInfo model.OrderInfo
		_ = json.Unmarshal(msgs[i].Body, &orderInfo)

		fmt.Printf("获取到订单超时消息: %v\n", time.Now())
		//查询订单的支付状态，如果已支付什么都不做，如果未支付，归还库存
		var order model.OrderInfo
		if result := global.DB.Model(model.OrderInfo{}).Where(model.OrderInfo{OrderSn: orderInfo.OrderSn}).First(&order); result.RowsAffected == 0 {
			return consumer.ConsumeSuccess, nil
		}
		if order.Status != "TRADE_SUCCESS" {
			tx := global.DB.Begin()
			//订单存在且未支付完成，修改订单的状态为关闭订单
			order.Status = "TRADE_CLOSED"
			tx.Save(&order)

			//归还库存，我们可以模仿order中发送一个消息到 order_reback中去
			p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"127.0.0.1:9876"}))
			if err != nil {
				panic("生成producer失败")
			}

			if err = p.Start(); err != nil {
				panic("启动producer失败")
			}

			//此处为普通消息
			_, err = p.SendSync(context.Background(), primitive.NewMessage("inventory_return", msgs[i].Body))
			if err != nil {
				tx.Rollback()
				fmt.Printf("发送失败: %s\n", err)
				return consumer.ConsumeRetryLater, nil
			}

			//if err = p.Shutdown(); err != nil {
			//	panic("关闭producer失败")
			//}
			tx.Commit()
			return consumer.ConsumeSuccess, nil
		}
	}
	return consumer.ConsumeSuccess, nil
}
