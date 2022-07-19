package repository

import (
	"fmt"
	"linebot/internal/model"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDbContext() {
	fmt.Println("初始化數據庫")
	// db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatalf("Error opening database: %q", err)
	// }

	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		fmt.Println("connet database fail,pleaes check parametre", err)

		//退出程序
		os.Exit(1)
	}

	//migrate table
	_ = db.AutoMigrate(&model.Camp{})
}

func InitDb_local() {
	dsn := "postgres://wkkckoevwhzsti:756a9057eb6ed63db56541deeb46335868aad00e5f26e0d942d5f021f7df062d@ec2-52-71-69-66.compute-1.amazonaws.com:5432/d2qfr5f81g7vlm"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("connet database fail,pleaes check parametre", err)

		//退出程序
		os.Exit(1)
	}

	//migrate table
	_ = db.AutoMigrate(&model.Camp{}, &model.Account{}, &model.Tag{}, &model.TagMap{})
	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//test
	db.Create(&model.Camp{
		CampName:       "Test",
		AddressCountry: "hi",
		AddressCity:    "123",
		AddressDetail:  "9999",
		//TagList:        []string{"get"},
	})

	//	result := db.First(&model.Camp{})

}

func BeginTranscation(db *gorm.DB, process func(tx *gorm.DB) error) error {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := process(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error

	return err
}
