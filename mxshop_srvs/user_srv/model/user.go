package models

import (
	"time"

	"gorm.io/gorm"
)

// 定义公共字段
type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeleteAt  gorm.DeletedAt
	isDeleted bool `gorm:"column:is_deleted"`
}

// User /*
// 表结构
type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"` //通过手机号码查询用户
	Password string     `gorm:"type:varchar(100);not null"`
	Nickname string     `gorm:"type:varchar(20);"` //允许为空
	Birthday *time.Time `gorm:"type:datetime"`     //防止同步时出错
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女，male表示男'"`
	Role     int        `gorm:"column:role;default:1;type:int;comment: '1表示普通用户，2表示管理员用户'"` //用于区别当前用户是否是管理员
}
