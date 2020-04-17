package usecase

import "github.com/watarun54/go-video-chat/api/domain"

type CommentRepository interface {
	FindById(id int) (domain.Comment, error)
	FindAll() (domain.Comments, error)
}
