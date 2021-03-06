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
	roomController := controllers.NewRoomController(NewSqlHandler())
	commentController := controllers.NewCommentController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", func(c echo.Context) error { return authController.Login(c) })
	e.POST("/signup", func(c echo.Context) error { return userController.Create(c) })

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(controllers.NewJWTConfig()))
	api.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	api.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	api.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
	api.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })

	api.GET("/comments", func(c echo.Context) error { return commentController.Index(c) })
	api.GET("/comments/:id", func(c echo.Context) error { return commentController.Show(c) })
	api.POST("/comments", func(c echo.Context) error { return commentController.Create(c) })
	api.PUT("/comments/:id", func(c echo.Context) error { return commentController.Update(c) })
	api.DELETE("/comments/:id", func(c echo.Context) error { return commentController.Delete(c) })

	api.GET("/rooms", func(c echo.Context) error { return roomController.Index(c) })
	api.GET("/rooms/:id", func(c echo.Context) error { return roomController.Show(c) })
	api.POST("/rooms", func(c echo.Context) error { return roomController.Create(c) })
	api.PUT("/rooms/:id", func(c echo.Context) error { return roomController.Update(c) })
	api.DELETE("/rooms/:id", func(c echo.Context) error { return roomController.Delete(c) })

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
