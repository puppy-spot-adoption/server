package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"puppyspot-backend/pkg/utils"
	"time"

	"github.com/go-mail/mail"
)

type SuccessMessage struct {
	Success    bool   `json:"success"`
	PaymentID  string `json:"paymentid"`
	Email      string `json:"email"`
	IsVerified bool   `json:"isverified"`
}

type FailureMessage struct {
	Success     bool   `json:"success"`
	ErrorNumber int    `json:"errornumber"`
	Message     string `json:"message"`
}

func HandlePaypalSumbit(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		fmt.Println(err)
		return
	} else {

		// Get form values
		paymentID := r.FormValue("paymentID")
		puppyID := r.FormValue("puppyID")
		price := r.FormValue("price")
		payerEmail := r.FormValue("payerEmail")
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
					sendPaypalMail2(payerEmail, paymentMethod, paymentID, puppyID, price, filePath, w)
				}

			}
		}
	}
}
func sendPaypalMail2(payerEmail, paymentMethod, paymentID, puppyID, price, filePath string, w http.ResponseWriter) {
	var emailAdd = os.Getenv("EMAIL")
	var emailPassword = os.Getenv("APP_PASSWORD")
	var emailHost = os.Getenv("EMAIL_HOST")
	var emialPort = 465
	// Create a new mailer
	m := mail.NewMessage()
	m.SetHeader("From", emailAdd)
	m.SetHeader("To", "info.puppyspotadoption@gmail.com")
	m.SetAddressHeader("Cc", emailAdd, "Puppy Spot")
	m.SetHeader("Subject", "THE BREAD IS HERE!!!")
	m.SetBody("text/html", "<h1>Hello White,</h1><br><p>someone made a "+paymentMethod+" purchase, <strong>Congratulations!!!</strong></p><br><p>details are as followed</p><br><ul><li>Payment ID of payer: "+paymentID+" </li><li>Payer Email: "+payerEmail+" </li><li>Amount to pay: "+price+" </li><li>Puppy ID: "+puppyID+" </li></ul>")

	// Attach the file
	m.Attach(filePath)

	// Send email
	d := mail.NewDialer(emailHost, emialPort, emailAdd, emailPassword)
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

func HandleCryptoSumbit(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	// Get form values
	paymentID := r.FormValue("paymentID")
	puppyID := r.FormValue("puppyID")
	price := r.FormValue("price")
	payerEmail := r.FormValue("payerEmail")
	// paymentMethod := r.FormValue("paymentMethod")
	blockChain := r.FormValue("blockChain")
	cryptoPrice := r.FormValue("cryptoPrice")

	sendCryptoMail(payerEmail, paymentID, puppyID, price, blockChain, cryptoPrice, w)
}
func sendCryptoMail(payerEmail, paymentID, puppyID, price, blockChain, cryptoPrice string, w http.ResponseWriter) {
	var emailAdd = os.Getenv("EMAIL")
	var emailPassword = os.Getenv("APP_PASSWORD")
	var emailHost = os.Getenv("EMAIL_HOST")
	var emialPort = 465
	// Create a new mailer
	m := mail.NewMessage()
	m.SetHeader("From", emailAdd)
	m.SetHeader("To", "info.puppyspotadoption@gmail.com")
	m.SetAddressHeader("Cc", emailAdd, "Puppy Spot")
	m.SetHeader("Subject", "THE BREAD IS HERE!!!")
	m.SetBody("text/html", "<h1>Hello White,</h1><br><p>someone made a crypto currency purchase, <strong>Congratulations!!!</strong></p><br><p>details are as followed</p><br><ul><li>Payment paymentID: "+paymentID+" </li><li>Payer Email: "+payerEmail+" </li><li>Payment ID: "+paymentID+" </li><li>Amount to pay: "+price+" </li><li>crypto price: "+cryptoPrice+" </li><li>Crypto Currency: "+blockChain+" </li><li>puppyID: "+puppyID+" </li></ul>")

	// Send email
	d := mail.NewDialer(emailHost, emialPort, emailAdd, emailPassword)
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

func HandleBankTransferSumbit(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		fmt.Println(err)
		return
	} else {

		// Get form values
		paymentID := r.FormValue("paymentID")
		puppyID := r.FormValue("puppyID")
		price := r.FormValue("price")
		payerEmail := r.FormValue("payerEmail")
		paymentMethod := r.FormValue("paymentMethod")
		accountName := r.FormValue("accountName")

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
					sendBankTransferMail(payerEmail, paymentMethod, paymentID, puppyID, price, filePath, accountName, w)
				}

			}
		}
	}
}
func sendBankTransferMail(payerEmail, paymentMethod, paymentID, puppyID, price, filePath, accountName string, w http.ResponseWriter) {
	var emailAdd = os.Getenv("EMAIL")
	var emailPassword = os.Getenv("APP_PASSWORD")
	var emailHost = os.Getenv("EMAIL_HOST")
	var emialPort = 465
	// Create a new mailer
	m := mail.NewMessage()
	m.SetHeader("From", emailAdd)
	m.SetHeader("To", "info.puppyspotadoption@gmail.com")
	m.SetAddressHeader("Cc", emailAdd, "Puppy Spot")
	m.SetHeader("Subject", "THE BREAD IS HERE!!!")
	m.SetBody("text/html", "<h1>Hello White,</h1><br><p>someone made a "+paymentMethod+" purchase, <strong>Congratulations!!!</strong></p><br><p>details are as followed</p><br><ul><li>Payment ID of payer: "+paymentID+" </li><li>Payer Email: "+payerEmail+" </li><li>Amount to pay: "+price+" </li><li>Puppy ID: "+puppyID+" </li><li>Payers account name: "+accountName+" </li></ul>")

	// Attach the file
	m.Attach(filePath)

	// Send email
	d := mail.NewDialer(emailHost, emialPort, emailAdd, emailPassword)
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


