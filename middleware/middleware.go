package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var LocalToken []string

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenCookie, err := c.Cookie("Authorization")
		auth := false

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unautorized User/Token Access",
			})
			c.Abort()
			return
		}

		for _, token := range LocalToken {
			if tokenCookie == token {
				auth = true
				break
			}
		}

		if auth == true {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized User/Token Access",
			})
			c.Abort()
			return
		}
	}
}

// func AuthenticateMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var dataHeader header
// 		if err := c.BindHeader(&dataHeader); err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"message": "Failed To Autorized User",
// 			})
// 			c.Abort()
// 		}

// 		// fmt.Println(dataHeader)

// 		if dataHeader.Auth == "token_guru" {
// 			c.Next()
// 		} else {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"message": "Unautorized User/Token Access",
// 			})
// 			c.Abort()
// 		}
// 	}
// }
