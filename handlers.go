package main

import (
	"database/sql"
	"fmt"
	"os"
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
	Secret      string `form:"secret"`
}

// DTO = Data Transfer Object. Will represent the
// clean and sanitized version of the input data.
type ProblemDTO struct {
	Title       string
	Text        string
	Constraints string
	Hint        string
	Solution    string
	IsProject   bool
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
			&problem.IsProject); err != nil {
			return problems, err
		}
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
	)
	if err != nil {
		//return an empty struct if error
		return problem, fmt.Errorf("error scanning Problem instance: %w", err)
	}

	return problem, nil
}
