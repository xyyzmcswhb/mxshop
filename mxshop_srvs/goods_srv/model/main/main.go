package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mxshop_srvs/goods_srv/global"
	models "mxshop_srvs/goods_srv/model"
	"os"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func genMD5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	//return md5.Sum([]byte(code))//需转换成字符串
	return hex.EncodeToString(Md5.Sum(nil))

}

func main() {
	Mysql2Es()
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:94144325hb@tcp(127.0.0.1:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢 SQL 阈值
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      true,        // 禁用彩色打印
	//	},
	//)
	//
	//// 全局模式
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		TablePrefix:   "",
	//		SingularTable: true,
	//	},
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	panic(err)
	//}
	////设置全局的logger，这个logger在我们执行每个sql语句的时候会打印每一行sql
	////sql才是最重要的，本着这个原则我尽量的给大家看到每个api背后的sql语句是什么
	//
	////定义一个表结构， 将表结构直接生成对应的表 - migrations
	//// 迁移 schema
	//_ = db.AutoMigrate(&models.Category{}, &models.Banner{},
	//	&models.Brands{}, &models.Goods{}, &models.GoodsCategoryBrand{}) //此处应该有sql语句

}

func Mysql2Es() {
	dsn := "root:94144325hb@tcp(127.0.0.1:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	host := fmt.Sprintf("http://127.0.0.1:9200")
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	global.EsClient, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false),
		elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	var goods []models.Goods
	db.Find(&goods)
	for _, g := range goods {
		esModel := models.EsGoods{
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

		_, err = global.EsClient.Index().Index(esModel.GetIndexName()).BodyJson(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
		if err != nil {
			return
		}
	}

}
