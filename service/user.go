package service

import (
	"context"
	"myapp/config"
	"myapp/graph/model"
)

//UserCreate Create
func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	db := config.ConnectCockroach()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	user := model.User{
		Name: input.Name,
	}
	if err := db.Table("user").Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

//UserUpdate Update
func UserUpdate(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	db := config.ConnectCockroach()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Table("user").Where("id = ?", input.ID).Update("name", input.Name).Error; err != nil {
		return nil, err
	}

	return UserGetByID(ctx, input.ID)
}

//UserUpdate Delete
func UserDelete(ctx context.Context, id int) (string, error) {
	db := config.ConnectCockroach()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Table("user").Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return "", err
	}

	return "Success", nil
}

//UserGetByID Get By ID
func UserGetByID(ctx context.Context, id int) (*model.User, error) {
	db := config.ConnectCockroach()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user model.User
	if err := db.Table("user").Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

//UserGetAll Get All
func UserGetAll(ctx context.Context) ([]*model.User, error) {
	db := config.ConnectCockroach()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var users []*model.User
	if err := db.Table("user").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
