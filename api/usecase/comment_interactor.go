package usecase

import "github.com/watarun54/go-video-chat/api/domain"

type CommentInteractor struct {
	CommentRepository CommentRepository
}

func (interactor *CommentInteractor) Comment(c domain.Comment) (comment domain.Comment, err error) {
	comment, err = interactor.CommentRepository.FindOne(c)
	return
}

func (interactor *CommentInteractor) Comments(c domain.Comment) (comments domain.Comments, err error) {
	comments, err = interactor.CommentRepository.FindAll(c)
	return
}

func (interactor *CommentInteractor) Add(c domain.Comment) (comment domain.Comment, err error) {
	comment, err = interactor.CommentRepository.Store(c)
	return
}
