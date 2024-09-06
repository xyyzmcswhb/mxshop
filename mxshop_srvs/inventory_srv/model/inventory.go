package model

import (
	"database/sql/driver"
	"encoding/json"
)

// 仓库
//
//	type Stock struct {
//		BaseModel
//		Name string
//		Address string
//	}
type GoodsDetail struct {
	Goods int32 `json:"goods"`
	Num   int32 `json:"num"`
}

type GoodsDetailList []GoodsDetail

func (l *GoodsDetailList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &l)
}

// 实现sql.Scanner接口，将value扫描至jsonb
func (g GoodsDetailList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 库存
type Inventory struct {
	BaseModel
	Goods  int32 `gorm:"type:int;index"`
	Stocks int32 `gorm:"type:int"` //库存数量
	//Stock  Stock//仓库服务，具体哪个仓库
	Version int32 `gorm:"type:int"` //分布式锁，乐观锁
}

// 库存
type InventoryNew struct {
	BaseModel
	Goods  int32 `gorm:"type:int;index"`
	Stocks int32 `gorm:"type:int"` //库存数量
	//Stock  Stock//仓库服务，具体哪个仓库
	Version int32 `gorm:"type:int"` //分布式锁，乐观锁
	Freeze  int32 `gorm:"type:int"` //冻结库存

}

type Delivery struct {
	Goods   int32  `gorm:"type:int;index"`
	Nums    int32  `gorm:"type:int"`
	OrderSn string `gorm:"type:varchar(255)"`
	Status  string `gorm:"type:varchar(255)"` //1.表示已扣减 2.表示已归还
}

type StockSellDetail struct {
	OrderSn string          `gorm:"type:varchar(255);index:idx_order_sn,unique"`
	Status  int32           `gorm:"type:varchar(255)"` //1.表示已扣减 2.表示已归还
	Detail  GoodsDetailList `gorm:"type:varchar(255)"`
}

func (StockSellDetail) TableName() string {
	return "stockSellDetail"
}
