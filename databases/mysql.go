package databases

import (
	"fmt"

	"github.com/tuacoustic/go-gin-example/entities"
	"github.com/tuacoustic/go-gin-example/utils/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect to database ~> MYSQL
func MysqlConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		// log.Fatalf("models.Setup err: %v", err)
		return nil, err
	}

	// defer func() {
	// 	dbDB, err := db.DB()
	// 	if err != nil {
	// 		log.Fatalf("failed to get DB instance: %v", err)
	// 	}
	// 	dbDB.Close()
	// }()
	return db, nil
}

func MysqlAuto() bool {
	db, err := MysqlConnect()
	if err != nil {
		// log.Fatal(err)
		return false
	}

	err = db.AutoMigrate(&entities.User{}, &entities.YoutubeTranscripts{})
	if err != nil {
		// log.Fatal(err)
		return false
	}

	return true
}
