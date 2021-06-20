package model

import (
	"golearn/config"
	"time"
)

type User struct {
	ID       	int    `json:"id"`
	Email    	string `json:"email"`
	Password 	string `json:"password"`
	Name     	string `json:"name"`
	Phone     	string `json:"phone"`
	CreatedAt	string `json:"create_at"`
	UpdatedAt	string `json:"updated_at"`
	DeletedAt 	string `json:"deleted_at"`
}

func GetUsers(start, count int) ([]User, error) {
	db := config.DbConn()
	defer db.Close()
	rows, err := db.Query("SELECT id, email, password, name, phone FROM users LIMIT ? OFFSET ?", count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.Name, &u.Phone); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (u *User) CreateUser() error {
	db := config.DbConn()
	defer db.Close()

	u.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	u.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	_, err := db.Exec("INSERT INTO users(email, password, name, phone, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)", u.Email, u.Password, u.Name, u.Phone, u.CreatedAt, u.UpdatedAt)

	if err != nil {
		return err
	}
	return nil
}
