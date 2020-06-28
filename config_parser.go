package mailbear

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// GetConfigFromFile parses a config file into a Config struct
func GetConfigFromFile(configFile string) (config *Config, err error) {

	config = &Config{}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		err = errors.Wrapf(err, "couldn't open config file %q", configFile)
		return
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		err = errors.Wrap(err, "couldn't parse config file")
		return
	}

	return
}
