package actions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/joepena/monsters/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type AuthClaims struct {
	UserId string `json:"user_email"`
	// Auth payload in here
	jwt.StandardClaims
}

func getAuthToken(u *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["user_email"] = u.Email

	tokenString, _ := token.SignedString(SERVER_SECRET)

	return tokenString, nil
}

func createUserHandler(c buffalo.Context) error {
	u := &models.User{}

	err := c.Bind(u)
	if err != nil {
		return errors.WithStack(err)
	}

	token, err := getAuthToken(u)
	if err != nil {
		log.Error(err)
	}

	u.AuthToken = token

	err = u.Create()

	if err != nil {
		log.Error(err)
		return err
	}

	return 	c.Render(201, render.JSON(map[string]string{"token": u.AuthToken, "email": u.Email, "userId": u.ID}))
}

func loginHandler(c buffalo.Context) error {
	u := &models.User{}

	err := c.Bind(u)
	if err != nil {
		return errors.WithStack(err)
	}

	validUser := u.Authenticate()
	if !validUser {
		err := errors.New("invalid credentials")
		log.Error(err)
		return err
	}

	return c.Render(201, render.JSON(map[string]string{"token": u.AuthToken, "userId": u.ID}))

}
