package infrastructure

import (
	"encoding/json"
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

func buildAuthParameters(context echo.Context, command interfaces.ICommand) (*value.RequestData, error) {
	model := new(value.RequestData)
	model.Args = command.GetModelValidate().Modal
	model.Authorization = context.Request().Header.Get("Authorization")

	user, err := authService.VerifyAndDecode(model.Authorization)

	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(user)

	json.Unmarshal(jsonBody, &model.UserInfo)

	return model, err
}

type AuthenticationHandler struct{}

func (h *AuthenticationHandler) Handle(context echo.Context, command interfaces.ICommand) error {
	model, err := buildAuthParameters(context, command)

	if err != nil {
		return onUnauthorized(context, err)
	}

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

func NewAuthHandler() IHandler {
	return &AuthenticationHandler{}
}
