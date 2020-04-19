package database

import (
	"github.com/watarun54/go-video-chat/api/domain"
)

type RoomRepository struct {
	SqlHandler
}

func (repo *RoomRepository) FindOne(r domain.Room) (room domain.Room, err error) {
	if err = repo.Debug().Take(&room, r.ID).Related(&room.Users, "Users").Error; err != nil {
		return
	}
	return
}

func (repo *RoomRepository) FindAll() (rooms domain.Rooms, err error) {
	if err = repo.Debug().Preload("Users").Find(&rooms).Error; err != nil {
		return
	}
	return
}

func (repo *RoomRepository) Store(r domain.Room) (room domain.Room, err error) {
	if err = repo.Debug().Create(&r).Error; err != nil {
		return
	}
	room = r
	return
}
