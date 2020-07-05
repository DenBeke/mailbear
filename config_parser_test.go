package mailbear_test

import (
	"testing"

	"github.com/DenBeke/mailbear"
	"github.com/stretchr/testify/require"
)

const (
	configFile      = "config_sample.yml"
	nonExistingFile = "random file that does not exist"
	invalidFile     = "config.go"
)

func TestGetConfigFromFile(t *testing.T) {

	config, err := mailbear.GetConfigFromFile(configFile)

	require.NoError(t, err, "should be able to parse config file")

	require.Len(t, config.Forms, 1, "should have a form in the sample config file.")

}

func TestGetConfigFromFileNotExists(t *testing.T) {

	_, err := mailbear.GetConfigFromFile(nonExistingFile)

	require.Error(t, err, "should have error on no existing file")

}

func TestGetConfigFromFileInvalid(t *testing.T) {

	_, err := mailbear.GetConfigFromFile(invalidFile)

	require.Error(t, err, "should have error on non-json file")

}
