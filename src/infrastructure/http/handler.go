package infrastructure

import (
	"net/http"
	interfaces "webapi/src/domain/contract/interface"
	value "webapi/src/domain/contract/value"

	"github.com/go-playground/validator"
	echo "github.com/labstack/echo/v4"
)

func onUnprocessableEntity(context echo.Context, err error) error {
	r := value.ResponseError{}

	r.StatusCode = http.StatusUnprocessableEntity

	for _, e := range err.(validator.ValidationErrors) {
		r.Append(2, e.Field()+".invalid")
	}

	return context.JSONPretty(http.StatusUnprocessableEntity, r, " ")
}

func onSuccess(context echo.Context, r value.ResponseData) error {
	return context.JSONPretty(http.StatusOK, r, "  ")
}

func onFaliure(context echo.Context, r value.ResponseError) error {
	return context.JSONPretty(r.StatusCode, r, "  ")
}

func buildParameters(context echo.Context, command interfaces.ICommand) *value.RequestData {
	model := new(value.RequestData)
	model.Args = command.GetModelValidate().Modal

	return model
}

// Handler - HTTP handler
type Handler struct{}

// Handle - Execute http request
func (h *Handler) Handle(context echo.Context, command interfaces.ICommand) error {
	model := buildParameters(context, command)

	if err := context.Bind(&model.Args); err != nil {
		return onUnprocessableEntity(context, err)
	}

	if err := context.Validate(model.Args); err != nil {
		return onUnprocessableEntity(context, err)
	}

	rs, errC := command.Execute(*model)

	if errC != nil {
		return onFaliure(context, *errC)
	}

	return onSuccess(context, rs)
}

// NewHandler ...
func NewHandler() IHandler {
	return &Handler{}
}
