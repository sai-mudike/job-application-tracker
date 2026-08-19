package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:password@localhost:5432/job_tracker?sslmode=disable")

	if err != nil {
		panic("Could not connec to postgres")
	}

	err = DB.Ping()
	if err != nil {
		panic("could not connec to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUsersTable := `
CREATE TABLE IF NOT EXISTS users(
id SERIAL PRIMARY KEY,
user_name VARCHAR(255) NOT NULL UNIQUE,
hashed_pass TEXT NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Unable to create users table")
	}

	createApplicationsTable := `
CREATE TABLE IF NOT EXISTS applications(
id SERIAL PRIMARY KEY,
user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL ,
company_name VARCHAR(255) NOT NULL,
job_title VARCHAR(255) NOT NULL,
job_url TEXT NOT NULL,
location VARCHAR(255) NOT NULL,
employment_type VARCHAR(50) NOT NULL,
salary_min INTEGER CHECK(salary_min >= 0),
salary_max INTEGER CHECK(salary_max >= 0),
status VARCHAR(50) NOT NULL,
applied_at TIMESTAMP WITH TIME ZONE,
notes TEXT,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
CHECK(salary_min IS NULL OR salary_max IS NULL OR salary_max >= salary_min),
UNIQUE(user_id,job_url),
CHECK (status IN (
    'applied',
    'interview',
    'offer',
    'rejected',
    'withdrawn',
    'accepted'
)),
CHECK (employment_type IN(
'full_time',
    'part_time',
    'contract',
    'internship',
    'temporary'
	))
);
`
	_, err = DB.Exec(createApplicationsTable)

	if err != nil {
		panic("could not create the applications table")
	}

	createResumeTable := `
CREATE TABLE IF NOT EXISTS resumes(
id SERIAL PRIMARY KEY,
user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
name VARCHAR(255) NOT NULL,
file_path TEXT NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
UNIQUE(user_id,file_path)
);
`
	_, err = DB.Exec(createResumeTable)

	if err != nil {
		panic("could not create resumes table")
	}
}
