package repositories

import (
	"errors"
	"fmt"

	"github.com/sai-mudike/job-application-tracker/internals/database"
	"github.com/sai-mudike/job-application-tracker/internals/models"
	"github.com/sai-mudike/job-application-tracker/internals/services"
)

func CreateUser(user *models.User) error {
	query := `
	INSERT INTO users(user_name,hashed_pass) VALUES($1,$2)
	RETURNING id;
	`
	smt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer smt.Close()

	hashedPassword, err := services.HashPassword(user.Password)
	if err != nil {
		return err
	}

	row := smt.QueryRow(user.UserName, hashedPassword)

	err = row.Scan(&user.Id)
	return err
}

func VerifyUser(user models.User) error {
	query := `
	SELECT id,hashed_pass FROM users WHERE user_name=$1;
	`

	row := database.DB.QueryRow(query, user.UserName)
	var hashed_pass string

	err := row.Scan(&user.Id, &hashed_pass)
	if err != nil {
		fmt.Println(err)
		return errors.New("Invalid credentials Please enter valid email id")
	}

	isPassValid := services.CheckPassword(user.Password, hashed_pass)

	if !isPassValid {
		return errors.New("Invalid credentials Please check password")
	}

	return nil

}
