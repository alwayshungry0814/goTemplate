package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int        `json:"id" db:"id"`
	Email       string     `json:"email" db:"email"`
	PhoneNumber string     `json:"phone_number" db:"phone_number"`
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func GetUsers(DB *sqlx.DB) ([]User, error) {
	users := []User{}
	err := DB.Select(&users, "SELECT id, email, phone_number FROM users;")
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *User) GetUser(DB *sqlx.DB) error {
	return DB.QueryRowx("SELECT email, phone_number FROM users WHERE id=$1;", u.ID).StructScan(u)
}

func (u *User) CreateUser(DB *sqlx.DB) error {

	err := DB.Get(u, "INSERT INTO users(id, email, phone_number, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING *",
		u.ID,
		u.Email,
		u.PhoneNumber,
		time.Now(),
		time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateUser(DB *sqlx.DB) error {

	err := DB.Get(u, "UPDATE users SET email=$1, phone_number=$2, updated_at=$3 WHERE id=$4 RETURNING *",
		u.Email,
		u.PhoneNumber,
		time.Now(),
		u.ID)

	if err != nil {
		return err
	}

	return nil

}

func (u *User) DeleteUser(DB *sqlx.DB) error {
	_, err := DB.Exec("DELETE FROM users WHERE id=$1", u.ID)

	return err
}
