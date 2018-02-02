package actions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/url"
	"regexp"
	"Monsters/models"
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

	// logic to attach token to user obj
	log.Infof("token generated %v", tokenString)

	return tokenString, nil
}

func createUserHandler(c buffalo.Context) error {
	token, err := getAuthToken(c)
	if err != nil {
		log.Error(err)
	}

	f := c.Request().Form

	if !validateCreateUserReqForm(f) {
		c.Render(400, render.JSON(map[string]string{"message": "email and/or password are wrong. password must be at least 8 characters."}))
		return nil
	}

	u := models.User{
		ID:           token,
		Email:        f.Get("email"),
		Password:     f.Get("password"),
		PasswordHash: "",
	}

	err = u.Create(models.GetDBInstance())

	if err != nil {
		log.Error(err)
		return err
	}

	c.Render(201, render.JSON(map[string]string{
		"token": u.ID,
		"email": u.Email,
	}))

	return nil
}

func validateCreateUserReqForm(values url.Values) bool {
	email := values.Get("email")
	password := values.Get("password")

	if email == "" || password == "" {
		return false
	}

	if !emailRegexp.MatchString(email) {
		return false
	}

	if len(password) < 8 {
		return false
	}

	return true
}

func loginHandler(c buffalo.Context) error {
	f := c.Request().Form

	u := models.User{
		Email:    f.Get("email"),
		Password: f.Get("password"),
	}
	log.Info(u)
	validUser := u.Authenticate()
	if !validUser {
		err := errors.New("invalid credentials")
		log.Error(err)
		return err
	} else {
		/**
		// create valid session

		s := c.Session()
		s.Set("email", u.Email)
		err := s.Save()
		if err != nil {
			return err
		}
		**/
	}

	return nil
}
