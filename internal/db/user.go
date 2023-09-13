package db

import (
	"context"
	"gorm.io/gorm"
)

var userTableName = "user"

type User struct {
	gorm.Model
	Username  string `gorm:"column:username;type:varchar(254);not null"`
	Password  string `gorm:"column:password;type:varchar(254);not null"`
	Phone     string `gorm:"column:phone;type:varchar(100)"`
	Email     string `gorm:"column:email;type:varchar(100)"`
	AvaUrl    string `gorm:"column:ava_url;type:varchar(254)"`
	Signature string `gorm:"column:signature;type:text"`
}

func (u *User) TableName() string {
	return userTableName
}

func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUserById(ctx context.Context, ID uint64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", ID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CreateUser(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(user).Error
}

func UpdateUser(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Updates(user).Error
}
