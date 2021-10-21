package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/xuri/excelize/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	f, err := excelize.OpenFile(filepath.Join(path, "emails.xlsx"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	for _, row := range rows {
		from := mail.NewEmail("crazyoptimist", "hey@crazyoptimist.net")
		subject := "Sending with SendGrid is Fun"
		to := mail.NewEmail("Me", row[0])
		plainTextContent := "and easy to do anywhere, even with Go"
		htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		response, err := client.Send(message)
		if err != nil {
			log.Fatal(err)
		}
		if response.StatusCode == 202 {
			log.Println("Successfully delivered to: ", row[0])
		} else {
			log.Println("Email delivery failed: ", response.Body)
		}
	}

}
