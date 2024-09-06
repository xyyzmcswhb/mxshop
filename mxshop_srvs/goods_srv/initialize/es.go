package initialize

import (
	"context"
	"fmt"
	"log"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"os"

	"github.com/olivere/elastic/v7"
)

func InitEs() {
	//初始化连接
	host := fmt.Sprintf("http://%s:%d/", global.ServerConfig.EsInfo.Host, global.ServerConfig.EsInfo.Port)
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	var err error
	global.EsClient, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false),
		elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	//新建mapping和index
	//首先判断index当前index是否存在
	exists, err := global.EsClient.IndexExists(model.EsGoods{}.GetIndexName()).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		if _, err = global.EsClient.CreateIndex(model.EsGoods{}.GetIndexName()).BodyString(model.EsGoods{}.GetMapping()).Do(context.Background()); err != nil {
			panic(err)
		}
	}
}
