package middleware

import (
	//controller "campaign/Controller"
	Models "campaign/Model"
	"campaign/auth"

	"campaign/helper"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)


func AuthMiddleware(authService auth.Service, userRepository Models.Repository) gin.HandlerFunc{

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader,"Bearer"){
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized,"error", nil)
			c.Copy().AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	
		tokenString := ""
		arrayToken := strings.Split(authHeader," ")
		if len(arrayToken) == 2{
			tokenString = arrayToken[1]
		}
		token, err:=  authService.ValidateToken(tokenString)
		if err != nil{
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok:= token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid{
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return


		}

		userID := int(claim["user_id"].(float64))
		user, err := userRepository.FindById(userID)
		if err != nil{
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("current user", user)
	}



}