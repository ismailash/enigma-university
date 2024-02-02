package manager

import "github.com/eulbyvan/enigma-university/repository"

type RepoManager interface {
	NewUserRepo() repository.UserRepository
	// CourseRepo
	// Enrollment (transaction)
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) NewUserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
