package service

import (
	"log"
	"proj_1/internal/domain"
	"proj_1/internal/dto"
)

type UserService struct {
}

// Signup(email string, pasword string, phone string)
func (s UserService) Signup(input dto.UserSignup) (string, error) {
	log.Println(input)

	//call db to create user
	return "this-is-my-token-id", nil
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	//perform some db operation

	//business logic

	return nil, nil
}

func (s UserService) Login(input any) (string, error) {

	return "", nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {

	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {

	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {

	return nil
}

func (s UserService) GetProfile(id uint, input any) (*domain.User, error) {

	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {

	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) CreateCart(id uint, u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (interface{}, error) {

	return nil, nil
}
