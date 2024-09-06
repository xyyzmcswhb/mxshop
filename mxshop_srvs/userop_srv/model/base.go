package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type GormList []string

func (l *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &l)
}

// 实现sql.Scanner接口，将value扫描至jsonb
func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

type BaseModel struct {
	ID        int32     `gorm:"primarykey;type:int" json:"id"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeleteAt  gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}
