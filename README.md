# echo-logger
A logging middleware for minimal golang Echo framework with logrus

## Usage

```go
import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echologrus "github.com/spirosoik/echo-logrus"
)

func main() {
	e := echo.New()
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

  // Usage of middleware
	mw := echologrus.NewLoggerMiddleware(logger)
	e.Logger = mw
	e.Use(mw.Hook())

	e.GET("/", func(c echo.Context) error {
		c.Logger().Debug("test") // test logging
		c.Logger().Warn("test") // test logging
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}

```

There is an example in [_example](_examples/) folder and you can run it:

```bash
# in one terminal
go run _examples/main.go
```

```bash
# in second terminal
curl localhost:3000
```

# Contributors

[Roel Reijerse](https://github.com/rollulus) Software Engineer