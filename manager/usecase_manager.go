package manager

import "github.com/eulbyvan/enigma-university/usecase"

type UseCaseManager interface {
	NewUserUseCase() usecase.UserUseCase
	// NewCourseUseCase
	// NewEnrollmentUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.NewUserRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
