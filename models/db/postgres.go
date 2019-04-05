package db

import "github.com/qna/models"

// Postgres ..
type Postgres struct {
}

// Connect ..
func (p Postgres) Connect(connectionStr string) error {
	return nil
}

//Save ..
func (p Postgres) Save(q models.Question) (int, error) {
	return 0, nil
}
