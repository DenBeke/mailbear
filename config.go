package mailbear

import (
	"net/url"
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
	Forms map[string]Form `yaml:"forms"`
}

// Form represents the form in the config
type Form struct {
	Key            string   `yaml:"key"`
	AllowedDomains []string `yaml:"allowed_domains"`
	ToEmail        []string `yaml:"to_email"`
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
