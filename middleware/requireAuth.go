package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/RickHPotter/flutter_rest_api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Abort(c *gin.Context, errorMessage string) {
	c.IndentedJSON(http.StatusUnauthorized, gin.H{
		"Error": errorMessage,
	})
	c.AbortWithStatus(http.StatusUnauthorized)
}

func RequireAuth(c *gin.Context) {
	// ! Get JWT Cookie off request
	tokenStr, err := c.Cookie("Authorisation")
	if err != nil {
		Abort(c, "Failed to fetch cookie.")
		return
	}

	// ! Decode JWT
	// Parse() is useless, better use the ParseWithClaims and declare empty MapClaims{},
	// ParseWithClaims() takes the tokenStr and a function for looking up the key,
	// both used to validation and verification of the signature, if sucess key is returned
	token, err := jwt.ParseWithClaims(
		tokenStr,

		jwt.MapClaims{},

		func(jwtoken *jwt.Token) (interface{}, error) {
			// verifying the signature by trying to convert *jwt.Token.Method into a *jwt.SigningMethod
			if _, ok := jwtoken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("_unexpected signing method %v found_", jwtoken.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		},
	)
	if err != nil {
		Abort(c, "Failed to decode cookie. Was it touched?")
		return
	}

	// ! Validate
	// only after ParseWithClaims(), a token could be invalid, aka not populated
	if !token.Valid {
		Abort(c, "Token not valid. Was it touched or corrupt?")
		return
	}

	// converting *jwt.Token.Claims to jwt.MapClaims so as to validate integrity
	// from interface{} to map[string]interface{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check "exp"claim
		// Note that claims["exp"] is an interface{} and needs to be converted
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			Abort(c, "Session Cookie has expired!")
			return
		}

		// "sub" claim is usually user.ID, no change here
		var user models.User
		models.DB.First(&user, claims["sub"])

		// validate if a user was indeed found with such ID
		if user.ID == 0 {
			Abort(c, "Failed to find a user with given JWT.")
			return
		}

		// optional: attach data to request body, in this case, user
		c.Set("user", user)

		// ! Continue
		// this is a handle before another handle, therefore you need to
		// tell the context to proceed to the next handle, as far as I can see
		c.Next()
	} else {
		Abort(c, "Something wrong with JTW Claims. Wonder what.")
		return
	}

}
