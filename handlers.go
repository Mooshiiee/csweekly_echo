package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type user struct {
	ID   int
	Name string
}

type problem struct {
	ID          int
	Text        string
	Title       string
	Hint        string
	Constraints string
	Solution    string
	IsProject   bool
	DateTime    string
	WeekNumber  int
	Poster      string
	Link        string
	Difficulty  string
}

// Represents the Binding of Raw post data to the Problem struct.
// This data should be handled as dirty and dangerous.
type ProblemPost struct {
	Title       string `form:"title"`
	Text        string `form:"text"`
	Constraints string `form:"constraints"`
	Hint        string `form:"hint"`
	Solution    string `form:"solution"`
	IsProject   bool   `form:"isproject"`
	WeekNumber  int    `form:"weeknumber"`
	Poster      string `form:"poster"`
	Link        string `form:"link"`
	Difficulty  string `form:"difficulty"`
	Secret      string `form:"secret"`
}

// takes in sql.db object, will return an array of users along with error
func QueryUsers(db *sql.DB) ([]user, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query : %v\n", err)
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user

		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
			//return users as nil, and error as the formatted message
		}

		users = append(users, user)
		fmt.Println(user.ID, user.Name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
		//return users as nil, and error as the formatted message
	}

	return users, nil
	//return users and nil for the error.
}

func QueryProblems(db *sql.DB) ([]problem, error) {
	rows, err := db.Query("SELECT * FROM problems ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//initiate a problem slice to hold data from the rows
	var problems []problem

	//loop through rows and use Scan() to assign db data to struct data
	for rows.Next() {
		//variable to hold instance data
		var problem problem
		if err := rows.Scan(
			&problem.ID,
			&problem.Title,
			&problem.Text,
			&problem.Hint,
			&problem.Constraints,
			&problem.Solution,
			&problem.IsProject,
			&problem.DateTime,
			&problem.WeekNumber,
			&problem.Poster,
			&problem.Link,
			&problem.Difficulty); err != nil {
			return problems, err
		}
		//parse DateTime
		parsedTime, err := time.Parse(time.RFC3339, problem.DateTime)
		if err != nil {
			return problems, err
		}
		//Reformat Datetime and replace the data
		formattedTime := parsedTime.Format("01-02-2006 03:04 PM")
		problem.DateTime = formattedTime

		problems = append(problems, problem)
	}
	if err = rows.Err(); err != nil {
		return problems, err
	}
	return problems, err
}

// takes in an sql.DB connection and the id from the request url, then returns
func QuerySingleProblem(db *sql.DB, idString string) (problem, error) {

	var problem problem

	err := db.QueryRow("SELECT * FROM problems WHERE id = ?", idString).Scan(
		&problem.ID,
		&problem.Title,
		&problem.Text,
		&problem.Hint,
		&problem.Constraints,
		&problem.Solution,
		&problem.IsProject,
		&problem.DateTime,
		&problem.WeekNumber,
		&problem.Poster,
		&problem.Link,
		&problem.Difficulty,
	)
	if err != nil {
		//return an empty struct if error
		return problem, fmt.Errorf("error scanning Problem instance: %w", err)
	}

	//parse DateTime
	parsedTime, err := time.Parse(time.RFC3339, problem.DateTime)
	if err != nil {
		return problem, fmt.Errorf("error parsing DateTime: %w", err)
	}
	//Reformat Datetime and replace the data
	formattedTime := parsedTime.Format("01-02-2006 03:04 PM")
	problem.DateTime = formattedTime

	return problem, nil
}
