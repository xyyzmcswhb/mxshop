package shopcart

import (
	"context"
	"mxshop-api/order-web/api"
	"mxshop-api/order-web/forms"
	"mxshop-api/order-web/global"
	"mxshop-api/order-web/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func List(c *gin.Context) {
	//获取购物车商品
	userId, _ := c.Get("userId")
	rsp, err := global.OrderSrvClient.CartItemList(context.Background(), &proto.UserInfo{
		Id: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("查询购物车列表失败")
		api.HandleGrpcError2Http(err, c)
		return
	}

	//商品ID
	goodsids := make([]int32, 0)
	for _, item := range rsp.Data {
		goodsids = append(goodsids, item.GoodsId)
	}

	if len(goodsids) == 0 {
		c.JSON(200, gin.H{
			"msg": "没有商品",
		})
		return
	}

	//请求商品服务获取商品信息
	goodsRsp, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: goodsids,
	})
	if err != nil {
		zap.S().Errorw("批量查询商品失败")
		api.HandleGrpcError2Http(err, c)
		return
	}

	reMap := gin.H{
		"total": rsp.Total,
	}

	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		for _, good := range goodsRsp.Data {
			if good.Id == item.GoodsId {
				tmpMap := map[string]interface{}{}
				tmpMap["id"] = item.Id
				tmpMap["goods_id"] = item.GoodsId
				tmpMap["good_name"] = good.Name
				tmpMap["good_image"] = good.GoodsFrontImage
				tmpMap["good_price"] = good.ShopPrice
				tmpMap["nums"] = item.Nums
				tmpMap["checked"] = item.Checked

				goodsList = append(goodsList, tmpMap)
			}
		}
	}
	reMap["data"] = goodsList
	c.JSON(http.StatusOK, reMap)
}

func New(c *gin.Context) {
	itemForm := forms.ShopCartItemForm{}
	if err := c.ShouldBind(&itemForm); err != nil {
		api.HandleGrpcError2Http(err, c)
		return
	}

	//检查商品是否存在
	_, err := global.GoodsSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("[List] 查询【商品信息】失败")
		api.HandleGrpcError2Http(err, c)
		return
	}

	invRsp, err := global.InvSrvClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("商品库存不存在")
		api.HandleGrpcError2Http(err, c)
		return
	}
	//如果添加到购物车的数量和库存的数量不一致
	if invRsp.Num < itemForm.Nums {
		c.JSON(500, gin.H{
			"msg": "商品库存不足",
		})
		return
	}

	userId, _ := c.Get("userId")
	rsp, err := global.OrderSrvClient.CreateCartItem(context.Background(), &proto.CartItemReq{
		GoodsId: itemForm.GoodsId,
		UserId:  int32(userId.(uint)),
		Nums:    itemForm.Nums,
	})

	if err != nil {
		zap.S().Errorw("添加到购物车失败")
		api.HandleGrpcError2Http(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": rsp.Id,
	})

}

func Detail(c *gin.Context) {}

func Delete(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式出错",
		})
		return
	}

	//删除当前用户
	userId, _ := c.Get("userId")
	_, err = global.OrderSrvClient.DeleteCartItem(context.Background(), &proto.CartItemReq{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
	})

	if err != nil {
		zap.S().Errorw("删除购物车记录失败")
		api.HandleGrpcError2Http(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func Update(c *gin.Context) {
	itemForm := forms.ShopCartItemUpdateForm{}
	if err := c.ShouldBind(&itemForm); err != nil {
		api.HandleGrpcError2Http(err, c)
		return
	}

	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式出错",
		})
		return
	}

	userId, _ := c.Get("userId")
	request := proto.CartItemReq{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
		Nums:    itemForm.Nums,
		Checked: false,
	}
	if itemForm.Checked != nil {
		request.Checked = *itemForm.Checked
	}
	_, err = global.OrderSrvClient.UpdateCartItem(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("更新购物车记录出错")
		api.HandleGrpcError2Http(err, c)
		return
	}

	c.Status(http.StatusOK)
}
