package middlewares

import (
	"fmt"
	"net/http"
	"post_management/database"
	"post_management/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header from the request
		authHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is present
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is misssing"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Split the Authorization header to get the token
		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || authToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token format"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Extract the token string
		tokenString := authToken[1]

		// Parse the token using the jwt package
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("auth-api-jwt-secret"), nil
		})

		// Check if there was an error parsing the token or if the token is not valid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if the token has expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token Expired"})
			c.AbortWithStatus((http.StatusUnauthorized))
			return
		}

		// If the token is valid, retrieve the user ID from the claims
		var user models.User
		database.DB.Where("ID = ?", claims["id"]).Find(&user)

		// If the user ID is 0, it means the user does not exist
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set the current user in the context for further use in handlers
		c.Set("currentUser", user)
		c.Next()
	}
}
