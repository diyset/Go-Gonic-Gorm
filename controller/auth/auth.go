package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"my-rest/utility"
	"net/http"
	"time"
)
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func LoginHandler(c *gin.Context) {
	var user Credential
	valid := true
	err := c.Bind(&user)
	if err != nil {
		valid = false
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	if user.Username != "myname" {
		valid = false
		fmt.Println("myname")
		c.JSON(http.StatusUnauthorized, gin.H{
			utility.ResponseStatus():  http.StatusUnauthorized,
			utility.ResponseMessage(): "wrong username or password",
		})
		c.Abort()
	}
	if user.Password != "myname123" {
		valid = false
		fmt.Println("myname123")
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
		c.Abort()
	}
	if valid {
		sign := jwt.New(jwt.GetSigningMethod("HS256"))
		token, err := sign.SignedString([]byte("diyansetiyadi"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("diyansetiyadi"), nil
	})

	// if token.Valid && err == nil {
	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"timestamp": time.Now(),
			"message":   "not authorized",
			"error":     err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}