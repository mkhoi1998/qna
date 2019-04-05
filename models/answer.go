package models

import "time"

//Answer ..
type answer struct {
	ID         int
	QuestionID int
	Content    string
	Owner      string
	CreastedAt time.Time
}
