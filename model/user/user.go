package user

import (
	"fmt"

	"SMZT/model"

	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model
	HashPassword string `json:"hash_password" gorm:"column:hash_password;" binding:"required"`
	StudentId    string `json:"student_id" gorm:"column:student_id;unique™"`
}

func (u *UserModel) TableName() string {
	return "users"
}

// Create ... create user
func (u *UserModel) CreateUser() error {
	tx := model.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Save ... save user.
func (u *UserModel) Save() error {
	return model.DB.Self.Save(u).Error
}

// Get Information
func (user *UserModel) GetInfo(email string) error {
	fmt.Printf("email:%s\n", email)
	if err := model.DB.Self.Model(UserModel{}).
		Where("email = ?", email).
		First(user).Error; err != nil {
		fmt.Println("error", err)
		return err
	}
	return nil
}

// func IfExist(id, email, name string) error {
// 	var user1 UserModel
// 	var user2 UserModel
// 	var user3 UserModel

// 	err1 := model.DB.Self.Debug().Where("student_id=?", id).Find(&user1).Error
// 	err2 := model.DB.Self.Debug().Where("email=?", email).Find(&user2).Error
// 	err3 := model.DB.Self.Debug().Where("name=?", name).Find(&user3).Error

// 	s := []string{""}
// 	i := 0

// 	if err1 == nil {
// 		s = append(s, "*学号*")
// 		i++
// 	}

// 	if err2 == nil {
// 		s = append(s, "*邮箱*")
// 		i++
// 	}

// 	if err3 == nil {
// 		s = append(s, "*姓名*")
// 		i++

// 	}

// 	if i > 0 {
// 		s = append(s, "已被注册")
// 	}

// 	if i > 0 {
// 		return errors.New(fmt.Sprintf("%s", s))
// 	}

// 	return nil

// }
