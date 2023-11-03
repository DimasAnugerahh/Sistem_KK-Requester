package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func EmailSender(email_tujuan string, link_gambar string) {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "dimasanugrah883@gmail.com")
	m.SetHeader("To", email_tujuan)
	m.SetHeader("Subject", "Layanan Kartu keluarga online")
	m.SetBody("text/plain", link_gambar)
	htmlContent := `<html>
			<body>
				<h3>Halo, terima kasih telah menggunakan layanan ini</h3>
				<p>Berikut adalah kartu keluarga anda</p>
				
				<img src="` + link_gambar + `" alt="External Image">
			</body>
		</html>`
	m.SetBody("text/html", htmlContent)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASS"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
