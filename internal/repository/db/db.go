package db

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDbContext() {
	fmt.Println("初始化數據庫")
	// DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatalf("Error opening database: %q", err)
	// }

	fmt.Println("dsn:", os.Getenv("DATABASE_URL"))
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		fmt.Println("connet database fail,pleaes check parametre", err)

		//退出程序
		os.Exit(1)
	}

	// DB.Migrator().DropTable(&product.Camp{}, &model.Account{}, &model.Tag{}, &model.TagMap{})

	// //migrate table
	// _ = DB.AutoMigrate(&model.Camp{}, &model.Account{}, &model.Tag{}, &model.TagMap{})
}

func BeginTranscation(DB *gorm.DB, process func(tx *gorm.DB) error) error {
	tx := DB.Begin()

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
