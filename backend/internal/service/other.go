// для остальных функций
package service

import (
	"crypto/sha256"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"text/template"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/spf13/viper"
)

const (
	emailConfirmationTTL = 24 * time.Hour
	minCode              = 100000
	maxCode              = 999999
	baseDelay            = 5 * time.Second // Начальное время задержки
	delayMultiplier      = 2               // Множитель для экспоненциальной задержки
)

func generatePasswordHash(password string) string {
	hash := sha256.New()

	var passwordBytes = []byte(password)
	passwordBytes = append(passwordBytes, os.Getenv("SALT")...)

	hash.Write(passwordBytes)

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func sendEmailWithCode(login, templateName string, code int) error {
	// Sender data
	from := viper.GetString("smtp.from")
	// password := "drYk8ykHR399"
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := viper.GetString("smtp.smtpHost")
	smtpPort := viper.GetInt("smtp.smtpPort")

	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", login)
	message.SetHeader("Subject", "Verification code for our app!")

	t, err := template.ParseFiles(templateName)
	if err != nil {
		return err
	}
	message.AddAlternativeWriter("text/html", func(w io.Writer) error {
		return t.Execute(w, struct {
			Message string
		}{
			Message: fmt.Sprintf("%d", code),
		})
	})

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	if err := d.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
func generateCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxCode-minCode+1) + minCode
}

func exponentialBackoff(attempt int) time.Duration {
	// Рассчитываем время задержки по экспоненциальной формуле
	delay := float64(baseDelay) * math.Pow(delayMultiplier, float64(attempt-1))
	return time.Duration(delay)
}
