package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Meta  Meta   `json:"meta"`
	Error string `json:"error"`
}

type Error struct {
	Response     errorResponse `json:"response"`
	Code         int           `json:"code"`
	ErrorMessage error
}

const (
	E_DUPLICATE            = "duplicate"
	E_NOT_FOUND            = "not_found"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
	E_DUPLICATE_ENTITY     = "duplicate_entity"
	E_UNAUTHORIZED         = "unauthorized"
	E_BAD_REQUEST          = "bad_request"
	E_VALIDATION           = "validation"
	E_SERVER_ERROR         = "server_error"
)

type errorConstant struct {
	Duplicate                Error
	NotFound                 Error
	RouteNotFound            Error
	UnprocessableEntity      Error
	DuplicateEntity          Error
	Unauthorized             Error
	BadRequest               Error
	Validation               Error
	InternalServerError      Error
	EmailOrPasswordIncorrect Error
	ConvertionNotFound       Error
	NotEnoughStock           Error
	InsufficientStockProduct Error
}

var ErrorConstant errorConstant = errorConstant{
	Duplicate: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Created value already exists",
			},
			Error: E_DUPLICATE,
		},
		Code: http.StatusConflict,
	},
	EmailOrPasswordIncorrect: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Email or password is incorrect",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	NotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Data not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	RouteNotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Route not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	UnprocessableEntity: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Invalid parameters or payload",
			},
			Error: E_UNPROCESSABLE_ENTITY,
		},
		Code: http.StatusUnprocessableEntity,
	},
	DuplicateEntity: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Duplicate data",
			},
			Error: E_DUPLICATE_ENTITY,
		},
		Code: http.StatusUnprocessableEntity,
	},
	Unauthorized: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Unauthorized, please login",
			},
			Error: E_UNAUTHORIZED,
		},
		Code: http.StatusUnauthorized,
	},
	BadRequest: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Bad Request",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	Validation: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Invalid parameters or payload",
			},
			Error: E_VALIDATION,
		},
		Code: http.StatusBadRequest,
	},
	InternalServerError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Something bad happened",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
}

func ErrorBuilder(res *Error, message error, customMessage ...string) *Error {
	res.ErrorMessage = message

	if strings.Contains(strings.Join([]string{E_VALIDATION, E_BAD_REQUEST, E_DUPLICATE}, ","), res.Response.Error) {
		res.Response.Meta.Message = message.Error()
	}
	if len(customMessage) > 0 {
		res.Response.Meta.Message = ""
		for i := range customMessage {
			res.Response.Meta.Message += customMessage[i]
		}
	}
	return res
}

func CustomErrorBuilder(code int, err string, message string) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: message,
			},
			Error: err,
		},
		Code: code,
	}
}

func ErrorResponse(err error) *Error {
	re, ok := err.(*Error)
	if ok {
		return re
	} else {
		return ErrorBuilder(&ErrorConstant.InternalServerError, err)
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d", e.Code)
}

func (e *Error) ParseToError() error {
	return e
}

func (e *Error) Send(c echo.Context) error {
	var errorMessage string
	if e.ErrorMessage != nil {
		errorMessage = fmt.Sprintf("%+v", errors.WithStack(e.ErrorMessage))
	}
	logrus.Error(errorMessage)
	sentry.CaptureException(e)

	return c.JSON(e.Code, e.Response)
}
