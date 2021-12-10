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

func NewGormRepository() *GormRepo {
	db = configDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	defer db.Close()
	return &GormRepo{}
}

func (g *GormRepo) CreateUser(u *models.User) *models.User {
	db = configDB()
	db.NewRecord(u)
	result := db.Create(&u)

	defer db.Close()
	if result.Error != nil {
		panic(result.Error)
	}

	return u
}

func (g *GormRepo) GetUser(Id int) (*models.User, error) {
	db = configDB()
	var getUser models.User

	defer db.Close()
	if err := db.First(&getUser, Id).Error; err != nil {
		return nil, err
	}
	return &getUser, nil
}

func (g *GormRepo) UpdateUser(Id int, user *models.User) (*models.User, error) {
	db = configDB()
	defer db.Close()

	us := &models.User{}

	updatedUser := db.Model(us).Where("id = ?", Id).Updates(user)

	if updatedUser.Error != nil {
		return nil, updatedUser.Error
	}

	return us, nil
}

func (g *GormRepo) GetAllUsers() []*models.User {
	db = configDB()
	var allUsers []*models.User

	defer db.Close()

	db.Find(&allUsers)

	return allUsers
}

func (*GormRepo) DeleteUser(Id int) error {
	db = configDB()
	var user models.User
	result := db.Where("id = ?", Id).Delete(&user)
	fmt.Println(result)
	defer db.Close()

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (*GormRepo) CreatePost(p *models.Post) *models.Post {
	db = configDB()
	db.NewRecord(p)

	result := db.Create(&p)

	defer db.Close()

	if result.Error != nil {
		panic(result.Error)
	}

	return p
}

func configDB() *gorm.DB {
	config.Connect()
	return config.GetDB()
}
