package usecase

import (
	"fmt"

	"github.com/eulbyvan/enigma-university/model"
	"github.com/eulbyvan/enigma-university/repository"
)

type UserUseCase interface {
	// Register(newUser model.User) string
	FindById(id string) (model.User, error)
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

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
