package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/Hulhay/jk-pengker/shared"
	"github.com/Hulhay/jk-pengker/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(u usecase.Token) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeaderBearer := c.GetHeader("Authorization")

		bearer := strings.Split(authHeaderBearer, " ")

		if bearer[0] != "Bearer" {
			res := shared.BuildErrorResponse("Failed to process request", "unauthorizes")
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		authHeader := bearer[1]

		if authHeader == "" {
			response := shared.BuildErrorResponse("Failed", "No token found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := u.ValidateToken(authHeader)
		if token.Valid {
			_ = token.Claims.(jwt.MapClaims)
		} else {
			log.Println(err)
			response := shared.BuildErrorResponse("Token is not valid", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
