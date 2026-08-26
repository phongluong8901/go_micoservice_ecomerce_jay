package service

import (
	"errors"
	"log"
	"proj_1/internal/domain"
	"proj_1/internal/dto"
	"proj_1/internal/helper"
	"proj_1/internal/repository"
	"time"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

// Signup(email string, pasword string, phone string)
func (s UserService) Signup(input dto.UserSignup) (string, error) {
	log.Println(input) //log create table User

	// call helper to hash
	hPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword, //change with new hash password
		Phone:    input.Phone,
	})

	//generate token
	log.Println(user)

	// userInfo := fmt.Sprintf("%v, %v, %v", user.ID, user.Email, user.UserType)

	//call db to create user
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	//perform some db operation

	//business logic
	user, err := s.Repo.FindUser(email)

	return &user, err
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	//compare password
	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	//generate token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) isVerifiedUser(id uint) bool {
	currentUser, err := s.Repo.FindUserById(id)

	return err == nil && currentUser.Verified
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	//if user already verified
	if s.isVerifiedUser(e.ID) {
		return 0, nil
	}

	//generate vetification code
	code, err := s.Auth.GenerateCode()
	if err != nil {
		return 0, err
	}

	// update user
	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = s.Repo.UpdateUser(e.ID, user)
	if err != nil {
		return 0, errors.New("unable to update vertification code")
	}

	//send SMS

	// return vertification code
	return code, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	//if user already verified
	if s.isVerifiedUser(id) {
		log.Println("verified...")
		return errors.New("user already verified")
	}

	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}

	//dùng để kiểm tra xem thời gian đã quá hạn (hết hạn) hay chưa.
	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	updateUser := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, updateUser)
	if err != nil {
		return errors.New("unable to update verify user")
	}

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
