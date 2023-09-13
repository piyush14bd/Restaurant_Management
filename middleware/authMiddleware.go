package middleware

import (
	"fmt"
	helper "golang-restaurant-management/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No uthorization header provided")})
			c.Abort()
			return
		}
		claims, err := helper.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return

		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("Last_name", claims.Last_name)
		c.Set("Uid", claims.Uid)
		// c.Set("User_type", claims.User_type)
		c.Next()
	}

}
