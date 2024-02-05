package usecase

import (
	"fmt"

	Middleware "github.com/eulbyvan/enigma-university/middleware"
	"github.com/eulbyvan/enigma-university/model"
	"github.com/eulbyvan/enigma-university/model/dto/req"
	"github.com/eulbyvan/enigma-university/repository"
)

type UserUseCase interface {
	// Register(newUser model.User) string
	Post(user model.User) error
	// View
	FindById(id string) (model.User, error)
	//
	FindAll() ([]model.User, error)
	// Edit
	Update(id string, user model.User) error
	// Remove
	Delete(id string) error
	//
	Login(credential req.Credential) (string, error)
	//
	Logout(token string) error
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) FindById(id string) (model.User, error) {
	// implementasikan
	// optional: boleh menambahkan business logic
	user, err := u.repo.GetById(id)

	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}

	return user, err
}

func (u *userUseCase) FindAll() ([]model.User, error) {
	// implementasikan
	// optional: boleh menambahkan business logic
	users, err := u.repo.GetAll()

	if err != nil {
		return []model.User{}, fmt.Errorf("error query")
	}

	return users, err
}

// create
func (u *userUseCase) Post(user model.User) error {
	// implementasikan
	// optional: boleh menambahkan business logic
	err := u.repo.Post(user)

	if err != nil {
		return fmt.Errorf("Failed to create data user")
	}

	return err
}

// delete
func (u *userUseCase) Delete(id string) error {
	// implementasikan
	// optional: boleh menambahkan business logic
	err := u.repo.Delete(id)

	if err != nil {
		return fmt.Errorf("Failed to Delete data user, error : %v", err)
	}

	return err
}

func (u *userUseCase) Update(id string, user model.User) error {
	// implementasikan
	// optional: boleh menambahkan business logic
	err := u.repo.Update(id, user)

	if err != nil {
		return fmt.Errorf("Failed to Update data user, error : %v", err)
	}

	return err
}

func (u *userUseCase) Login(credential req.Credential) (string, error) {
	auth := u.repo.Login(credential)

	if auth == false {
		return "", fmt.Errorf("Username or Password Invalid!!")
	}

	for _, token := range Middleware.LocalToken {
		if token == credential.Username {
			return credential.Username, nil
		}
	}

	Middleware.LocalToken = append(Middleware.LocalToken, credential.Username)

	return credential.Username, nil
}

func (u *userUseCase) Logout(token string) error {

	for i, tokenL := range Middleware.LocalToken {
		if tokenL == token {
			Middleware.LocalToken = append(Middleware.LocalToken[:i], Middleware.LocalToken[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("The user is not logged in yet!!")
}

// constructor
func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
