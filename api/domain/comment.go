package domain

type Comments []Comment

type Comment struct {
	ID     int
	Text   string
	UserID int
}
