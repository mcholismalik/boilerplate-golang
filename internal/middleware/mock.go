package middleware

import (
	"strconv"

	"github.com/mcholismalik/boilerplate-golang/pkg/util/response"

	"github.com/labstack/echo/v4"
)

func Mock(reqDto interface{}, resDto interface{}) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			mockS := c.Request().Header.Get("Mock")
			if mockS != "" {
				//checking
				if v, err := strconv.ParseBool(mockS); err == nil && v {
					if err := c.Bind(reqDto); err != nil {
						return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
					}
					// if err := c.Validate(reqDto); err != nil {
					// 	return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
					// }

					return response.SuccessResponse(resDto).Send(c)
				}
			}

			return next(c)
		}
	}
}
