package mailbear

import (
	"fmt"
	"net/url"

	"github.com/badoux/checkmail"
)

// Config represents the complete app config of Mail Bear.
//
// Global contains all the global configurations
// Forms is a map of individual forms
type Config struct {
	Global struct {
		SMTP struct {
			Host       string `yaml:"host"`
			Port       int    `yaml:"port"`
			User       string `yaml:"user"`
			Password   string `yaml:"password"`
			DisableTLS bool   `yaml:"disable_tls"`
			FromEmail  string `yaml:"from_email"`
			FromName   string `yaml:"from_name"`
		} `yaml:"smtp"`
		HTTP struct {
			Address string `yaml:"address"`
		} `yaml:"http"`
	} `yaml:"global"`
	Forms map[string]*Form `yaml:"forms"`
}

// Validate checks for all the required values in config
func (config *Config) Validate() error {

	if config.Global.SMTP.Host == "" {
		return fmt.Errorf("smpt host must be set")
	}
	if config.Global.SMTP.Port == 0 {
		return fmt.Errorf("smpt port must be set")
	}
	if config.Global.SMTP.FromEmail == "" {
		return fmt.Errorf("smtp from email must be set")
	}
	if config.Global.HTTP.Address == "" {
		return fmt.Errorf("http address must be set")
	}

	// Check for nil pointer on forms
	if config.Forms == nil {
		return fmt.Errorf("forms cannot be null")
	}
	if len(config.Forms) == 0 {
		return fmt.Errorf("expected to have at least one form")
	}

	// Check for duplicate form keys
	formKeys := map[string]bool{}

	// Validate forms
	for _, form := range config.Forms {
		err := form.Validate()
		if err != nil {
			return fmt.Errorf("invalid form with key %q: %v", form.Key, err)
		}
		// Check for duplicate form keys
		if _, ok := formKeys[form.Key]; ok {
			return fmt.Errorf("there exist already a form with the key: %q", form.Key)
		}
		formKeys[form.Key] = true
	}

	return nil
}

// Form represents the form in the config
type Form struct {
	Key            string   `yaml:"key"`
	AllowedDomains []string `yaml:"allowed_domains"`
	ToEmail        []string `yaml:"to_email"`
}

// Validate checks for all the required values in a form
func (form *Form) Validate() error {
	if form.Key == "" {
		return fmt.Errorf("'key' of form cannot be empty")
	}
	if len(form.AllowedDomains) == 0 {
		return fmt.Errorf("form should have at least one allowed domain in 'allowed_domains'")
	}
	if len(form.ToEmail) == 0 {
		return fmt.Errorf("form should have at least one recipient email address in 'to_email'")
	}
	for _, email := range form.ToEmail {
		err := checkmail.ValidateFormat(email)
		if err != nil {
			return fmt.Errorf("invalid email address %q in 'to_email", email)
		}
	}
	return nil
}

// OriginDomainAllowed checks whether the given origin is allowed to access the form.
func (form *Form) OriginDomainAllowed(origin string) bool {

	if origin == "" {
		return false
	}

	u, err := url.Parse(origin)
	if err != nil {
		return false
	}
	originDomain := u.Host

	for _, allowedDomain := range form.AllowedDomains {
		if allowedDomain == "*" {
			return true
		}
		if allowedDomain == originDomain {
			return true
		}
	}
	return false
}
