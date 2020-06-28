package mailbear

import (
	"fmt"

	"github.com/badoux/checkmail"
)

// FormSubmission represent a submitted form
type FormSubmission struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Subject string `json:"subject" form:"subject"`
	Content string `json:"content" form:"content"`
	FormID  string `json:"-"`
}

// Validate validates all fields of a form submission.
func (f *FormSubmission) Validate() error {
	if f.Name == "" {
		// name is optional
	}
	if f.Email == "" {
		return fmt.Errorf("field 'email'cannot be empty")
	}
	if f.Subject == "" {
		return fmt.Errorf("field 'subject' cannot be empty")
	}
	if f.Content == "" {
		return fmt.Errorf("field 'content' cannot be empty")
	}

	// Validate ID, although it should always be set manually
	if f.Content == "" {
		return fmt.Errorf("id cannot be empty")
	}

	// Validate format of email
	err := checkmail.ValidateFormat(f.Email)
	if err != nil {
		return fmt.Errorf("invalid email address: %v", err)
	}

	return nil
}
