package controllers

import (
	"strconv"

	"github.com/watarun54/go-video-chat/api/domain"
	"github.com/watarun54/go-video-chat/api/interfaces/database"
	"github.com/watarun54/go-video-chat/api/usecase"
)

type RoomController struct {
	Interactor usecase.RoomInteractor
}

func NewRoomController(sqlHandler database.SqlHandler) *RoomController {
	return &RoomController{
		Interactor: usecase.RoomInteractor{
			RoomRepository: &database.RoomRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RoomController) Show(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := domain.Room{
		ID: id,
	}
	room, err := controller.Interactor.Room(r)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, room)
	return
}

func (controller *RoomController) Index(c Context) (err error) {
	rooms, err := controller.Interactor.Rooms()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, rooms)
	return
}
