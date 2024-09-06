package goods

import (
	"context"
	"mxshop-api/goods-web/api"
	"mxshop-api/goods-web/forms"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	//商品的列表,主要业务逻辑为数据过滤
	request := &proto.GoodsFilterRequest{}
	priceMin := c.DefaultQuery("pmin", "0")
	priceMinInt, _ := strconv.Atoi(priceMin)
	priceMax := c.DefaultQuery("pmax", "1000000")
	priceMaxInt, _ := strconv.Atoi(priceMax)
	request.PriceMin = int32(priceMinInt)
	request.PriceMax = int32(priceMaxInt)

	isHot := c.DefaultQuery("ih", "0")
	if isHot == "1" {
		request.IsHot = true
	}

	isNew := c.DefaultQuery("in", "0")
	if isNew == "1" {
		request.IsNew = true
	}

	isTab := c.DefaultQuery("it", "0")
	if isTab == "1" {
		request.IsTab = true
	}

	categoryId := c.DefaultQuery("category", "0")
	categoryIdInt, _ := strconv.Atoi(categoryId)
	request.TopCategory = int32(categoryIdInt)

	pages := c.DefaultQuery("pages", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := c.DefaultQuery("pernum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)

	keywords := c.DefaultQuery("keywords", "")
	request.KeyWords = keywords

	brand := c.DefaultQuery("brand", "")
	brandInt, _ := strconv.Atoi(brand)
	request.Brand = int32(brandInt)

	//请求商品的service服务
	Rsp, err := global.GoodsSrvClient.GoodsList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("GetGoodsList fail", "查询商品列表失败", err)
		api.HandleGrpcError2Http(err, c)
		return
	}

	goodsList := make([]interface{}, 0)
	for _, value := range Rsp.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"category": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"is_hot":  value.IsHot,
			"is_new":  value.IsNew,
			"on_sale": value.OnSale,
		})
	}

	reMap := map[string]interface{}{
		"total": Rsp.Total,
		"data":  goodsList,
	}
	c.JSON(http.StatusOK, reMap)
}

func NewGoods(c *gin.Context) {
	goodsForm := forms.GoodsForm{}
	if err := c.ShouldBind(&goodsForm); err != nil {
		api.HandleGrpcError2Http(err, c)
		return
	}

	goodsClient := global.GoodsSrvClient
	rsp, err := goodsClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		Brand:           goodsForm.Brand,
	})
	if err != nil {
		api.HandleGrpcError2Http(err, c)
		return
	}
	//如何设置库存
	//TODO
	c.JSON(http.StatusOK, rsp)
}

func Detail(c *gin.Context) {
	//通过URL去获取ID
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	rsp, err := global.GoodsSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: int32(i),
	})
	if err != nil {
		api.HandleGrpcError2Http(err, c)
		return
	}

	//TODO 去库存服务查询库存
	detailRsp := map[string]interface{}{
		"id":          rsp.Id,
		"name":        rsp.Name,
		"goods_brief": rsp.GoodsBrief,
		"desc":        rsp.GoodsDesc,
		"shop_price":  rsp.ShopPrice,
		"images":      rsp.Images,
		"desc_images": rsp.DescImages,
		"ship_free":   rsp.ShipFree,
		"front_image": rsp.GoodsFrontImage,
		"category": map[string]interface{}{
			"id":   rsp.Category.Id,
			"name": rsp.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   rsp.Brand.Id,
			"name": rsp.Brand.Name,
			"logo": rsp.Brand.Logo,
		},
		"is_hot":  rsp.IsHot,
		"is_new":  rsp.IsNew,
		"on_sale": rsp.OnSale,
	}
	c.JSON(http.StatusOK, detailRsp)
}

func DeleteGoods(c *gin.Context) {
	id := c.Param("id")
	i, er := strconv.ParseInt(id, 10, 32)
	if er != nil {
		c.Status(http.StatusNotFound)
		return
	}
	_, err := global.GoodsSrvClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{
		Id: int32(i),
	})
	if err != nil {
		api.HandleGrpcError2Http(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

// 获取商品库存
func Stocks(c *gin.Context) {
	id := c.Param("id")
	_, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	//TODO 商品库存
	return
}

func UpdateStatus(c *gin.Context) {
	goodsStatusForm := forms.GoodsStatusForm{}
	if err := c.ShouldBind(&goodsStatusForm); err != nil {
		api.HandleGrpcError2Http(err, c)
		return
	}

	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	_, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:     int32(i),
		IsHot:  *goodsStatusForm.IsHot,
		IsNew:  *goodsStatusForm.IsNew,
		OnSale: *goodsStatusForm.OnSale,
	})
	if err != nil {
		api.HandleGrpcError2Http(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新状态成功",
	})
}

func UpdateGoods(c *gin.Context) {
	goodsForm := forms.GoodsForm{}
	if err := c.ShouldBind(&goodsForm); err != nil {
		api.HandleGrpcError2Http(err, c)
		return
	}

	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	_, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:              int32(i),
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		Brand:           goodsForm.Brand,
	})
	if err != nil {
		api.HandleGrpcError2Http(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
