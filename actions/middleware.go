package actions

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"Monsters/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func authenticateRequest(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
				log.Error(err)
				return nil, err
			}
			return []byte(SERVER_SECRET), nil
		})

		if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
			log.Infof("user: %v, expiration: %v", claims.UserId, claims.StandardClaims.ExpiresAt)
			user, err := models.GetDBInstance().GetUserById(tokenString)
			if err != nil {
				return err
			}
			c.Data()["User"] = user
			log.Infof("User model was atttached: %v", c.Data()["User"].(models.User))
		} else {
			return errors.New("Bad auth token!")
		}

		err = next(c)

		return err
	}
}
