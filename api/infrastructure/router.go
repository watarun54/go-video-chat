package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/watarun54/go-video-chat/api/interfaces/controllers"
)

func Init() {
	// Echo instance
	e := echo.New()

	userController := controllers.NewUserController(NewSqlHandler())
	commentController := controllers.NewCommentController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })

	e.GET("/comments", func(c echo.Context) error { return commentController.Index(c) })
	e.GET("/comments/:id", func(c echo.Context) error { return commentController.Show(c) })

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
