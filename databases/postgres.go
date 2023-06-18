package databases

import (
	"fmt"
	"log"

	"github.com/tuacoustic/go-gin-example/entities"
	"github.com/tuacoustic/go-gin-example/utils/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to database ~> MYSQL
func PostgresConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
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

func PgAuto() bool {
	db, err := PostgresConnect()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatal(err)
	}

	return true
}
