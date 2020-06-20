package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echologrus "github.com/spirosoik/echo-logrus"
)

func main() {
	e := echo.New()
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = &logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	}
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	mw := echologrus.NewLoggerMiddleware(logger)
	e.Logger = mw
	e.Use(mw.Hook())

	e.GET("/", func(c echo.Context) error {
		c.Logger().Debug("test")
		c.Logger().Warn("test")
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
