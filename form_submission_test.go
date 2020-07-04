package mailbear_test

import (
	"testing"

	"github.com/DenBeke/mailbear"

	"github.com/stretchr/testify/require"
)

func TestFormSubmissionValid(t *testing.T) {

	testInputValid := []mailbear.FormSubmission{
		{
			// all fields filled
			Name:    "Foo Bar",
			Email:   "foo@bar",
			Subject: "Some Subject",
			Content: "Some Content\nwith newline",
			FormID:  "some-random-id",
		},
		{
			// without name
			Email:   "foo@bar",
			Subject: "Some Subject",
			Content: "Some Content\nwith newline",
			FormID:  "some-random-id",
		},
	}

	for _, test := range testInputValid {
		err := test.Validate()
		require.NoError(t, err, "should be valid FormSubmission")
	}

}

func TestFormSubmissionInvalid(t *testing.T) {

	testInputInvalid := []mailbear.FormSubmission{
		{
			// all fields filled except id
			Name:    "Foo Bar",
			Email:   "foo@bar",
			Subject: "Some Subject",
			Content: "Some Content\nwith newline",
		},
		{
			// invalid email
			Name:    "Foo Bar",
			Email:   "foogegez",
			Subject: "Some Subject",
			Content: "Some Content\nwith newline",
			FormID:  "some-random-id",
		},
		{
			// subject not set
			Name:    "Foo Bar",
			Email:   "foo@bar",
			Content: "Some Content\nwith newline",
			FormID:  "some-random-id",
		},
		{
			// content not set
			Name:    "Foo Bar",
			Email:   "foo@bar",
			Subject: "Some Subject",
			FormID:  "some-random-id",
		},
	}

	for _, test := range testInputInvalid {
		err := test.Validate()
		require.Error(t, err, "should be invalid FormSubmission")
	}

}
