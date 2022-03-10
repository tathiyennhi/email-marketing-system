package validate

import (
	"golang-email-marketing-system/model"
	"golang-email-marketing-system/util/email"
	"net/mail"
)

func ValidateUser(user model.User) []string {
	if !IsValidEmail(user.Email) {
		return []string{user.Title, user.FirstName, user.LastName, user.Email}
	}

	return []string{}
}

func FiltUsers(users []model.User) ([]model.User, [][]string) {
	userMissingMailList := make([][]string, 0)
	userListValid := make([]model.User, 0)

	for _, user := range users {
		if userInfo := ValidateUser(user); len(userInfo) > 0 {
			userMissingMailList = append(userMissingMailList, userInfo)
			continue
		}

		userListValid = append(userListValid, user)
	}

	return userListValid, userMissingMailList
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func StandardUserJSON(users []model.User, mailTemplate model.Mail) []model.UserJSON {
	response := make([]model.UserJSON, 0, len(users))

	for _, item := range users {
		response = append(response, model.UserJSON{
			From:     mailTemplate.From,
			To:       item.Email,
			MimeType: mailTemplate.MimeType,
			Subject:  mailTemplate.Subject,
			Body:     email.StandardMessage(mailTemplate.Body, item),
		})
	}

	return response
}
