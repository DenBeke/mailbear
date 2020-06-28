package mailbear

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neko-neko/echo-logrus/log"
)

func (m *MailBear) handleForm(c echo.Context) error {
	formID := c.Param("id")

	// check if form exists
	if !m.formExists(formID) {
		return c.JSON(http.StatusNotFound, "the given form does not exist")
	}

	// get form data from body
	data := &FormSubmission{FormID: formID}
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// validate form data
	if err := data.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// send the mail
	err := m.SendMail(data)
	if err != nil {
		log.Warnf("couldn't send mail: %v", err)
		return c.JSON(http.StatusInternalServerError, "couldn't send the mail")
	}

	return c.JSON(http.StatusOK, data)

}
