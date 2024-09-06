package model

import (
	"context"
	"mxshop_srvs/goods_srv/global"
	"strconv"

	"gorm.io/gorm"
)

// 商品品牌分类
type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32       `json:"parent"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	Logo string `gorm:"type:varchar(200);default:' ';not null" json:"logo"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32    `gorm:"type:int;index:idx_category_brand,unique"` //与品牌ID建立联合唯一索引
	Category   Category //分类的外键

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

func (GoodsCategoryBrand) TableName() string {
	//重载表名
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(255);not null"`
	Url   string `gorm:"type:varchar(255);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

// 商品表结构的设计
type Goods struct {
	BaseModel
	CategoryID int32    `gorm:"type:int;not null"`
	Category   Category //分类的外键
	BrandsID   int32    `gorm:"type:int;not null"`
	Brands     Brands   //品牌外键

	OnSale   bool `gorm:"default:false;not null;comment:'是否已经上架'"`  //是否已经上架
	ShipFree bool `gorm:"default:false;not null;comment:'是否免运费'"`   //是否免运费
	IsNew    bool `gorm:"default:false;not null;comment:'是否为新品'"`   //是否为新品
	IsHot    bool `gorm:"default:false;not null;comment:'是否为热卖商品'"` //是否为热卖商品

	Name            string   `gorm:"type:varchar(50);not null;comment:'商品名称'"`
	GoodsSn         string   `gorm:"type:varchar(50);not null;comment:'商品编号'"`
	ClickNum        int32    `gorm:"type:int;default:0;not null;comment:'商品点击数'"`
	SoldNum         int32    `gorm:"type:int;default:0;not null;comment:'销量'"`
	FavNum          int32    `gorm:"type:int;default:0;not null;comment:'被收藏数量'"`
	MarketPrice     float32  `gorm:"not null;comment:'市面售价'"`
	ShopPrice       float32  `gorm:"not null;comment:'本店售价'"`
	GoodsBrief      string   `gorm:"type:varchar(255);not null;comment:'商品简介'"`
	Images          GormList `gorm:"type:varchar(1000);not null"`
	DescImages      GormList `gorm:"type:varchar(1000);not null"`
	GoodsFrontImage string   `gorm:"type:varchar(255);not null;comment:'商品封面图'"`
}

// 钩子函数，做转换
func (g *Goods) AfterCreate(tx *gorm.DB) (err error) {
	esModel := EsGoods{
		ID:          g.ID,
		CategoryID:  g.CategoryID,
		BrandsID:    g.BrandsID,
		OnSale:      g.OnSale,
		ShipFree:    g.ShipFree,
		IsNew:       g.IsNew,
		IsHot:       g.IsHot,
		Name:        g.Name,
		ClickNum:    g.ClickNum,
		SoldNum:     g.SoldNum,
		FavNum:      g.FavNum,
		MarketPrice: g.MarketPrice,
		GoodsBrief:  g.GoodsBrief,
		ShopPrice:   g.ShopPrice,
	}

	//id为必传字段
	_, err = global.EsClient.Index().Index(esModel.GetIndexName()).
		BodyJson(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	if err != nil {
		return
	}
	return nil
}

// 钩子函数做更新
func (g *Goods) AfterUpdate(tx *gorm.DB) (err error) {
	esModel := EsGoods{
		ID:          g.ID,
		CategoryID:  g.CategoryID,
		BrandsID:    g.BrandsID,
		OnSale:      g.OnSale,
		ShipFree:    g.ShipFree,
		IsNew:       g.IsNew,
		IsHot:       g.IsHot,
		Name:        g.Name,
		ClickNum:    g.ClickNum,
		SoldNum:     g.SoldNum,
		FavNum:      g.FavNum,
		MarketPrice: g.MarketPrice,
		GoodsBrief:  g.GoodsBrief,
		ShopPrice:   g.ShopPrice,
	}

	_, err = global.EsClient.Update().Index(esModel.GetIndexName()).
		Doc(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	if err != nil {
		return
	}
	return nil
}

func (g *Goods) AfterDelete(tx *gorm.DB) (err error) {
	_, err = global.EsClient.Delete().Index(EsGoods{}.GetIndexName()).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
