// the main entrypoint to the application using echo
// framework. this pacakge only handles routing,
// and should not contain application logic.
package main

import (
	"csweekly-echo/db"
	"database/sql"
	"html/template"
	"io"
	"net/http"

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

	e.POST("/submit-post", postSubmitProblem)

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

func postSubmitProblem(c echo.Context) error {
	var postData ProblemPost
	err := c.Bind(&postData)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	problem, err := PostProblem(postData)
}
