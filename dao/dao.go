package dao

import (
	"NetClip2/model"

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

func init() {
	dsn := "root:1234.com@@tcp(localhost:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
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
