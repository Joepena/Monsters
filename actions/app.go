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
		userGroup.Middleware.Skip(authenticateRequest, userDataHandler)
		userGroup.GET("/{userID}", userDataHandler)
		userGroup.GET("/{userID}/animations", userAnimationsHandler)
		userGroup.PUT("/monster", renameMonsterHandler)
		userGroup.PUT("/monster/stats", updateMonsterStatsHandler)
		userGroup.POST("/monster", addMonsterHandler)
		userGroup.POST("/monster/attack", addMonsterAttackHandler)
		userGroup.POST("/battle", addBattleResultHandler)

		app.Middleware.Skip(authenticateRequest, leaderboardDataHandler)
		app.GET("/leaderboard", leaderboardDataHandler)

		dexGroup := app.Group("/dex")
		dexGroup.POST("/monster", createMonsterHandler)
		dexGroup.POST("/attack", createAttackHandler)
	}

	return app
}
