package controllers

import(
	"teak/config"
	"teak/models"

	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"github.com/spf13/viper"
)

func LoginUser(c echo.Context) error {
	type LoginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var user models.User
	db := config.DB
	login := new(LoginData)
	c.Bind(login)

	if db.Preload("Role").Where("username = ?", login.Username).First(&user).RecordNotFound() {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Invalid username or password"})
	} else if !user.IsActive {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "User disabled"})
	} else if CheckPasswordHash(login.Password, user.Password) {

		// Set custom claims
		claims := &config.JwtCustomClaims{
			user.ID,
			user.Username,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(viper.GetString("jwtSign")))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	} else {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Invalid username or password"})
	}

	return echo.ErrUnauthorized
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}