package mailbear

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	logMiddleware "github.com/neko-neko/echo-logrus/v2"
	nekoLog "github.com/neko-neko/echo-logrus/v2/log"
	log "github.com/sirupsen/logrus"
)

type httpError struct {
	Error string
}

func constructHTTPError(err error) httpError {
	return httpError{err.Error()}
}

// Serve craftboard server
func Serve(config *Config) {

	// prep mailbear
	m := &MailBear{config: config}

	// prep echo
	e := echo.New()

	e.HideBanner = true

	nekoLog.Logger().SetOutput(os.Stdout)
	nekoLog.Logger().SetLevel(echoLog.INFO)
	// nekoLog.Logger().SetFormatter(&logrus.JSONFormatter{
	// 	TimestampFormat: time.RFC3339,
	// })
	e.Logger = nekoLog.Logger()
	e.Use(logMiddleware.Logger())

	// Cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	g := e.Group("/api/v1")
	g.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to mail bear! 🐻")
	})

	g.POST("/forms/:id", m.handleForm)

	log.Println("Starting Mail Bear...")

	e.Logger.Fatal(e.Start(config.Global.HTTP.Address))
}