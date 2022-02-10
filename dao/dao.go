package dao

import (
	"NetClip2/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *model.User)
	SearchUser(username string) *model.User
	DeleteUser(username string)
	UpdateUserContent(user *model.User)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func LoadEnv() string{
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Println("unable to load .env file")

    }

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME)
	return DB_DSN

}

func init() {
	dsn := LoadEnv()
	// dsn := "root:1234.com@@tcp(localhost:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	//创建表
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	//插入数据
	// p := Product{Name: "Product1", Price: 100}
	// db.Create(&p)
}



func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(&user)
}

func (mgr *manager) SearchUser(username string) *model.User {
	var user model.User
	mgr.db.First(&user, "username = ?", username)
	return &user
}
func (mgr *manager) DeleteUser(username string) {
	mgr.db.Where("username = ?", username).Unscoped().Delete(&model.User{})
}
func (mgr *manager) UpdateUserContent(user *model.User) {
	mgr.db.Model(&user).Update("Cliptext", user.Cliptext)
	mgr.db.Model(&user).Update("Expired", user.Expired)
}
