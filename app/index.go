package app

import (
	"fmt"

	"anla.io/hound/handler"
	"anla.io/hound/middleware"
	"anla.io/hound/models"

	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"

	"github.com/houndgo/houndgo/logfile"
	"github.com/iris-contrib/middleware/cors"
	"github.com/jinzhu/configor"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// InitApp is
func InitApp() {

	configor.Load(&models.Config, "app.yml")
	fmt.Printf("config port : %#v", models.Config)

	logfile.Mkdir(logfile.LogFIlePath)

	f := newLogFile()
	defer f.Close()

	app := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PATCH", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization", "Content-Type", "Accept"},
		AllowCredentials: true, // allows everything, use that to change the hosts.
	})

	app.Use(crs)

	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("MySecret2"), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler:  handler.JWTError,
	})

	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())
	// app.UseGlobal(func(ctx iris.Context) {
	// 	ctx.Header("Access-Control-Allow-Origin", "*")
	// 	ctx.Next()
	// })

	app.Use(middleware.Before)

	// attach the file as logger, remember, iris' app logger is just an io.Writer.
	app.Logger().SetOutput(newLogFile())

	app.Options("/*", handler.OptionsHandler)

	app.Get("/", handler.IndexHand)

	v1 := app.Party("/api/v1")
	{
		v1.Get("/", jwtHandler.Serve, myHandler)
		v1.Post("/", handler.UserHand{}.Create)
	}

	// navigate to defafult config http://localhost:8080
	if err := app.Run(iris.Addr(":"+models.Config.APP.Port), iris.WithoutBanner); err != nil {
		if err != iris.ErrServerClosed {
			app.Logger().Warn("Shutdown with error: " + err.Error())
		}
	}
}

func myHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)

	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")

	ctx.Writef("%s", user.SigningString)
}
