package repositories

import (
	"fmt"
	"project/config"
	"project/models"

	"github.com/jinzhu/gorm"
)

type GormRepo struct{}

var (
	db *gorm.DB
)

func NewGormRepository() UserRepository {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&models.User{})
	defer db.Close()
	return &GormRepo{}
}

func (g *GormRepo) CreateUser(u *models.User) *models.User {
	config.Connect()
	db = config.GetDB()
	db.NewRecord(u)
	result := db.Create(&u)

	defer db.Close()
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(result)
	return u
}

func (g *GormRepo) GetUser(Id int) *models.User {
	config.Connect()
	db = config.GetDB()
	var getUser models.User

	defer db.Close()
	_ = db.Where("ID = ?", Id).Find(&getUser)
	return &getUser
}
