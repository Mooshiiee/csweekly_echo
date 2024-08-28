// the main entrypoint to the application using echo
// framework. this pacakge only handles routing,
// and should not contain application logic.
package main

import (
	"csweekly-echo/db"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type Template struct {
	templates *template.Template
}

// implement the echo.Renderer intereface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Static("/", "public")
	e.Use(middleware.Logger())

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "cookie:_csrf",
		CookiePath:     "/",
		CookieDomain:   "csweekly.xyz",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	//init database
	database, err := db.InitDB()
	if err != nil {
		e.Logger.Fatal("failed to load database: %s", err)
	}

	// load html file templates
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return indexHandler(c, database)
	})

	e.GET("/problem/:id", func(c echo.Context) error {
		return getProblem(c, database)
	})

	e.GET("/submit", getSubmitPage)

	e.POST("/submit-post", func(c echo.Context) error {
		return postSubmitProblem(c, database)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func indexHandler(c echo.Context, database *sql.DB) error {
	rows, err := QueryProblems(database)

	if err != nil {
		c.Render(http.StatusInternalServerError, "index", err)
	}
	return c.Render(http.StatusOK, "index", rows)
}

func getProblem(c echo.Context, database *sql.DB) error {
	idString := c.Param("id")
	row, err := QuerySingleProblem(database, idString)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "index", err)
	}
	return c.Render(http.StatusOK, "problem", row)
}

func getSubmitPage(c echo.Context) error {
	return c.Render(http.StatusOK, "submit", nil)
}

// takes in POST request of a Problem object + Security key and then inserts the data
// to database if authorized
func postSubmitProblem(c echo.Context, database *sql.DB) error {
	var formData ProblemPost
	err := c.Bind(&formData)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if formData.Secret != os.Getenv("SECRET_KEY") {
		return c.String(http.StatusUnauthorized, "Invalid Secret Key")
	}

	ctx := c.Request().Context()

	datetime_now := time.Now().Format(time.RFC3339)
	fmt.Println(datetime_now)

	result, err := database.ExecContext(ctx,
		"INSERT INTO problems (title, text, constraints, hint, solution, isproject, datetime, weeknumber, poster, link, difficulty) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		formData.Title,
		formData.Text,
		formData.Constraints,
		formData.Hint,
		formData.Solution,
		formData.IsProject,
		datetime_now,
		formData.WeekNumber,
		formData.Poster,
		formData.Link,
		formData.Difficulty,
	)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Check and get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return c.String(http.StatusInternalServerError, "error: could not retrieve last insert ID")
	}

	// Checkc and get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.String(http.StatusInternalServerError, "error: could not retrieve rows affected")
	}

	// Return the response string
	response := fmt.Sprintf("Created ID: %d, Rows Affected: %d", id, rowsAffected)
	return c.String(http.StatusCreated, response)
}
