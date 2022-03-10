### Email Marketing System Installation
###### go version 1.17
###### $go mod tidy

##### Setup Email Sender
###### Create .env file in main directory

```
MAIL=<your-email@example.com>
PASSWORD=<email-password>
```
##### Argument List 
```
$go run main.go email_template.json customers.csv customer.json errors.csv
```
##### Unit Test cmd

```
$go test ./... -v
```
