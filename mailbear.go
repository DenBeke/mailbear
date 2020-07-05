package mailbear

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/mail.v2"
)

// MailBear will handle all the logic behind the forms.
type MailBear struct {
	config *Config
}

// formExists checks whether the given form exists in the config.
func (m *MailBear) formExists(formID string) bool {

	if m.config.Forms == nil {
		return false
	}

	for _, form := range m.config.Forms {
		if form.Key == formID {
			return true
		}
	}

	return false
}

// getFormByID gets the form with the given id from the config.
func (m *MailBear) getFormByID(formID string) *Form {

	for _, form := range m.config.Forms {
		if form.Key == formID {
			return form
		}
	}

	return nil
}

func buildMailBody(formSubmission *FormSubmission) string {
	const template = `
	<p>Hello,</p>
	<p>Someone has just submitted a new form on your website.</p>
	<p>Kind regards,<br>MailBear</p>
	<p><br></p>
	<p><b>Name:</b> %s</p>
	<p><b>Email:</b> %s</p>
	<p><b>Subject:</b> %s</p>
	<p><b>Content:</b><br><br>%s</p>
	`

	return fmt.Sprintf(template, formSubmission.Name, formSubmission.Email, formSubmission.Subject, strings.ReplaceAll(formSubmission.Content, "\n", "<br>"))
}

// SendMail sends a formsubmission to the receiver of the form.
func (m *MailBear) SendMail(formSubmission *FormSubmission) error {

	form := m.getFormByID(formSubmission.FormID)
	if form == nil {
		return fmt.Errorf("form does not exist")
	}

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	msg := mail.NewMessage()
	msg.SetHeader("From", m.config.Global.SMTP.FromEmail /*, m.config.Global.SMTP.FromName */)
	msg.SetHeader("To", form.ToEmail...)
	msg.SetAddressHeader("Reply-To", formSubmission.Email, formSubmission.Name)
	msg.SetHeader("Subject", fmt.Sprintf("New submission with subject: %s", formSubmission.Subject))
	msg.SetBody("text/html", buildMailBody(formSubmission))

	d := mail.NewDialer(
		m.config.Global.SMTP.Host,
		m.config.Global.SMTP.Port,
		m.config.Global.SMTP.User,
		m.config.Global.SMTP.Password,
	)

	if m.config.Global.SMTP.DisableTLS {
		d.StartTLSPolicy = mail.NoStartTLS
	}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(msg); err != nil {
		return errors.Wrap(err, "couldn't send the email")
	}

	/*
			err := smtp.SendMail(
				fmt.Sprintf("%s:%d", m.config.Global.SMTP.Host, m.config.Global.SMTP.Port),
				nil,
				m.config.Global.SMTP.FromEmail,
				[]string{form.ToEmail},
				[]byte(formSubmission.Content),
			)

		if err != nil {
			return errors.Wrap(err, "couldn't send the email")
		}
	*/

	return nil
}
