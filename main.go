package main

import (
	"todo_app/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())  // EchoフレームワークでCORSを有効にする

	e.GET("/todos", handlers.GetTasks)
	e.POST("/todos", handlers.AddTask)
	e.PUT("/todos/:id", handlers.UpdateTask)
	e.DELETE("/todos/:id", handlers.DeleteTask)

	e.GET("/users", handlers.GetUsers)
	e.POST("/users", handlers.AddUser)
	e.GET("/users/:id", handlers.GetUserByID)

	e.Static("/", "public")

	// e.GET("/", func(c echo.Context) error {
	// 	return c.File("public/index.html") // `public/index.html`を返す
	// })

	e.Logger.Fatal(e.Start(":8080"))
}