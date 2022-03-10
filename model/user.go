package model

type User struct {
	Title     string
	FirstName string
	LastName  string
	Email     string
}

type UserJSON struct {
	From     string
	To       string
	Subject  string
	MimeType string
	Body     string
}
