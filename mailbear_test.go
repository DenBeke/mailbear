package mailbear

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var config = &Config{
	Forms: map[string]*Form{
		"some-form": {
			Key: "some-key",
		},
	},
}

func TestMailBearFormExists(t *testing.T) {
	m := MailBear{config: config}
	exists := m.formExists("test")
	require.False(t, exists, "form shouldn't exist")

	exists = m.formExists("some-key")
	require.True(t, exists, "form should exist")
}

func TestMailBearGetFormByID(t *testing.T) {
	m := MailBear{config: config}

	form := m.getFormByID("some-key")
	require.NotNil(t, form, "should have returned form")
	require.Equal(t, "some-key", form.Key)

	form = m.getFormByID("non-existing-key")
	require.Nil(t, form, "should have returned nil")
}
