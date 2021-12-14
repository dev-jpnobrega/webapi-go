package infrastructure

import (
	"encoding/json"
	"log"
	"net/http"
	interfaces "webapi/src/domain/contract/interface"
	value "webapi/src/domain/contract/value"
	authService "webapi/src/infrastructure/helper"

	echo "github.com/labstack/echo/v4"
)

func onUnauthorized(context echo.Context, err error) error {
	r := value.ResponseError{}

	r.StatusCode = http.StatusUnauthorized

	return context.JSONPretty(http.StatusUnprocessableEntity, r, " ")
}

func buildAuthParameters(context echo.Context, command interfaces.ICommand) *value.RequestData {
	model := new(value.RequestData)
	model.Args = command.GetModelValidate().Modal
	model.Authorization = context.Request().Header.Get("Authorization")

	return model
}

// AuthenticationHandler - HTTP handler with authentication
type AuthenticationHandler struct{}

// Handle - Execute http request
func (h *AuthenticationHandler) Handle(context echo.Context, command interfaces.ICommand) error {
	model := buildAuthParameters(context, command)

	user, err := authService.VerifyAndDecode(model.Authorization)

	jsonBody, err := json.Marshal(user)

	log.Print(jsonBody, err)

	json.Unmarshal(jsonBody, &model.UserInfo)

	log.Print(model.UserInfo)

	if err != nil {
		return onUnauthorized(context, err)
	}

	// model.UserInfo = user.

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

// NewAuthHandler ...
func NewAuthHandler() IHandler {
	return &AuthenticationHandler{}
}
