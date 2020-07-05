package mailbear

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *MailBear) handleForm(c echo.Context) error {
	formID := c.Param("id")

	// check if form exists
	if !m.formExists(formID) {
		return c.JSON(http.StatusNotFound, mailbearRespone("the given form does not exist"))
	}

	// get form data from body
	data := &FormSubmission{FormID: formID}
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, mailbearRespone(err.Error()))
	}

	// validate form data
	if err := data.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, mailbearRespone(err.Error()))
	}

	// check if domain allowed
	origin := c.Request().Header.Get("Origin")
	form := m.getFormByID(data.FormID)
	if !form.OriginDomainAllowed(origin) {
		return c.JSON(http.StatusForbidden, mailbearRespone("you're not allowed to send from this domain"))
	}

	// send the mail
	err := m.SendMail(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, mailbearRespone("couldn't send the mail"))
	}

	return c.JSON(http.StatusOK, mailbearRespone("form was submitted successfully"))

}
