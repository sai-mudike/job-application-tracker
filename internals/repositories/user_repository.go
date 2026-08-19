package repositories

import (
	"errors"

	"github.com/sai-mudike/job-application-tracker/internals/database"
	appErrors "github.com/sai-mudike/job-application-tracker/internals/errors"
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
	if errors.As(err, &appErrors.PQErr) {
		if appErrors.PQErr.Code == "23505" {
			return appErrors.ErrEmailAlreadyExists
		}
	}
	return err
}

func VerifyUser(user *models.User) error {
	query := `
	SELECT id,hashed_pass FROM users WHERE user_name=$1;
	`

	row := database.DB.QueryRow(query, user.UserName)
	var hashed_pass string

	err := row.Scan(&user.Id, &hashed_pass)
	if err != nil {
		return appErrors.ErrUserNotFound
	}

	isPassValid := services.CheckPassword(user.Password, hashed_pass)

	if !isPassValid {
		return appErrors.ErrInvalidCredentials
	}

	return nil

}

func UserExists(userID int64) (bool, error) {
	var exists bool

	err := database.DB.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE id = $1
		)
	`, userID).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
