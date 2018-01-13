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
	UserId string `json:"user_id"`
	// Auth payload in here
	jwt.StandardClaims
}

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

	c.Render(201, render.JSON(map[string]string{"token": u.ID, "email": u.Email}))

	return nil
}

func loginHandler(c buffalo.Context) error {
	email := c.Value("email").(string)
	password := c.Value("password").(string)

	u := models.User{
		Email:    email,
		Password: password,
	}

	validUser := u.Authenticate()
	if !validUser {
		err := errors.New("invalid credentials")
		log.Error(err)
		c.Flash().Add("error", err.Error())
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
