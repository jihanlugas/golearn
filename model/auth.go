package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golearn/config"
	"time"
)

const JWT_KEY = "MyRandom"

var JwtKey = []byte(JWT_KEY)

type Auth struct {
	Email		string	`json:"email"`
	Password 	string	`json:"password"`
}

type Credentials struct {
	Email		string	`json:"email"`
	Password 	string	`json:"password"`
}

type Claims struct {
	Email	string	`json:"email"`
	jwt.StandardClaims
}

func (c *Credentials) Signin() error {
	db := config.DbConn()
	defer db.Close()

	var u User

	err := db.QueryRow("SELECT id, email, password, name, phone FROM users where email = ?",
		c.Email).Scan(&u.ID, &u.Email, &u.Password, &u.Name, &u.Phone)

	if err != nil {
		return err
	}

	if u.Password != HashPassword(c.Password) {
		return errors.New("Wrong password")
	}

	return nil
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 60)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtKey)
}