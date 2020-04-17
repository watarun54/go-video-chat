package database

import (
	"github.com/watarun54/go-video-chat/api/domain"
)

type CommentRepository struct {
	SqlHandler
}

func (repo *CommentRepository) FindById(id int) (comment domain.Comment, err error) {
	if err = repo.Find(&comment, id).Error; err != nil {
		return
	}
	return
}

func (repo *CommentRepository) FindAll() (comments domain.Comments, err error) {
	if err = repo.Find(&comments).Error; err != nil {
		return
	}
	return
}
