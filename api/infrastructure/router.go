package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/watarun54/go-video-chat/api/interfaces/controllers"
)

func Init() {
	// Echo instance
	e := echo.New()

	authController := controllers.NewAuthController(NewSqlHandler())
	userController := controllers.NewUserController(NewSqlHandler())
	commentController := controllers.NewCommentController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", func(c echo.Context) error { return authController.Login(c) })

	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(controllers.NewJWTConfig()))
	api.GET("/comments", func(c echo.Context) error { return commentController.Index(c) })
	api.GET("/comments/:id", func(c echo.Context) error { return commentController.Show(c) })

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
