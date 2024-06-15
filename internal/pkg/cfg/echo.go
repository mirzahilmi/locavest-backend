package cfg

import (
	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
)

func NewEcho() echo.Echo {
	r := echo.New()
	r.JSONSerializer = parser{}
	return *echo.New()
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
