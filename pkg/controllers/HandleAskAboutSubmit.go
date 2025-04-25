
package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"puppyspot-backend/pkg/utils"
	"time"

	"github.com/go-mail/mail"
)


func HandleAskAboutMail (w http.ResponseWriter, r *http.Request){
	utils.EnableCors(w, r)

	// Get form values
	firstName := r.FormValue("firstName")
    lastName := r.FormValue("lastName")
    emailAddress := r.FormValue("emailAddress")
    phone := r.FormValue("phone")
    state := r.FormValue("state")
    text := r.FormValue("text")
    puppyID := r.FormValue("puppyID")

    var emailAdd = os.Getenv("EMAIL")
	var emailPassword = os.Getenv("APP_PASSWORD")
	var emailHost = os.Getenv("EMAIL_HOST")
	var emialPort = 587

	if firstName != "" {
		// Create a new mailer
		m := mail.NewMessage()
		m.SetHeader("From", emailAdd)
		m.SetHeader("To", "info.puppyspotadoption@gmail.com")
		m.SetAddressHeader("Cc", emailAdd, "Puppy Spot")
		m.SetHeader("Subject", "WANT TO KNOW MORE")
		m.SetBody("text/html", "<h1>Hello White,</h1><br><p>someone wants to know more about this puppy: "+puppyID+", <strong>Congratulations!!!</strong></p><br><p>details of asker, are as followed</p><br><ul><li>First Name : "+firstName+" </li><li>Last Name: "+lastName+" </li><li>Email: "+emailAddress+" </li><li>phone: "+phone+" </li><li>state: "+state+" </li><li>text: "+text+" </li></ul>")

		// Send email
		d := mail.NewDialer(emailHost, emialPort, emailAdd, emailPassword)
		d.Timeout = 120 * time.Second
		d.StartTLSPolicy = mail.MandatoryStartTLS

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

	

}
