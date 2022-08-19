package repository

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
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

	fmt.Println("dsn:", os.Getenv("DATABASE_URL"))
	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		fmt.Println("connet database fail,pleaes check parametre", err)

		//退出程序
		os.Exit(1)
	}

	db.Migrator().DropTable(&Camp{}, &User{}, &Tag{}, &TagMap{})

	//migrate table
	_ = db.AutoMigrate(&Camp{}, &User{}, &Tag{}, &TagMap{})
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
