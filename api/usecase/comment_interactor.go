package usecase

import "github.com/watarun54/go-video-chat/api/domain"

type CommentInteractor struct {
	CommentRepository CommentRepository
}

func (interactor *CommentInteractor) CommentById(id int) (comment domain.Comment, err error) {
	comment, err = interactor.CommentRepository.FindById(id)
	return
}

func (interactor *CommentInteractor) Comments() (comments domain.Comments, err error) {
	comments, err = interactor.CommentRepository.FindAll()
	return
}
