package controllers

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"

	"github.com/watarun54/go-video-chat/api/domain"
	"github.com/watarun54/go-video-chat/api/interfaces/database"
	"github.com/watarun54/go-video-chat/api/usecase"
)

type jwtCustomClaims struct {
	UID  int    `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

func NewJWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: signingKey,
	}
}

type AuthController struct {
	Interactor usecase.UserInteractor
}

func NewAuthController(sqlHandler database.SqlHandler) *AuthController {
	return &AuthController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *AuthController) Login(c Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, _ := controller.Interactor.UserByEmail(u.Email)
	if user.ID == 0 || user.Password != u.Password {
		c.JSON(500, NewError(errors.New("invalid name or password")))
		return
	}

	claims := &jwtCustomClaims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return
	}

	c.JSON(200, map[string]string{
		"token": t,
	})
	return
}

func userIDFromToken(c Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	uid := claims.UID
	return uid
}
