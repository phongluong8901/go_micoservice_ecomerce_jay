package repository

import (
	"errors"
	"log"
	"proj_1/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)

	//more
}

type userRepository struct {
	db *gorm.DB
}

// constructors
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {
	err := r.db.Create(&usr).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return usr, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User

	//tìm kiếm bản ghi đầu tiên khớp với điều kiện email và gán kết quả vào biến user
	err := r.db.First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user doese not exist")
	}

	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	//tìm kiếm bản ghi dựa vào Khóa chính (Primary Key), cụ thể ở đây là tìm người dùng (User) theo id.
	err := r.db.First(&user, id).Error
	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user doese not exist")
	}

	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	//dùng để cập nhật dữ liệu cho một bản ghi dựa theo id, đồng thời lấy lại toàn bộ thông tin mới nhất của bản ghi đó sau khi cập nhật và đổ ngược lại vào biến user.
	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error
	if err != nil {
		log.Printf("error on update %v", err)
		return domain.User{}, errors.New("failed upadte user")
	}

	return user, nil
}
