package email

import (
	"golang-email-marketing-system/model"
	"golang-email-marketing-system/util/date"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardMessage(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		template string
		user     model.User
		expected string
	}{
		{
			"Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
			model.User{},
			"Hi   , \nToday, " + date.FormatDate() + ", we would like to tell you that... Sincerely,\nThe Marketing Team",
		},
		{
			"Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
			model.User{
				Title:     "test",
				FirstName: "test",
				LastName:  "test",
			},
			"Hi test test test, \nToday, " + date.FormatDate() + ", we would like to tell you that... Sincerely,\nThe Marketing Team",
		},
		{
			"Hi {{FIRST_NAME}} {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
			model.User{
				Title:     "test",
				FirstName: "test",
				LastName:  "test",
			},
			"Hi test test, \nToday, " + date.FormatDate() + ", we would like to tell you that... Sincerely,\nThe Marketing Team",
		},
		{
			"Hi {{TITLE}} {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
			model.User{
				Title:     "test",
				FirstName: "test",
				LastName:  "test",
			},
			"Hi test test, \nToday, " + date.FormatDate() + ", we would like to tell you that... Sincerely,\nThe Marketing Team",
		},
		{
			"Hi {{TITLE}} {{FIRST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
			model.User{
				Title:     "test",
				FirstName: "test",
				LastName:  "test",
			},
			"Hi test test, \nToday, " + date.FormatDate() + ", we would like to tell you that... Sincerely,\nThe Marketing Team",
		},
		{
			"Hi, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
			model.User{
				Title:     "test",
				FirstName: "test",
				LastName:  "test",
			},
			"Hi, \nToday, " + date.FormatDate() + ", we would like to tell you that... Sincerely,\nThe Marketing Team",
		},
		{
			"Hi {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
			model.User{
				Title:     "test",
				FirstName: "test",
				LastName:  "test",
			},
			"Hi test, \nToday, " + date.FormatDate() + ", we would like to tell you that... Sincerely,\nThe Marketing Team",
		},
		//remove 2 fields cases
	}

	for _, test := range tests {
		assert.Equal(StandardMessage(test.template, test.user), test.expected)
	}
}

func TestConvertJSONtoStruct(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		path             string
		isConvertSuccess bool
	}{
		{
			"../../email_template.json",
			true,
		},
		{
			"email_template.json",
			false,
		},
		{
			".json",
			false,
		},
		{
			"",
			false,
		},
	}

	for _, test := range tests {
		data, err := ConvertJSONtoStruct(test.path)

		if test.isConvertSuccess {
			assert.NotEmpty(data)
			assert.NoError(err)
			continue
		}

		assert.Empty(data)
		assert.Error(err)
	}
}

func TestSendEmails(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		data      []model.UserJSON
		isSuccess bool
	}{
		{
			[]model.UserJSON{
				{
					"From",
					"To",
					"Subject",
					"MimeType",
					"Body",
				},
				{
					"The Marketing Team<marketing@example.com>",
					"tathiyennhi.test@gmail.com",
					"A new product is being launched soon...",
					"text/html",
					"Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
				},
			},
			false,
		},
		{
			[]model.UserJSON{
				{
					"The Marketing Team<marketing@example.com>",
					"abc@gmail.com",
					"A new product is being launched soon...",
					"text/html",
					"Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
				},
				{
					"The Marketing Team<marketing@example.com>",
					"tathiyennhi.test@gmail.com",
					"A new product is being launched soon...",
					"text/html",
					"Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
				},
			},
			false,
		},
	}

	for _, test := range tests {
		err := SendEmails(test.data)

		if test.isSuccess {
			assert.NoError(err)
			continue
		}

		assert.Error(err)
	}
}

func TestSendEmail(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		data      model.UserJSON
		isSuccess bool
	}{
		{
			model.UserJSON{"From", "To", "Subject", "MimeType", "Body"},
			false,
		},
		{
			model.UserJSON{
				"The Marketing Team<marketing@example.com>",
				"tathiyennhi.test@gmail.com",
				"A new product is being launched soon...",
				"text/html",
				"Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
			},
			true,
		},
		{
			model.UserJSON{"From", "", "Subject", "MimeType", "Body"},
			false,
		},
	}

	for _, test := range tests {
		err := SendEmail(test.data)

		if test.isSuccess {
			assert.NoError(err)
			continue
		}

		assert.Error(err)
	}
}
