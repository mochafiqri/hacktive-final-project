package middleware

import (
	"finalProject/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UserAuth(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, please login",
		})
		return
	}

	bearer := strings.Split(auth, "Bearer ")

	if len(bearer) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, please login",
		})
		return
	}

	tokStr := bearer[1]

	tok, err := helper.ValidateToken(tokStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("%+v\n", tok)
	fmt.Println(tok.UserId)
	c.Set("user_id", tok.UserId)
	c.Set("email", tok.Email)

	fmt.Println(c.Get("user_id"))
	c.Next()
}
