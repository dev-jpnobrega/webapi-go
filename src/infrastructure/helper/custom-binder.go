package infrastructure

import (
	"github.com/go-playground/validator"
	echo "github.com/labstack/echo/v4"
)

// CustomBinder - CustomBinder validate struct
type CustomBinder struct{}

// Bind - To echo server
func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	db := new(echo.DefaultBinder)
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
