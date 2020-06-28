package mailbear

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
	FormKey  string `yaml:"form_key"`
	FormCors string `yaml:"form_cors"`
	ToEmail  string `yaml:"to_email"`
}
