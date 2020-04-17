package controllers

import (
	"strconv"

	_ "github.com/watarun54/go-video-chat/api/domain"
	"github.com/watarun54/go-video-chat/api/interfaces/database"
	"github.com/watarun54/go-video-chat/api/usecase"
)

type CommentController struct {
	Interactor usecase.CommentInteractor
}

func NewCommentController(sqlHandler database.SqlHandler) *CommentController {
	return &CommentController{
		Interactor: usecase.CommentInteractor{
			CommentRepository: &database.CommentRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *CommentController) Show(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	comment, err := controller.Interactor.CommentById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, comment)
	return
}

func (controller *CommentController) Index(c Context) (err error) {
	comments, err := controller.Interactor.Comments()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, comments)
	return
}
