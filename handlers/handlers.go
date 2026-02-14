package handlers

import (
	"inapp_payment_service/graph"
	"inapp_payment_service/model"
	
	"context"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
)

func Handler(ctx echo.Context) error {
	if ctx.Request().Method != http.MethodPost {
		return echo.ErrMethodNotAllowed
	}

	ctx.Request().Body = http.MaxBytesReader(ctx.Response(), ctx.Request().Body, 1<<20) // 1MB limit

	var operation map[string]interface{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&operation)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	query, ok := operation["query"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing query field")
	}

	var variables map[string]interface{}
	if opVars, ok := operation["variables"]; ok && opVars != nil {
		variables = opVars.(map[string]interface{})
	} else {
		variables = make(map[string]interface{})
	}
	reqCtx := context.WithValue(ctx.Request().Context(), model.RequestKey, ctx.Request())

	result := executeQuery(reqCtx, query, variables)

	if result.HasErrors() {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Errors)
	}

	return ctx.JSON(http.StatusOK, result)
}

func executeQuery(ctx context.Context, query string, variables map[string]interface{}) *graphql.Result {
	inatePaymentSchema, err := schema.GetSchema()
	if err != nil {
		return nil
	}
	res := graphql.Do(graphql.Params{
		Schema:         *inatePaymentSchema,
		RequestString:  query,
		VariableValues: variables,
		Context:        ctx,
	})
	return res
}
