package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/RickHPotter/flutter_rest_api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

/*
! OPERATIONS
*/

func SignUp(c *gin.Context) {
	// ! get email and password off body req
	var body models.RequestBody
	if c.BindJSON(&body) != nil || body.Email == "" || body.Password == "" {
		BadReq(c, "Failed to fetch request body. Is it empty or incomplete?")
		return
	}

	// checking if email already exists in the DB is another query, avoidable,
	// but then I'd have holes in the debugging side
	// ! check if email is in use
	user, _ := models.GetUser("email = ?", body.Email)
	if user.ID != 0 {
		BadReq(c, "Failed to Sign Up. Email already in use. Login?")
		return
	}

	// !  hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		BadReq(c, "Failed to encode the password. Is it empty?")
		return
	}

	// !  store in the db
	user.Email, user.Password = body.Email, string(hash)
	if err := models.CreateUser(*user); err != nil {
		BadReq(c, "Failed to create a record in the DB. Is it down?")
		return
	}

	// !  notify the user
	c.IndentedJSON(http.StatusAccepted, gin.H{
		"Message": "Hell Yeah, Mate! " + user.Email + " is logged in.",
	})
}

func Login(c *gin.Context) {
	// !  get email and password off body req
	var body models.RequestBody
	if c.BindJSON(&body) != nil || body.Email == "" || body.Password == "" {
		BadReq(c, "Failed to fetch request body. Is it empty?")
		return
	}

	// !  check if email exists
	user, err := models.GetUser("email = ?", body.Email)
	if err != nil || user.ID == 0 {
		BadReq(c, "Failed to find an account with this email. Is it wrong?")
		return
	}

	// !  check if password matches
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		BadReq(c, "Failed to match this password with given email. Is it wrong?")
		return
	}

	// !  generate a JWT
	// Creates a new JWT with given claims (they're convetional, not created at will)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 60).Unix(),
	})

	// (*jwt.Token).SignedString returns a complete Token with the chosen Signature and Secret
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))

	// i have no idea why a jwt wouldn't be created, but the lack of a secret in .env
	if err != nil {
		BadReq(c, "Failed to create a JWToken. What just happened?\n"+err.Error())
		return
	}

	// ! Create a Cookie
	// Cookie in the same site, http.SameSitLaxMode is equal to SameSite = Lax
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorisation", tokenStr, 3600*24*60, "", "", false, true)

	// check if cookie was indeed created
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		BadReq(c, "Cookie not created. Why???")
		return
	}

	// !  notify the user
	c.IndentedJSON(http.StatusAccepted, gin.H{
		"Message":   "You're logged in, my homie!",
		"Token":     tokenStr,
		"expiresAt": claims["exp"].(int64),
		"User":      user,
	})
}

func Logout(c *gin.Context) {
	// ! destroy the JWT Cookie
	c.SetCookie("Authorisation", "", -60, "", "", false, true)

	// ! notify the user
	c.IndentedJSON(http.StatusAccepted, gin.H{
		"Message": "You're logged out, my homie!",
	})
}
