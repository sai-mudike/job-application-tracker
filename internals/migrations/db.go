package migrations

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:password@localhost:5432/job_tracker?sslmode=disable")

	if err != nil {
		panic(fmt.Sprintf("Could not coonect to the DB %v,", err))
	}

	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Connection lost with Database %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createCompanyTable := `
	CREATE TABLE IF NOT EXISTS companies (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	name VARCHAR(50) NOT NULL,
	website VARCHAR(250),
	location VARCHAR(100),
	industry VARCHAR(100),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(createCompanyTable)

	if err != nil {
		panic(fmt.Sprintf("Could not create the companies table%v", err))
	}
}
