package mailbear

import (
	"net/http"
	"os"
	"time"

	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	logMiddleware "github.com/neko-neko/echo-logrus/v2"
	nekoLog "github.com/neko-neko/echo-logrus/v2/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	// Validate config
	err := config.Validate()
	if err != nil {
		log.Fatalf("Config file is not valid: %v", err)
	}

	// prep mailbear
	m := &MailBear{config: config}

	// prep echo
	e := echo.New()

	e.HideBanner = true

	// request logging
	nekoLog.Logger().SetOutput(os.Stdout)
	nekoLog.Logger().SetLevel(echoLog.INFO)
	// nekoLog.Logger().SetFormatter(&logrus.JSONFormatter{
	// 	TimestampFormat: time.RFC3339,
	// })
	e.Logger = nekoLog.Logger()
	e.Use(logMiddleware.Logger())

	// Prometheus metrics
	configMetrics := echoPrometheus.NewConfig()
	configMetrics.Namespace = "mailbear"
	e.Use(echoPrometheus.MetricsMiddlewareWithConfig(configMetrics))

	go func() {
		m := echo.New()
		m.HideBanner = true
		m.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		m.Logger.Fatal(m.Start(config.Global.Metrics.Address))
	}()

	// Rate limitting
	e.Use(RateLimitMiddleware(1*time.Minute, 5))

	// Cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	g := e.Group("/api/v1")
	g.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, mailbearRespone("Welcome to mail bear! 🐻"))
	})

	g.POST("/form/:id", m.handleForm)

	log.Println("Starting Mail Bear...")

	e.Logger.Fatal(e.Start(config.Global.HTTP.Address))
}
