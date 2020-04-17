package usecase

import "github.com/watarun54/go-video-chat/api/domain"

type CommentRepository interface {
	FindOne(c domain.Comment) (domain.Comment, error)
	FindAll(c domain.Comment) (domain.Comments, error)
	Store(c domain.Comment) (domain.Comment, error)
	Update(c domain.Comment) (domain.Comment, error)
	DeleteById(c domain.Comment) error
}
