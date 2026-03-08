package mailer

import (
	"bytes"
	"embed"
	"text/template"
	"time"

	"github.com/go-mail/mail/v2"
)

// `templateFS` is an embedded filesystem that will store email templates.
// The templates are used to populate the email subject and body dynamically.
var templateFS embed.FS

// `Mailer` struct encapsulates the `Dialer` used for sending emails and the `sender` email address.
// This abstraction makes it easy to configure and reuse for sending multiple emails.
type Mailer struct {
	dialer *mail.Dialer
	sender string
}

// `New` creates a new instance of `Mailer`. It takes SMTP configuration parameters (host, port, username, password)
// and the `sender` email address. The `Dialer` is set with a timeout of 5 seconds to avoid long waits for failed connections.
func New(host string, port int, username, password, sender string) Mailer {
	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = 5 * time.Second

	return Mailer{
		dialer: dialer,
		sender: sender,
	}
}

// `Send` is the core method for sending an email. It takes in the recipient email address, a template file name, and data to be populated in the template.
// It loads and parses the template, constructs the email message, and attempts to send it using the `Dialer`. If sending fails, it retries up to 3 times.
func (m Mailer) Send(recipient, templateFile string, data interface{}) error {
	// Load and parse the email template from the embedded filesystem.
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	// Buffer to store the parsed email subject.
	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	// Buffer to store the parsed plain text email body.
	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return err
	}

	// Buffer to store the parsed HTML email body.
	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return err
	}

	// Construct the email message.
	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	// Attempt to send the email. If it fails, retry up to 3 times with a delay.
	err = m.dialer.DialAndSend(msg)
	if err != nil {
		// Retry sending email up to 3 times if the initial attempt fails.
		for i := 1; i <= 3; i++ {
			time.Sleep(500 * time.Millisecond) // Wait 500ms before retrying.
			err = m.dialer.DialAndSend(msg)
			if err == nil {
				return nil
			}
		}
		return err
	}

	return nil
}
