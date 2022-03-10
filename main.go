package main

import (
	"golang-email-marketing-system/constant"
	"golang-email-marketing-system/util/email"
	"golang-email-marketing-system/util/file"
	"golang-email-marketing-system/util/validate"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		log.Panicln("Missing argument. Please check again.")
	}

	emailTemplateFilePath := os.Args[1]
	customersCsvFilePath := os.Args[2]
	customerJsonFilePath := os.Args[3]
	errorsFilePath := os.Args[4]

	mailTemplate, err := email.ConvertJSONtoStruct(emailTemplateFilePath)

	if err != nil {
		log.Panic(err)
	}

	users, err := file.ReadCSVFile(customersCsvFilePath)

	if err != nil {
		log.Panic(err)
	}

	emptyData := [][]string{
		{constant.Title, constant.FirstName, constant.LastName, constant.Email},
	}
	userListValid, userMissingMailList := validate.FiltUsers(users)
	emptyData = append(emptyData, userMissingMailList...)
	file.WriteCSV(errorsFilePath, emptyData)

	userValids := validate.StandardUserJSON(userListValid, mailTemplate)
	err = file.WriteJSON(userValids, customerJsonFilePath)

	if err != nil {
		log.Panic(err)
	}

	if err := email.SendEmails(userValids); err != nil {
		log.Panic(err)
	}
}
