package router

import (
	"mxshop-api/goods-web/api/brands"
	"mxshop-api/goods-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitBrandsRouter(Router *gin.RouterGroup) {
	//路由组
	brandsRouter := Router.Group("brands")
	{
		brandsRouter.GET("list", brands.BrandList)                                                            //品牌列表
		brandsRouter.POST("add_new_brand", middlewares.JWTAuth(), middlewares.IsadminAuth(), brands.NewBrand) //此接口需要管理员权限，新增品牌
		brandsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), brands.DeleteBrand)     //删除品牌
		brandsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), brands.UpdateBrand)        //更新品牌
	}

	CategoryBrandRouter := Router.Group("categorybrand")
	{
		CategoryBrandRouter.GET("list", brands.CategoryBrandList)                                                                    //分类品牌列表
		CategoryBrandRouter.POST("add_new_categorybrand", middlewares.JWTAuth(), middlewares.IsadminAuth(), brands.NewCategoryBrand) //此接口需要管理员权限，新增分类品牌
		CategoryBrandRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), brands.DeleteCategoryBrand)             //删除品牌分类
		CategoryBrandRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), brands.UpdateCategoryBrand)                //更新品牌分类
		CategoryBrandRouter.GET("/:id", brands.GetCategoryBrandList)                                                                 //通过分类id获取其下所有品牌
	}
}
