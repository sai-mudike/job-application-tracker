package repositories

import (
	"errors"

	"github.com/sai-mudike/job-application-tracker/internals/database"
	appErrors "github.com/sai-mudike/job-application-tracker/internals/errors"
	"github.com/sai-mudike/job-application-tracker/internals/models"
)

func CreateResume(resume models.Resume) error {
	query := `
INSERT INTO resumes(user_id,name,file_path)
VALUES($1,$2,$3);
`

	smt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer smt.Close()

	_, err = smt.Exec(resume.UserID, resume.Name, resume.File_path)

	if err != nil {

		if errors.As(err, &appErrors.PQErr) {
			if appErrors.PQErr.Code == "23505" {
				return appErrors.ErrResumeAlreadyExists
			}
		}
		return err
	}
	return nil
}

func GetAllResumes(user_id int64) ([]models.Resume, error) {
	query := `
SELECT user_id,name,file_path FROM resumes WHERE user_id = $1;
`
	rows, err := database.DB.Query(query, user_id)
	if err != nil {
		return nil, appErrors.ErrInternal
	}
	defer rows.Close()

	var resumes []models.Resume

	for rows.Next() {
		var resume models.Resume
		err := rows.Scan(&resume.UserID, &resume.Name, &resume.File_path)

		if err != nil {
			return nil, err
		}

		resumes = append(resumes, resume)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resumes, nil

}

func GetResumeByID(resumeID, userID int64) (*models.Resume, error) {

	query := `
	SELECT name,file_path FROM resumes WHERE id=$1 AND user_id=$2; 
	`
	smt, err := database.DB.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer smt.Close()

	row := smt.QueryRow(resumeID, userID)

	var resume models.Resume

	err = row.Scan(&resume.Name, &resume.File_path)

	if err != nil {
		return nil, appErrors.ErrResumeNotFound

	}
	return &resume, nil
}

func DeleteResume(resume_id, user_id int64) error {
	query := `
	DELETE FROM resumes WHERE id=$1 AND user_id =$2;
	`

	_, err := database.DB.Exec(query, resume_id, user_id)
	if err != nil {
		return err
	}
	return nil
}
