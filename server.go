package main

import (
	"net/http"
	"os"
	"io"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.GET("/", func( c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/save-multi", saveMulti)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.POST("/users", createUser)
func createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("show", show)
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

// e.POST("/save-multi", saveMulti)
func saveMulti(c echo.Context) error {
	name := c.FormValue("name")

	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	defer dst.Close()


	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "<b>")
}
