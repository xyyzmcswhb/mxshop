package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	models "mxshop_srvs/inventory_srv/model"
	"os"
	"time"

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
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:94144325hb@tcp(127.0.0.1:3306)/mxshop_inventory_srv?charset=utf8mb4&parseTime=True&loc=Local"

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
	//设置全局的logger，这个logger在我们执行每个sql语句的时候会打印每一行sql
	//sql才是最重要的，本着这个原则我尽量的给大家看到每个api背后的sql语句是什么

	//定义一个表结构， 将表结构直接生成对应的表 - migrations
	// 迁移 schema
	//_ = db.AutoMigrate(&models.Inventory{}, &models.StockSellDetail{}) //此处应该有sql语句
	OrderDetail := models.StockSellDetail{
		OrderSn: "111",
		Status:  0,
		Detail:  []models.GoodsDetail{{1, 2}, {2, 3}},
	}
	db.Create(&OrderDetail)
	// 新增
	//db.Create(&Product{Code: sql.NullString{"D42", true}, Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整形主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//
	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{"", true}}) // 仅更新非零值字段
	//如果我们去更新一个product 只设置了price：200
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product， 并没有执行delete语句，逻辑删除
	//db.Delete(&product, 1)

	//盐值加密，将用户密码变成随机数加用户密码
	//salt, encodedPwd := password.Encode("generic password", nil)
	//newpassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println(newpassword)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check)

	//cutom options

	//
	//passwordinfo := strings.Split(newpassword, "$")
	//fmt.Println(passwordinfo)
	//check := password.Verify("generic password", passwordinfo[2], passwordinfo[3], options)
	//fmt.Println(check)
	//fmt.Println(genMD5("123455"))
}