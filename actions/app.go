package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"
	"github.com/gobuffalo/x/sessions"
	"github.com/joepena/monsters/models"
)

// TODO: remove this from source later
var SERVER_SECRET = []byte(envy.Get("SERVER_SECRET","super_secret"))

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		// init DB
		models.GetDBInstance()

		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			SessionName:  "_monsters_session",
		})

		app.Use(func (next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				// change the context to MonsterContext
				return next(MonsterContext{c})
			}
		})

		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Set the request content type to JSON
		app.Use(middleware.SetContentType("application/json"))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// middleware
		app.Use(authenticateRequest)

		authGroup := app.Group("/auth")
		authGroup.Middleware.Skip(authenticateRequest, createUserHandler, loginHandler) // do not verify a token for these registration/login handlers
		authGroup.POST("/user", createUserHandler)
		authGroup.POST("/login", loginHandler)

		userGroup := app.Group("/user")
		userGroup.GET("/{userID}", userDataHandler)
		userGroup.GET("/{userID}/assets", userAssetDataHandler)
		userGroup.PUT("/monster", renameMonsterHandler)
		userGroup.POST("/monster", addMonsterHandler)
		userGroup.POST("/monster/attack", addMonsterAttackHandler)

		dexGroup := app.Group("/dex")
		dexGroup.POST("/monster", createMonsterHandler)
		dexGroup.POST("/attack", createAttackHandler)

		app.GET("/download/{assetID}", downloadHandler)
	}

	return app
}
