package mailbear_test

import (
	"testing"

	"github.com/DenBeke/mailbear"
	"github.com/stretchr/testify/require"
)

var config = &mailbear.Config{
	Forms: map[string]*mailbear.Form{
		"test": {
			Key: "test",
			AllowedDomains: []string{
				"allowed.tld",
				"another.allowed.tld",
			},
		},
	},
}

func TestOriginDomainAllowed(t *testing.T) {

	require.Len(t, config.Forms, 1, "should have a form in the sample config.")

	require.True(t, config.Forms["test"].OriginDomainAllowed("http://allowed.tld"), "allowed domain")
	require.True(t, config.Forms["test"].OriginDomainAllowed("https://another.allowed.tld"), "allowed domain")

}

func TestOriginDomainNotAllowed(t *testing.T) {

	require.Len(t, config.Forms, 1, "should have a form in the sample config.")

	require.False(t, config.Forms["test"].OriginDomainAllowed("http://random.domain.tld"), "not allowed domain")

}
