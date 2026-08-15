package repository

import (
	"github.com/sai-mudike/job-application-tracker/internals/migrations"
	"github.com/sai-mudike/job-application-tracker/internals/models"
)

func SaveCompany(company *models.Company) error {

	query := `
	INSERT INTO companies(user_id,name,website,location,industry)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING id;
	`
	smt, err := migrations.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer smt.Close()

	rows := smt.QueryRow(company.User_id, company.Name, company.Website, company.Location, company.Industry)

	err = rows.Scan(&company.Id)

	if err != nil {
		return err
	}

	// id, err := result.LastInsertId()

	// c.Id = id

	return err

}

func GetAllCompanies() ([]models.CompanySummary, error) {

	var companies []models.CompanySummary

	query := `
	SELECT id,name FROM companies;
	`
	rows, err := migrations.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var company models.CompanySummary

		err := rows.Scan(&company.Id, &company.Name)

		if err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, err

}

func GetCompanyByID(id int64) (models.Company, error) {
	query := `
SELECT id,name,website,location,industry FROM companies
WHERE companies.id = $1;
`
	smt, err := migrations.DB.Prepare(query)
	if err != nil {
		return models.Company{}, err
	}

	row := smt.QueryRow(id)

	var company models.Company

	err = row.Scan(&company.Id, &company.Name, &company.Website, &company.Location, &company.Industry)
	if err != nil {
		return models.Company{}, err
	}

	return company, nil

}
