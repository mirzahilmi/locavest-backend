package cfg

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func NewEcho(log *zerolog.Logger) *echo.Echo {
	r := echo.New()
	r.JSONSerializer = parser{}
	r.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		log.Error().Err(err).Caller().Msg("")
		c.JSON(report.Code, report)
	}
	return r
}

// See https://github.com/labstack/echo/blob/master/json.go
type parser struct{}

func (d parser) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := json.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

func (d parser) Deserialize(c echo.Context, i interface{}) error {
	return json.NewDecoder(c.Request().Body).Decode(i)
}
