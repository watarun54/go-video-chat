package usecase

import "github.com/watarun54/go-video-chat/api/domain"

type (
	IRoomRepository interface {
		FindOne(r domain.Room) (domain.Room, error)
		FindAll() (domain.Rooms, error)
		Store(r domain.Room) (domain.Room, error)
	}

	RoomInteractor struct {
		RoomRepository IRoomRepository
	}
)

func (interactor *RoomInteractor) Room(r domain.Room) (room domain.Room, err error) {
	room, err = interactor.RoomRepository.FindOne(r)
	return
}

func (interactor *RoomInteractor) Rooms() (rooms domain.Rooms, err error) {
	rooms, err = interactor.RoomRepository.FindAll()
	return
}

func (interactor *RoomInteractor) Add(r domain.Room) (room domain.Room, err error) {
	room, err = interactor.RoomRepository.Store(r)
	return
}
