package repositories

import (
	"errors"
	"fmt"

	"github.com/sai-mudike/job-application-tracker/internals/database"
	appErrors "github.com/sai-mudike/job-application-tracker/internals/errors"
	"github.com/sai-mudike/job-application-tracker/internals/models"
)

func CreateApplication(application *models.Application) error {
	query := `
INSERT INTO applications(user_id,company_name,job_title,job_url,location,employment_type,salary_min,salary_max,status,applied_at,notes)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
RETURNING id;
`
	smt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer smt.Close()

	row := smt.QueryRow(application.User_id, application.Company_name, application.Job_title, application.Job_url, application.Location, application.Employment_type, application.Salary_min, application.Salary_max, application.Status, application.Applied_at, application.Notes)

	err = row.Scan(&application.Id)
	if errors.As(err, &appErrors.PQErr) {
		if appErrors.PQErr.Code == "23505" {
			return appErrors.ErrDuplicateApplication
		}
		if appErrors.PQErr.Code == "23514" {
			return appErrors.ErrInvalidSalaryRange
		}
	}
	return err
}

func GetAllApplications(user_id int64, queryList models.ApplicationQuery) ([]models.Application, error) {
	query := `
	SELECT * FROM applications WHERE applications.user_id=$1
	`
	args := []any{user_id}
	argsCount := 2

	if queryList.Status != "" {
		query += fmt.Sprintf(" AND status=$%d", argsCount)
		args = append(args, queryList.Status)
		argsCount++
	}

	if queryList.EmploymentType != "" {
		query += fmt.Sprintf(" AND employment_type=$%d", argsCount)
		args = append(args, queryList.EmploymentType)
		argsCount++
	}
	if queryList.OrderBy != "" {
		query += fmt.Sprintf(" ORDER BY %s %s", queryList.SortBy, queryList.OrderBy)
	}

	if queryList.Page >= 1 {
		offset := (queryList.Page - 1) * queryList.Limit
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argsCount, argsCount+1)
		args = append(args, queryList.Limit, offset)

	}
	fmt.Println(query)
	rows, err := database.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var applications []models.Application
	for rows.Next() {
		var singleApplication models.Application
		err = rows.Scan(&singleApplication.Id, &singleApplication.User_id, &singleApplication.Company_name, &singleApplication.Job_title, &singleApplication.Job_url, &singleApplication.Location, &singleApplication.Employment_type, &singleApplication.Salary_min, &singleApplication.Salary_max, &singleApplication.Status, &singleApplication.Applied_at, &singleApplication.Notes, &singleApplication.Created_at, &singleApplication.Updated_at)

		if err != nil {
			return nil, err
		}

		applications = append(applications, singleApplication)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return applications, nil

}

func GetApplicationByID(applicationID, userID int64) (models.Application, error) {
	query := `
 SELECT * FROM applications
 WHERE id =$1 AND user_id =$2;
 `
	row := database.DB.QueryRow(query, applicationID, userID)

	var singleApplication models.Application
	err := row.Scan(&singleApplication.Id, &singleApplication.User_id, &singleApplication.Company_name, &singleApplication.Job_title, &singleApplication.Job_url, &singleApplication.Location, &singleApplication.Employment_type, &singleApplication.Salary_min, &singleApplication.Salary_max, &singleApplication.Status, &singleApplication.Applied_at, &singleApplication.Notes, &singleApplication.Created_at, &singleApplication.Updated_at)

	if err != nil {
		return models.Application{}, appErrors.ErrApplicationNotFound
	}
	return singleApplication, nil

}

func UpdateApplication(updatedApplication models.Application, id int64) error {
	query := `
UPDATE applications
SET company_name=$1,job_title=$2,job_url=$3,location=$4,employment_type=$5,salary_min=$6,salary_max=$7,status=$8,applied_at=$9,notes=$10,updated_at=CURRENT_TIMESTAMP
WHERE id=$11
;
`

	smt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}
	_, err = smt.Exec(updatedApplication.Company_name, updatedApplication.Job_title, updatedApplication.Job_url, updatedApplication.Location, updatedApplication.Employment_type, updatedApplication.Salary_min, updatedApplication.Salary_max, updatedApplication.Status, updatedApplication.Applied_at, updatedApplication.Notes, id)

	return err
}

func DeleteApplication(id int64) error {
	query := `
	DELETE FROM applications WHERE id =$1;
	`
	_, err := database.DB.Exec(query, id)

	return err
}
