package handlers

import (
	"fmt"
	"hollyways/models"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func sendMail(status string, transaction models.Transaction) {
	var CONFIG_SMTP_HOST = os.Getenv("SMTP_HOST")
	var CONFIG_SMTP_PORT, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	var CONFIG_SENDER_EMAIL = os.Getenv("SENDER_EMAIL")
	var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
	var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

	if status != transaction.Status && (status == "success") {
		transactionId := strconv.Itoa(int(transaction.ID))
		donation := strconv.Itoa(transaction.Donation)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_EMAIL)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Thanks For Donation in Hollyways")
		mailer.SetBody("text/html", fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8" />
				<meta http-equiv="X-UA-Compatible" content="IE=edge" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<title>Document</title>
			</head>
			<body>
				<section>
					<h2> Dear, %s </h2>
					<h3> thanks for your suport donation %s, Rp. %s</h3>
					<h4> your transaction %s have received by my team </h4>
				</section>
			</body>
		</html>
		`, transaction.User.FullName, transaction.Project.ProjectName, donation, transactionId))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("mail send to" + transaction.User.Password)
	}
}
