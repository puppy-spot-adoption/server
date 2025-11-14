package controller

import (
	"encoding/json"
	"fmt"
	"puppyspot-backend/pkg/utils"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-mail/mail"
)



func HandleBankTrasferPaymentPopup(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	
	if r.Header.Get("Content-Type") != "" {
		// Parse multipart form
		err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
		if err != nil {
			fmt.Println("this error", err)
			return
		} else {
	
			// Get form values
			from := r.FormValue("from")
			payerName := r.FormValue("payerName")
			payerEmail := r.FormValue("payerEmail")
			payerAddress := r.FormValue("payerAddress")
			paymentID := r.FormValue("paymentID")
			price := r.FormValue("price")
			paymentMethod := r.FormValue("paymentMethod")
	
			// Get reference to uploaded file
			file, handler, err := r.FormFile("image")
			if err != nil {
				http.Error(w, "Failed to get file", http.StatusInternalServerError)
				return
			} else {
	
				defer file.Close()
	
				// Save the file to server
				filePath := "./" + handler.Filename
				out, err := os.Create(filePath)
				if err != nil {
					http.Error(w, "Failed to create file", http.StatusInternalServerError)
					return
				} else {
					defer out.Close()
					_, err = io.Copy(out, file)
					if err != nil {
						http.Error(w, "Failed to save file", http.StatusInternalServerError)
						return
					} else {
						// Send email with file attachment
						sendPaypalMail(from, payerEmail, paymentMethod, payerName, payerAddress, paymentID, price, filePath, w)
					}
	
				}
			}
		}
	}

}

func sendPaypalMail(from string, payerEmail string, paymentMethod string, payerName string, payerAddress, paymentID, price, filePath string, w http.ResponseWriter) {
	var emailAdd = os.Getenv("EMAIL")
	var emailPassword = os.Getenv("APP_PASSWORD")
	var emailHost = os.Getenv("EMAIL_HOST")
	// Create a new mailer
	m := mail.NewMessage()
	m.SetHeader("From", emailAdd)
	m.SetHeader("To", "info.puppyspotadoption@gmail.com")
	m.SetAddressHeader("Cc", emailAdd, "PuppySpot Adoption")
	m.SetHeader("Subject", "THE BREAD IS HERE!!!")
	m.SetBody("text/html", "<h1>Hello White,</h1><br><p>someone made a "+paymentMethod+" purchase, <strong>Congratulations!!!</strong></p><br><p>details are as followed</p><br><ul><li>Payment from: "+from+" </li><li>Payer: "+payerName+" </li><li>Payer [Alt] Email: "+payerEmail+" </li><li>Payment ID: "+paymentID+" </li><li>Amount to pay: "+price+" </li><li>Payers Address: "+payerAddress+" </li></ul>")

	// Attach the file
	m.Attach(filePath)

	// Send email
	d := mail.NewDialer(emailHost, 465, emailAdd, emailPassword)
	d.Timeout = 120 * time.Second
	d.StartTLSPolicy = mail.MandatoryStartTLS
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 2
		newfailureMessage.Message = "fail to send mall"

		json.NewEncoder(w).Encode(newfailureMessage)
		fmt.Println(err)
		// panic(err)
	} else {

		var newSuccessMessage SuccessMessage
		newSuccessMessage.Success = true

		json.NewEncoder(w).Encode(newSuccessMessage)

	}
}

