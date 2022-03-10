package file

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"golang-email-marketing-system/constant"
	"golang-email-marketing-system/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetColumnName(columnNames []string) map[string]int {
	keys := make(map[string]int)
	for idx, item := range columnNames {
		upperCase := strings.ToUpper(item)
		_, ok := keys[upperCase]
		if !ok {
			keys[upperCase] = idx

		}
	}
	return keys

}

func GetInfo(keys map[string]int, key string, userDetail []string) string {
	value, ok := keys[key]
	if !ok {
		return ""
	}
	if len(userDetail)-1 < value {
		return ""
	}
	return userDetail[value]

}

func ReadCSVFile(path string) ([]model.User, error) {
	f, err := os.Open(path)

	if err != nil {
		return []model.User{}, err
	}

	fmt.Print()

	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		return []model.User{}, err
	}

	fmt.Print(data)
	users := make([]model.User, 0)
	keys := make(map[string]int)

	for index, item := range data {
		fmt.Println(item)
		//check title column name
		if index == 0 {
			keys = GetColumnName(item)
			continue
		}
		user := model.User{}

		user.Title = GetInfo(keys, constant.Title, item)
		user.FirstName = GetInfo(keys, constant.FirstName, item)
		user.LastName = GetInfo(keys, constant.LastName, item)
		user.Email = GetInfo(keys, constant.Email, item)

		users = append(users, user)
	}
	fmt.Println(keys)

	fmt.Println(users)
	return users, nil
}

func WriteCSV(path string, data [][]string) error {
	if filepath.Ext(path) != ".csv" {
		return errors.New("Wrong extension!")
	}

	csvFile, err := os.Create(path)

	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}

	csvwriter.Flush()
	csvFile.Close()
	return nil
}

func WriteJSON(users []model.UserJSON, path string) error {
	if filepath.Ext(path) != ".json" {
		return errors.New("Wrong extension!")
	}

	file, err := json.MarshalIndent(users, "", " ")

	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, file, 0644)
}
