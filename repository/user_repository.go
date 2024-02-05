package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/eulbyvan/enigma-university/model"
	"github.com/eulbyvan/enigma-university/model/dto/req"
)

type UserRepository interface {
	// create
	Post(user model.User) error
	// read
	GetById(id string) (model.User, error)
	// read
	GetAll() ([]model.User, error)
	// update
	Update(id string, user model.User) error
	// delete
	Delete(id string) error
	// Login
	Login(credential req.Credential) bool
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

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetAll() ([]model.User, error) {
	// implementasikan cara get by id
	var users []model.User

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
				`

	rows, err := u.db.Query(query)

	for rows.Next() {

		rows.Scan(
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

		users = append(users, user)
	}

	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (u *userRepository) Post(user model.User) error {
	// implementasikan cara get by id
	// var user model.User
	query := `	INSERT INTO users(first_name, last_name, email, username, password, role, photo)
		VALUES ($1, $2, $3, $4, $5, $6, $7);`

	fmt.Println(query)

	roleId, errConv := strconv.Atoi(user.Role)

	if errConv != nil {
		return errConv
	}

	_, err := u.db.Exec(query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Username,
		user.Password,
		model.Role(roleId),
		user.Photo,
	)

	user.Password = ""

	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Update(id string, user model.User) error {
	// implementasikan cara get by id
	// var user model.User
	updateQuery := ""
	param := 1
	var stringUpdate []any

	if user.FirstName != "" {
		updateQuery += fmt.Sprintf("first_name=$%v, ", param)
		stringUpdate = append(stringUpdate, user.FirstName)
		param++
	}
	if user.LastName != "" {
		updateQuery += fmt.Sprintf("last_name=$%v, ", param)
		stringUpdate = append(stringUpdate, user.LastName)
		param++
	}
	if user.Email != "" {
		updateQuery += fmt.Sprintf("email=$%v, ", param)
		stringUpdate = append(stringUpdate, user.Email)
		param++
	}
	if user.Username != "" {
		updateQuery += fmt.Sprintf("username=$%v, ", param)
		stringUpdate = append(stringUpdate, user.Username)
		param++
	}
	if user.Password != "" {
		updateQuery += fmt.Sprintf("password=$%v, ", param)
		stringUpdate = append(stringUpdate, user.Password)
		param++
	}
	if user.Role != "" {
		updateQuery += fmt.Sprintf("role=$%v, ", param)
		roleId, err := strconv.Atoi(user.Role)
		if err != nil {
			return err
		}
		stringUpdate = append(stringUpdate, model.Role(roleId))
		param++
	}
	if user.Photo != "" {
		updateQuery += fmt.Sprintf("photo=$%v, ", param)
		stringUpdate = append(stringUpdate, user.Photo)
		param++
	}
	stringUpdate = append(stringUpdate, id)

	query := `UPDATE public.users
				SET 				
				` + updateQuery + ` 				 
				updated_at=NOW()
				WHERE id = $` + strconv.Itoa(param) + `;`

	_, err := u.db.Exec(query,
		stringUpdate...,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Delete(id string) error {
	// implementasikan cara get by id
	query := `	DElETE
				FROM	users
				WHERE	id = $1`

	_, err := u.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Login(credential req.Credential) bool {
	// implementasikan cara get by id
	var username string

	query := `SELECT username FROM users WHERE username=$1 AND password=$2`

	row := u.db.QueryRow(query, credential.Username, credential.Password)

	if row.Scan(&username); username != "" {
		return true
	}

	return false
}

// constructor
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
