package mailbear_test

import (
	"testing"

	"github.com/DenBeke/mailbear"
	"github.com/stretchr/testify/require"
)

func TestConfigValidateValid(t *testing.T) {

	// TODO

}

func TestConfigValidateInvalid(t *testing.T) {

	// TODO

}

func TestFormValidateValid(t *testing.T) {

	testInput := []mailbear.Form{
		{
			Key: "test",
			AllowedDomains: []string{
				"allowed.tld",
				"another.allowed.tld",
			},
			ToEmail: []string{
				"some@domain.tld",
			},
		},
	}

	for _, testForm := range testInput {
		err := testForm.Validate()
		require.NoError(t, err, "should be valid form")
	}

}

func TestFormValidateInvalid(t *testing.T) {

	testInput := []mailbear.Form{
		{
			// no allowed domains
			Key:            "test",
			AllowedDomains: []string{},
			ToEmail: []string{
				"some@domain.tld",
			},
		},
		{
			// no to email
			Key: "test",
			AllowedDomains: []string{
				"allowed.tld",
				"another.allowed.tld",
			},
		},
		{
			// no form key
			AllowedDomains: []string{
				"allowed.tld",
				"another.allowed.tld",
			},
			ToEmail: []string{
				"some@domain.tld",
			},
		},
		{
			// invalid email
			Key: "test",
			AllowedDomains: []string{
				"allowed.tld",
				"another.allowed.tld",
			},
			ToEmail: []string{
				"some-invalid-email",
			},
		},
	}

	for _, testForm := range testInput {
		err := testForm.Validate()
		require.Error(t, err, "should be invalid form")
	}

}

var simpleForm = mailbear.Form{
	AllowedDomains: []string{
		"allowed.tld",
		"another.allowed.tld",
	},
}

func TestOriginDomainAllowed(t *testing.T) {

	require.True(t, simpleForm.OriginDomainAllowed("http://allowed.tld"), "allowed domain")
	require.True(t, simpleForm.OriginDomainAllowed("https://another.allowed.tld"), "allowed domain")

}

func TestOriginDomainNotAllowed(t *testing.T) {

	require.False(t, simpleForm.OriginDomainAllowed("http://random.domain.tld"), "not allowed domain")

}
