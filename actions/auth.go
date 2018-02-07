package actions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/joepena/monsters/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/url"
	"regexp"
)

type AuthClaims struct {
	UserId string `json:"user_id"`
	// Auth payload in here
	jwt.StandardClaims
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func getAuthToken(c buffalo.Context) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["user_id"] = c.Request().Form.Get("email")

	tokenString, _ := token.SignedString(SERVER_SECRET)

	return tokenString, nil
}

func createUserHandler(c buffalo.Context) error {
	f := c.Request().Form

	err := validateCreateUserReqForm(f)
	if err != nil {
		return err
	}

	token, err := getAuthToken(c)
	if err != nil {
		log.Error(err)
	}

	u := models.User{
		AuthToken:    token,
		Email:        f.Get("email"),
		Password:     f.Get("password"),
		PasswordHash: "",
	}

	err = u.Create()

	if err != nil {
		log.Error(err)
		return err
	}

	return 	c.Render(201, render.JSON(map[string]string{"token": u.AuthToken, "email": u.Email, "userId": u.ID}))
}

func validateCreateUserReqForm(values url.Values) error {
	email := values.Get("email")
	password := values.Get("password")

	if email == "" {
		return errors.New("empty email was provided")
	}
	if password == "" || len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if !emailRegexp.MatchString(email) {
		return errors.New("provide a valid email address")
	}

	return nil
}

func loginHandler(c buffalo.Context) error {
	f := c.Request().Form

	u := models.User{
		Email:    f.Get("email"),
		Password: f.Get("password"),
	}

	validUser := u.Authenticate()
	if !validUser {
		err := errors.New("invalid credentials")
		log.Error(err)
		return err
	}

	return c.Render(201, render.JSON(map[string]string{"token": u.AuthToken, "userId": u.ID}))

}
