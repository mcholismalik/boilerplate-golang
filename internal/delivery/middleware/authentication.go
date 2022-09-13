package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"

	"github.com/mcholismalik/boilerplate-golang/internal/model/abstraction"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	ctxval "github.com/mcholismalik/boilerplate-golang/pkg/ctxval"
	res "github.com/mcholismalik/boilerplate-golang/pkg/util/response"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		jwtKey = os.Getenv(constant.JWT_KEY)
	)

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil).Send(c)
		}

		splitToken := strings.Split(authToken, "Bearer ")
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
		}

		var id uuid.UUID
		destructID := token.Claims.(jwt.MapClaims)["id"]
		if destructID != nil {
			id, err = uuid.Parse(destructID.(string))
			if err != nil {
				return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
			}
		}

		var name string
		destructName := token.Claims.(jwt.MapClaims)["name"]
		if destructName != nil {
			name = destructName.(string)
		} else {
			name = ""
		}

		var email string
		destructEmail := token.Claims.(jwt.MapClaims)["email"]
		if destructEmail != nil {
			email = destructEmail.(string)
		} else {
			email = ""
		}

		authCtx := &abstraction.AuthContext{
			ID:    id,
			Name:  name,
			Email: email,
		}
		ctx := ctxval.SetAuthValue(c.Request().Context(), authCtx)
		newRequest := c.Request().WithContext(ctx)
		c.SetRequest(newRequest)

		return next(c)
	}
}
