package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	db       *sql.DB
	ID       string
	Name     string
	Email    string
	Password string
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (c *User) Create(name string, email string, password string) (User, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO user (id, name, email, password) values ($1, $2, $3, $4)", id, name, email, password)
	if err != nil {
		return User{}, err
	}
	return User{Name: name, Email: email, Password: password}, nil
}

func (c *User) FindOne(userID string) (User, error) {
	_, err := c.db.Exec("SELECT name, email, password FROM user WHERE id = $1", userID)
	if err != nil {
		return User{}, err
	}
	return User{Name: "name", Email: "email", Password: "password"}, nil
}

func (c *User) FindAll() ([]User, error) {
	rows, err := c.db.Query("SELECT id, name, email, password FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var id, name, email, password string
		if err := rows.Scan(&id, &name, &email, &password); err != nil {
			return nil, err
		}
		users = append(users, User{ID: id, Name: name, Email: email, Password: password})
	}
	return users, nil
}
