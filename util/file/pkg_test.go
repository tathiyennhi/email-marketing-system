package file

import (
	"golang-email-marketing-system/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetColumnName(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		columnName []string
		data       map[string]int
		IsSuccess  bool
	}{
		{
			[]string{"Test", "Test1", "Test2", "Test3"},
			map[string]int{"TEST": 0, "TEST1": 1, "TEST2": 2, "TEST3": 3},
			true,
		},
	}

	for _, test := range tests {
		assert.Equal(GetColumnName(test.columnName), test.data)
	}
}

func TestGetInfo(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		data       map[string]int
		userDetail []string
		key        string
		expected   string
	}{
		{
			map[string]int{"FIRST_NAME,": 1, "TITLE": 0, "LAST_NAME": 2, "EMAIL": 3},
			[]string{"TITLE", "FIRST_NAME", "LAST_NAME", "EMAIL"},
			"TITLE",
			"TITLE",
		},
		{
			map[string]int{"FIRST_NAME,": 1, "TITLE": 0, "LAST_NAME": 2, "EMAIL": 3},
			[]string{"TITLE", "FIRST_NAME", "LAST_NAME", "EMAIL"},
			"abc",
			"",
		},
		{
			map[string]int{"FIRST_NAME,": 1, "TITLE": 0, "LAST_NAME": 2, "EMAIL": 3},
			[]string{"TITLE", "FIRST_NAME", "LAST_NAME"},
			"EMAIL",
			"",
		},
	}

	for _, test := range tests {
		assert.Equal(GetInfo(test.data, test.key, test.userDetail), test.expected)
	}
}

func TestReadCSVFile(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		path                string
		IsCreateSuccessfull bool
	}{
		{
			"../../customers.csv",
			true,
		},
		{
			"../../customers.json",
			false,
		},
		{
			"customers.csv",
			false,
		},
		{
			".csv",
			false,
		},
	}

	for _, test := range tests {

		users, err := ReadCSVFile(test.path)

		if test.IsCreateSuccessfull {
			assert.NotEmpty(users)
			assert.NoError(err)
			continue
		}

		assert.Empty(users)
		assert.Error(err)
	}
}

func TestWriteCSV(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		path                string
		data                [][]string
		IsCreateSuccessfull bool
	}{
		{
			"",
			[][]string{},
			false,
		},
		{
			"abc.csv",
			[][]string{},
			true,
		},
		{
			"abc.json",
			[][]string{
				{"asdsa", "sadsada"},
			},
			false,
		},
	}

	for _, test := range tests {
		err := WriteCSV(test.path, test.data)

		if test.IsCreateSuccessfull {
			assert.NoError(err)
			continue
		}

		assert.Error(err)
	}
}

func TestWriteJSON(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		path                string
		data                []model.UserJSON
		IsCreateSuccessfull bool
	}{
		{
			"",
			[]model.UserJSON{},
			false,
		},
		{
			"test-json.csv",
			[]model.UserJSON{},
			false,
		},
		{
			"test-json.json",
			[]model.UserJSON{
				{"test@gmail.com", "test@gmail.com", "Test Subject", "text/html", "Test Content"},
			},
			true,
		},
	}

	for _, test := range tests {
		err := WriteJSON(test.data, test.path)

		if test.IsCreateSuccessfull {
			assert.NoError(err)
			continue
		}

		assert.Error(err)
	}
}
