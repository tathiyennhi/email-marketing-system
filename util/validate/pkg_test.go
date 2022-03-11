package validate

import (
	"golang-email-marketing-system/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMail(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		mail     string
		expected bool
	}{
		{
			"",
			false,
		},
		{
			"tathiyennhi.test@gmail.com",
			true,
		},
		{
			"a@gmail.com",
			true,
		},
		{
			"@gmail.com",
			false,
		},
	}

	for _, test := range tests {
		assert.Equal(IsValidEmail(test.mail), test.expected)
	}
}

func TestValidateUser(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		data     model.User
		expected []string
	}{
		{
			model.User{},
			[]string{"", "", "", ""},
		},
		{
			model.User{
				"test", "test", "test", "test",
			},
			[]string{"test", "test", "test", "test"},
		},
		{
			model.User{
				"test", "test", "test", "test@gmail.com",
			},
			[]string{},
		},
	}

	for _, test := range tests {
		assert.Equal(ValidateUser(test.data), test.expected)
	}
}

func TestFiltUsers(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		data           []model.User
		hasUserValid   bool
		hasUserInvalid bool
	}{
		{
			[]model.User{},
			false,
			false,
		},
		{
			[]model.User{
				{"Mr", "John", "Smith", "ex@gmail.com"},
				{"Mrs", "Anna", "Smith", "anna@gmail.com"},
			},
			true,
			false,
		},
		{
			[]model.User{
				{"Mr", "John", "Smith", "ex@gmail.com"},
				{"Mrs", "Anna", "Smith", ""},
			},
			true,
			true,
		},
		{
			[]model.User{
				{"Mr", "John", "Smith", ""},
				{"Mrs", "Anna", "Smith", ""},
			},
			false,
			true,
		},
	}

	for _, test := range tests {
		userValids, userInvalids := FiltUsers(test.data)
		if test.hasUserInvalid {
			assert.NotEmpty(userInvalids)
		} else {
			assert.Empty(userInvalids)
		}

		if test.hasUserValid {
			assert.NotEmpty(userValids)
		} else {
			assert.Empty(userValids)
		}
	}
}
