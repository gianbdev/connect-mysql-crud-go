package repositories

import (
	"go-api-crud/config"
	"go-api-crud/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.DbDSN)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateUser(user *models.User) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.NamedExec(`INSERT INTO users (username, email, password) VALUES (:username, :email, :password)`, user)
	return err
}

func GetUserByUsername(username string) (*models.User, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user models.User
	err = db.Get(&user, "SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
