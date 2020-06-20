package echologrus

import (
	"io"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// Logrus : implement Logger
type Logrus struct {
	logrus.FieldLogger
}

// NewLoggerMiddleware factory method to create middleware
func NewLoggerMiddleware(logger logrus.FieldLogger) Logrus {
	return Logrus{logger}
}

// Level returns logger level. Ignore it
// logrus is responsible for the Level of logging
func (l Logrus) Level() log.Lvl {
	return log.OFF
}

// SetHeader is empty to satisfy interface. can
// be changed by Logrus
func (l Logrus) SetHeader(_ string) {}

// SetPrefix It's controlled by Logger
func (l Logrus) SetPrefix(s string) {}

// Prefix It's controlled by Logger
func (l Logrus) Prefix() string {
	return ""
}

// SetLevel is empty to satisfy the interface, the level
// is controlled via the logrus you pass in factory method
func (l Logrus) SetLevel(_ log.Lvl) {
}

// Output is empty to satisfy the interface. the output is
// controlled via the logrus you pass in factory method
func (l Logrus) Output() io.Writer {
	return os.Stdout
}

// SetOutput change output, default os.Stdout
func (l Logrus) SetOutput(w io.Writer) {
	l.SetOutput(w)
}

// Printj print json log
func (l Logrus) Printj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Print()
}

// Debugj debug json log
func (l Logrus) Debugj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Debug()
}

// Infoj info json log
func (l Logrus) Infoj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Info()
}

// Warnj warning json log
func (l Logrus) Warnj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Warn()
}

// Errorj error json log
func (l Logrus) Errorj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Error()
}

// Fatalj fatal json log
func (l Logrus) Fatalj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Fatal()
}

// Panicj panic json log
func (l Logrus) Panicj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Panic()
}

// Hook method to attach the middleware in echo
func (l Logrus) Hook() echo.MiddlewareFunc {
	return l.logger
}

func (l Logrus) logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()
	l.WithFields(map[string]interface{}{
		"method":  req.Method,
		"uri":     req.RequestURI,
		"status":  res.Status,
		"agent":   req.UserAgent(),
		"latency": stop.Sub(start).String(),
	}).Info()

	return nil
}

func (l Logrus) logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return l.logrusMiddlewareHandler(c, next)
	}
}
