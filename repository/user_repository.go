package repository

import (
	"database/sql"
	"log"

	"github.com/eulbyvan/enigma-university/model"
)

type UserRepository interface {
	// create
	// read
	GetById(id string) (model.User, error)
	// update
	// delete
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) GetById(id string) (model.User, error) {
	// implementasikan cara get by id
	var user model.User
	query := `	SELECT 	id,
						first_name,
						last_name,
						email,
						username,
						role,
						photo,
						created_at,
						updated_at
				FROM	users
				WHERE	id = $1`

	err := u.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Username,
		&user.Role,
		&user.Photo,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	log.Printf("DISINIIIIIIII BOSSSSS >>>>>>>>>>>>>>>>> %v", user)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// constructor
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
