package database

import (
	"github.com/watarun54/go-video-chat/api/domain"
)

type CommentRepository struct {
	SqlHandler
}

func (repo *CommentRepository) FindOne(c domain.Comment) (comment domain.Comment, err error) {
	if err = repo.Debug().Where(&c).Take(&comment).Error; err != nil {
		return
	}
	return
}

func (repo *CommentRepository) FindAll(c domain.Comment) (comments domain.Comments, err error) {
	if err = repo.Debug().Where(&c).Find(&comments).Error; err != nil {
		return
	}
	return
}

func (repo *CommentRepository) Store(c domain.Comment) (comment domain.Comment, err error) {
	if err = repo.Debug().Create(&c).Error; err != nil {
		return
	}
	comment = c
	return
}
