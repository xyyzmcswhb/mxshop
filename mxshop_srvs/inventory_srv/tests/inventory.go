package main

import (
	"context"
	"fmt"
	"mxshop_srvs/inventory_srv/proto"
	"sync"
)

func TestSetInv(goodsId, nums int32) {
	_, err := InvClient.SetStocks(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
		Num:     nums,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("设置库存成功")
}

func TestInvDetail(goodsId int32) {
	/*
		1.第一件扣减成功，第二件扣减失败，事务能否回滚
		2. 两件都扣减成功
	*/
	rsp, err := InvClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Num)
}

func TestSell(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := InvClient.Sell(context.Background(), &proto.SellInfo{
		GoodInfo: []*proto.GoodsInvInfo{
			{
				GoodsId: 421,
				Num:     1,
			},
			//{
			//	GoodsId: 422,
			//	Num:     1,
			//},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("库存扣减成功")
}

func TestReback() {
	//defer wg.Done()
	_, err := InvClient.ReBack(context.Background(), &proto.SellInfo{
		GoodInfo: []*proto.GoodsInvInfo{
			{
				GoodsId: 421,
				Num:     1,
			},
			//{
			//	GoodsId: 422,
			//	Num:     30,
			//},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("库存归还成功")
}
