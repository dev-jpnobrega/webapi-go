package infrastructure

import (
	interfaces "webapi/src/domain/contract/interface"

	echo "github.com/labstack/echo/v4"
)

// IHandler - Representation hanlder interface
type IHandler interface {
	Handle(context echo.Context, command interfaces.ICommand) error
}
