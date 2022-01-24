package connectdb

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"tranning_golang/config"
)

var DB *gorm.DB

func Connnectdb() bool {

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	configdatabase.Username,
	configdatabase.Password,
	configdatabase.Host,
	configdatabase.Port,
	configdatabase.Database,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // -> Log các câu lệnh truy vấn database trong terminal
	})

	// Dừng chương trình nếu quá trình kết nối tới database xảy ra lỗi
	if err != nil {
		return false
	} else {
	return true

	}
	

}