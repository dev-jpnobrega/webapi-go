package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type CustomBinder struct{}

func parseGet(i interface{}, r *http.Request) (err error) {
	dbJson := new(echo.BindUnmarshaler)

	r.ParseForm()

	m := map[string]string{}
	for k, v := range r.Form {
		m[k] = v[0]
	}

	data, _ := json.Marshal(m)

	log.Info(data, dbJson)

	err = json.Unmarshal(data, i)

	return
}

// Bind - To echo server
func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	db := new(echo.DefaultBinder)

	log.Info(err)

	r := c.Request()

	if r.Method == "GET" {
		parseGet(i, r)
		return
	}

	if err = db.Bind(i, c); err != nil {
		return err
	}

	return
}

// CustomValidator - Validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate - Validate interface
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
