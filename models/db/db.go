package db

import "github.com/qna/models"

// DB ..
type DB interface {
	Connect(connectionStr string) error
	Save(q models.Question) (int, error)
}
