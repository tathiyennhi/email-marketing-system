package email

import (
	"encoding/json"
	"golang-email-marketing-system/model"
	"golang-email-marketing-system/util/date"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func StandardMessage(template string, user model.User) string {
	template = strings.Replace(template, "{{TITLE}}", user.Title, 1)
	template = strings.Replace(template, "{{FIRST_NAME}}", user.FirstName, 1)
	template = strings.Replace(template, "{{LAST_NAME}}", user.LastName, 1)
	template = strings.Replace(template, "{{TODAY}}", date.FormatDate(), 1)

	return template
}

func ConvertJSONtoStruct(path string) (model.Mail, error) {
	// data := &Mail{}

	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		return model.Mail{}, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var user model.Mail

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	if err := json.Unmarshal(byteValue, &user); err != nil {
		return model.Mail{}, err
	}
	// fmt.Println(user.From)
	// fmt.Println(user.Subject)
	// fmt.Println(user.MimeType)
	// fmt.Println(user.Body)
	return user, nil
}

func GetENV() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	email := os.Getenv("MAIL")
	password := os.Getenv("PASSWORD")

	return email, password
}
