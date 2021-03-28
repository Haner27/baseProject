package repository

import (
	"baseProject/datasource"
	"baseProject/model"
	"gorm.io/gorm"
	"time"
)

type (
	IUserRepository interface {
		CreateUser(username, password, nickname string, gender int, birthday time.Time) *model.UserModel
		GetUserById(id int) *model.UserModel
		GetUsers(maxUserId, perPage int) []model.UserModel
	}
	UserRepository struct {
		mysqlDB *datasource.MysqlDB
	}
)

func NewUserRepository(mysqlDB *datasource.MysqlDB) *UserRepository {
	return &UserRepository{
		mysqlDB: mysqlDB,
	}
}

func (u *UserRepository) CreateUser(
	username, password, nickname string, gender int, birthday time.Time) *model.UserModel {

	tx := u.mysqlDB.Cli.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil
	}

	var user model.UserModel
	results := tx.Set("gorm:query_option", "FOR UPDATE").Where("username = ?", username).First(&user)
	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return nil
	}
	if results.Error == gorm.ErrRecordNotFound {
		user = model.UserModel{
			Username: username,
			Password: password,
			Nickname: nickname,
			Gender:   gender,
			Birthday: birthday,
		}
		err := tx.Create(&user).Error
		if err != nil {
			tx.Rollback()
		}
	}
	if err := tx.Commit().Error; err != nil {
		return nil
	}
	return &user
}

func (u *UserRepository) GetUserById(id int) *model.UserModel {
	var user model.UserModel
	results := u.mysqlDB.Cli.Where("id = ?", id).First(&user)
	if results.Error == gorm.ErrRecordNotFound {
		return nil
	}
	if results.Error == nil {
		return &user
	}
	return nil
}


func (u *UserRepository) GetUsers(maxUserId, perPage int) []model.UserModel {
	var users []model.UserModel
	u.mysqlDB.Cli.Where("id > ?", maxUserId).Order("id").Limit(perPage).Find(&users)
	return users
}
