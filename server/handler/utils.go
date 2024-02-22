package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hector-leite/meli-notification/domain/constants"
	"github.com/hector-leite/meli-notification/infra/context"
	"github.com/labstack/echo/v4"
)

func getContext(ctx echo.Context) context.Context {
	ct := ctx.Get(constants.ContextKeyDomain)

	context := ct.(context.Context)

	return context
}

func writeHeader(ctx echo.Context, statusCode int) {
	res := ctx.Response()
	if !res.Committed {
		res.WriteHeader(statusCode)
	}
}

func respond(ctx echo.Context, statusCode int, data []byte) error {
	writeHeader(ctx, statusCode)
	_, err := ctx.Response().Write(data)
	if err != nil {
		ctx.Logger().Errorf("Error writing response body: %s", err.Error())
	}

	return nil
}

func respondJSON(ctx echo.Context, statusCode int, data interface{}) error {
	var response []byte
	var err error
	if ctx.Echo().Debug {
		response, err = json.MarshalIndent(data, "", "  ")
	} else {
		response, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return respond(ctx, statusCode, response)
}

func respondNoContent(ctx echo.Context) error {
	writeHeader(ctx, http.StatusNoContent)
	return nil
}
