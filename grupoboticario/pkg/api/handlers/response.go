package handlers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/modules/auth"
	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/modules/share"
)

func Success(ctx *fiber.Ctx, data interface{}) error {
	if data == nil {
		data = make(map[int]interface{}, 0)
	}

	if reflect.TypeOf(data).Kind() == reflect.Slice && reflect.TypeOf(data).Size() == 0 {
		data = make([]interface{}, 0)
	}

	return ctx.Status(fiber.StatusOK).JSON(data)

}

func Error(ctx *fiber.Ctx, err error) error {
	code, res, isLog := GetResponseError(err)
	if isLog {
		log.WithFields(log.Fields{
			"path":          string(ctx.Request().URI().Path()),
			"status_code":   code,
			"client_ip":     ctx.IP(),
			"method":        string(ctx.Context().Method()),
			"user_agent":    string(ctx.Request().Header.UserAgent()),
			"response_body": res,
			"request_body":  string(ctx.Request().Body()),
		}).
			WithError(err).
			Error("Response error")
	}
	return ctx.Status(code).JSON(res)
}

type ResponseError struct {
	Code    string `json:"code" example:"API|CASHBACK|INTERNAL_SERVER_ERROR"`
	Message string `json:"message" example:"An unexpected error has occurred"`
	// @name ErrorResponse
}

func GetResponseError(err error) (int, ResponseError, bool) {
	if err == auth.ErrForbidden {
		return fiber.StatusForbidden, ResponseError{
			Code:    "API|FORBIDDEN",
			Message: "You need to be authenticated",
		}, false
	}

	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		return fiber.StatusUnprocessableEntity, ResponseError{
			Code:    "API|UNPROCESSABLE_ENTITY",
			Message: fmt.Sprintf("The field {%s} must be {%s}", jsonError.Field, jsonError.Type),
		}, false
	}

	if derr, ok := err.(share.DomainError); ok {
		if derr.Err == "" {
			derr.Err = "INTERNAL_SERVER_ERROR"
		}

		if derr.Description == "" {
			derr.Description = "An unexpected error has occurred"
		}

		return fiber.StatusBadRequest, ResponseError{
			Code:    derr.Error(),
			Message: derr.Description,
		}, true
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		code, res := parseValidatorError(validationErrors)
		return code, res, false
	}

	if httpErrs, ok := err.(fiber.MultiError); ok {
		code, res := parseFiberError(httpErrs)
		return code, res, false
	}

	return fiber.StatusInternalServerError, ResponseError{
		Code:    "API|INTERNAL_SERVER_ERROR",
		Message: "An unexpected error has occurred",
	}, true
}

func parseFiberError(httpErrs fiber.MultiError) (int, ResponseError) {
	for _, value := range httpErrs {
		if emField, ok := value.(fiber.EmptyFieldError); ok {
			return fiber.StatusUnprocessableEntity, ResponseError{
				Code:    fmt.Sprintf("API|REQUEST|%s_IS_REQUIRED", strings.ToUpper(emField.Key)),
				Message: fmt.Sprintf("The param {%s} is required", emField.Key),
			}
		}

		if cnField, ok := value.(fiber.ConversionError); ok {
			return fiber.StatusUnprocessableEntity, ResponseError{
				Code:    fmt.Sprintf("API|REQUEST|%s_IS_INVALID", strings.ToUpper(cnField.Key)),
				Message: fmt.Sprintf("The param {%s} must be {%s}", cnField.Key, cnField.Type.String()),
			}
		}
	}

	return fiber.StatusBadRequest, ResponseError{
		Code:    "API|REQUEST|INVALID_REQUEST",
		Message: "An unexpected error has occurred with your request params",
	}

}

func parseValidatorError(validationErrors validator.ValidationErrors) (int, ResponseError) {
	for _, value := range validationErrors {
		return fiber.StatusUnprocessableEntity, ResponseError{
			Code:    fmt.Sprintf("API|REQUEST|%s_MUST_BE_%s", strings.ToUpper(value.Field()), strings.ToUpper(value.Tag())),
			Message: fmt.Sprintf("The param {%s} must be {%s}", strings.ToLower(value.Field()), value.Tag()),
		}
	}

	return fiber.StatusBadGateway, ResponseError{
		Code:    "API|REQUEST|INVALID_REQUEST",
		Message: "An unexpected error has occurred with your request params",
	}

}
