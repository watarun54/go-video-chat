package usecase

import "github.com/watarun54/go-video-chat/api/domain"

type (
	IUserRepository interface {
		FindById(id int) (domain.User, error)
		FindByIds(ids []int) (domain.Users, error)
		FindByEmail(email string) (domain.User, error)
		FindAll() (domain.Users, error)
		Store(domain.User) (domain.User, error)
		Update(domain.User) (domain.User, error)
		DeleteById(domain.User) error
	}

	UserInteractor struct {
		UserRepository IUserRepository
	}
)

func (interactor *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return
}

func (interactor *UserInteractor) UserByEmail(email string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByEmail(email)
	return
}

func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

func (interactor *UserInteractor) DeleteById(u domain.User) (err error) {
	err = interactor.UserRepository.DeleteById(u)
	return
}
