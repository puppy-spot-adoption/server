package controller

import (
	"fmt"
	"net/http"
	"os"
	"puppyspot-backend/pkg/utils"
	"time"

	"github.com/go-mail/mail"
)




func HandleNotification(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		fmt.Println(err)
		return
	} else {
		// Get form values
		documentID := r.FormValue("documentID")
		message := r.FormValue("message")


		var emailAdd = os.Getenv("EMAIL")
		var emailPassword = os.Getenv("APP_PASSWORD")
		var emailHost = os.Getenv("EMAIL_HOST")

		// Create a new mailer
		m := mail.NewMessage()
		m.SetHeader("From", emailAdd)
		m.SetHeader("To", "info.puppyspotadoption@gmail.com")
		m.SetAddressHeader("Cc", emailAdd, "Puppy Spot")
		m.SetHeader("Subject", "Notification!!!")

		m.SetBody("text/html", "Someone just triggered something: <br/> <p>ID: "+documentID+"</p> <br/> <p>Message: "+message+"</p>")
	
		
		// Send email
		d := mail.NewDialer(emailHost, 587, emailAdd, emailPassword)
		d.Timeout = 120 * time.Second
		d.StartTLSPolicy = mail.MandatoryStartTLS	


		// Attempt to send email
		if err := d.DialAndSend(m); err != nil {
			http.Error(w, "Failed to send email", http.StatusInternalServerError)
			fmt.Println("Error sending email:", err)
			return
		}

		// Send a success response to the client
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Notification sent successfully"))	

	}
}


func HandleUserNotificationEmail(w http.ResponseWriter, r *http.Request){
		utils.EnableCors(w, r)
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		fmt.Println(err)
		return
	} else {
		// Get form values
		documentID := r.FormValue("documentID")
		email := r.FormValue("email")
		puppyName := r.FormValue("puppyName")
		breed := r.FormValue("breed")
		

		var emailAdd = os.Getenv("EMAIL")
		var emailPassword = os.Getenv("APP_PASSWORD")
		var emailHost = os.Getenv("EMAIL_HOST")

		// Create a new mailer
		m := mail.NewMessage()
		m.SetHeader("From", emailAdd)
		m.SetHeader("To", email)
		m.SetAddressHeader("Cc", emailAdd, "Puppy Spot Adoption")
		m.SetHeader("Subject", "Your Puppy Adoption Application is Being Reviewed!")

		trackingURL := "https://puppyspotadoption.shop/shop/puppy-tracker/" + documentID
		emailMessage := GenerateEmailTemplate(puppyName, breed, documentID, trackingURL)
		m.SetBody("text/html", emailMessage)
		
		// Send email
		d := mail.NewDialer(emailHost, 587, emailAdd, emailPassword)
		d.Timeout = 120 * time.Second
		d.StartTLSPolicy = mail.MandatoryStartTLS	


		// Attempt to send email
		if err := d.DialAndSend(m); err != nil {
			http.Error(w, "Failed to send email", http.StatusInternalServerError)
			fmt.Println("Error sending email:", err)
			return
		}

		// Send a success response to the client
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Notification sent successfully"))	

	}

}






func GenerateEmailTemplate(puppyName, breed, trackingID, trackingURL string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Puppy Adoption Application</title>
</head>
<body style="font-family: Arial, sans-serif; margin: 0; padding: 0; background-color: #f5f5f5;">
    <table width="100%%" cellpadding="0" cellspacing="0" style="max-width: 600px; margin: 20px auto; background: #fff; padding: 20px; border-radius: 8px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
        <tr>
            <td align="center" style="margin-bottom: 1rem;">
                <img src="https://mail.puppyspot.com/hs-fs/hubfs/Logo_sticker_ps.png?width=420&upscale=true&name=Logo_sticker_ps.png" alt="Puppy Spot Adoption Logo" width="200">
            </td>
        </tr>
        <tr>
            <td style="border-top: 1px solid #219653; margin: 20px 0;">&nbsp;</td>
        </tr>
        <tr>
            <td align="center">
                <h2>Your Puppy Adoption Application is Being Reviewed!</h2>
            </td>
        </tr>
        <tr>
            <td>
                <p style="font-size: 1.05rem;">Thank you for applying to adopt %s! 🐾 We’re excited to process your application and will review it shortly.</p>
                <h3><strong>Application Details:</strong></h3>
                <ul>
                    <li style="font-size: 1.05rem"><strong>Puppy Name:</strong> %s</li>
                    <li style="font-size: 1.05rem"><strong>Breed:</strong> %s</li>
                    <li style="font-size: 1.05rem"><strong>Tracking ID:</strong> %s</li>
                    <li style="font-size: 1.05rem"><strong>Application Status:</strong> Under Review</li>
                </ul>
                <p style="font-size: 1.05rem;">You can track your application status anytime using your <strong>Tracking ID</strong> at:<br>🔗 <a href="%s">Tracking Portal</a></p>
                <p style="font-size: 1.05rem;">We will let you know as soon as your application is approved or if we need any more information.</p>
            </td>
        </tr>
        <tr>
            <td style="border-top: 1px solid #219653; margin: 20px 0;">&nbsp;</td>
        </tr>
        <tr>
            <td align="center">
                <h2>Want to see even more adorable puppies?</h2>
                <p style="font-size: 1.05rem;">There's no shortage of options here at PuppySpot. Dive into our diverse selection of other breeds to find your perfect puppy!</p>
                <img src="https://ci3.googleusercontent.com/meips/ADKq_NYEKEY4-fJkw_WsvfPegDQPrcYLUZa3c5nwfiWOCLjmzIk3rnqNuS0qGP0rgxcWAgVAz5_g8v2Uhx36UtFau0CRMaE12tyfloMQMoN02RuWX1y7kNW-WS3qisHnO9w2pyzj3kIHuls3F_z2hF4qTHMAfsKuJaw8wHrv_sLkqz-cx_2ptZkZxKc=s0-d-e1-ft#https://mail.puppyspot.com/hs-fs/hubfs/Most%%20pop%%20breeds.png?width=640&upscale=true&name=Most%%20pop%%20breeds.png" width="250">
            </td>
        </tr>
        <tr>
            <td align="center">
                <a href="https://puppyspotadoption.shop/puppies-for-sale/" style="display: block; background-color: #27ae60; color: #ffffff; font-weight: bold; padding: 10px; text-align: center; text-decoration: none; border-radius: 32px; width: 90%%; margin: 10px auto; cursor: pointer;">See all available pups</a>
            </td>
        </tr>
        <tr>
            <td align="center" style="padding: 20px; background: #f1f1f1; border-top: 1px solid #ddd; margin-bottom: 1rem">
                <img src="https://mail.puppyspot.com/hs-fs/hubfs/ps_logo_outline.png?width=70&upscale=true&name=ps_logo_outline.png" alt="Puppy Spot Adoption Logo" width="40">
                <p style="font-size: 1.05rem;"><strong>OUR PERFECT PUPPY PROMISE</strong></p>
                <p style="font-size: 1.05rem;">We promise to do everything we can to provide you with your perfect puppy and ensure your experience leaves you with a big smile and a warm heart.</p>
            </td>
        </tr>
        <tr>
            <td align="center" style="border-top: 2px solid #ddd; padding: 10px;">
                <a href="https://wa.me/15023820019" style="margin: 0 0.5rem; cursor: pointer">Contact a Puppy Concierge: +1 (602) 382-0019</a>
            </td>
        </tr>
        <tr>
            <td align="center" style="color: gray; font-size: 14px; padding: 10px;">
                <p style="font-size: 1.05rem;">Use of the PuppySpotAdoption service and website is subject to our</p>
                <p style="font-size: 1.05rem;"><a href="https://puppyspotadoption.shop/privacy" style="cursor: pointer;">Privacy Policy</a> and <a href="https://puppyspotadoption.shop/terms-of-use" style="cursor: pointer;">Terms of Service</a></p>
                <p style="font-size: 1.05rem;">PuppySpot &copy; 2025, All rights reserved.</p>
                <p style="font-size: 1.05rem;">PO Box 239, Nutley, New Jersey, 07110</p>
            </td>
        </tr>
    </table>
</body>
</html>`, puppyName, puppyName, breed, trackingID, trackingURL)
}


