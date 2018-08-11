package main

import (
	"github.com/kataras/iris"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/iris-contrib/middleware/cors"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

/*
1. GraphQL Rest API

- sign up db record save 200
- login jwt token 200
- passowrd reset/change jwt token 200
- cors support

2. Note

- use sql(add an instructions on how to create DB schema or preferably migrations.)
- DDAY til sun
- QOC not features!!!
- coding convention and no please silly naming and best practices of frameworks


POST /login      log in
GET  /users      get All user info
GET  /user/1     get Single user info
POST /user       sign up
PUT  /user/1     password reset/hange


> Iris
> SQLite
> go-REST
> go-CORS
> jwt-go
*/

func AuthHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.Writef("%s", user.Signature)
}

func main() {

	port := "3000"
	app := makeNew()

}

func makeNew() *iris.Application {
	app := iris.New()

	// make jwtAuth
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	// make cors
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	/*
		POST /login      log in
		GET  /users      get All user info
		GET  /user/1     get Single user info
		POST /user       sign up
		PUT  /user/1     password reset/hange
	*/
	api := app.Party("/api", corsHandler).AllowMethods(iris.MethodOptions)
	{
		api.Post("/login", login)

		v1 := api.Party("/v1")
		{
			//v1.Use(jwtHandler.Serve)
			v1.Get("/users", users)
			//v1.Get("/users/{id}", user)
			//v1.Post("/user", signup)
			//v1.Post("/user/{id}", passwordChange)
		}
	}

	return app

}

func users(c iris.Context) {

}
