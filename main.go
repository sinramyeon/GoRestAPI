package main

import (
	"GoRestAPI/config"
	"GoRestAPI/model"
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/iris-contrib/middleware/cors"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
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
	config.Init()
	port := "3000"
	app := makeNew()
	app.Run(iris.Addr(":" + port))
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

	v1 := app.Party("/api/v1", corsHandler).AllowMethods(iris.MethodOptions)
	{

		v1.Use(jwtHandler.Serve)

		// Gets all users
		// Method:   GET
		// Resource: this to get all all users
		v1.Get("/users", func(ctx iris.Context) {

			results, err := model.GetUsers()
			if err != nil {
				fmt.Println(err)
				ctx.StatusCode(iris.StatusNotFound)

			} else {
				ctx.StatusCode(iris.StatusOK)
				ctx.JSON(results)
			}
		})

		// Gets a single user
		// Method:   GET
		// Resource: this to get all all users
		v1.Get("/users/{id: string}", func(ctx iris.Context) {
			msisdn := ctx.Params().Get("id")
			if msisdn == "" {
				ctx.StatusCode(iris.StatusBadRequest)
			}

			result, err := model.GetUser(msisdn)
			if err != nil {
				fmt.Println(err)
				ctx.StatusCode(iris.StatusBadRequest)
			}
			ctx.JSON(result)
		})

		// Method:   PUT
		// Resource: This is to update a user record
		v1.Put("/users/{id: string}", func(ctx iris.Context) {
			msisdn := ctx.Params().Get("id")
			params := &model.User{}
			params.ID, _ = strconv.Atoi(msisdn)
			err := ctx.ReadJSON(params)

			if err != nil {
				fmt.Println(err)
				ctx.StatusCode(iris.StatusBadRequest)
			}

			if msisdn == "" {
				ctx.StatusCode(iris.StatusBadRequest)
			} else {

				err := model.UpdateUser(params)
				if err != nil {
					fmt.Println(err)
					ctx.StatusCode(iris.StatusBadRequest)
				} else {
					ctx.StatusCode(iris.StatusOK)
				}
			}
		})

	}

	api := app.Party("/api", corsHandler).AllowMethods(iris.MethodOptions)
	{

		// Method:   POST
		// Resource: This is to sign up a new user
		api.Post("/user", func(ctx iris.Context) {
			params := &model.User{}
			err := ctx.ReadJSON(params)

			if err != nil {
				fmt.Println(err)
				ctx.StatusCode(iris.StatusBadRequest)
			} else {
				err := model.CreateUser(params)
				if err != nil {
					fmt.Println(err)
					ctx.StatusCode(iris.StatusBadRequest)
				} else {
					ctx.StatusCode(iris.StatusOK)
				}
			}

		})

		// Method : POST
		// Resource : This is to login
		api.Post("/login", func(ctx iris.Context) {

			auth := new(model.Auth)
			err := ctx.ReadJSON(auth)
			if err != nil {
				fmt.Println(err)
				ctx.StatusCode(iris.StatusBadRequest)
				return
			}

			list, err := model.GetUsers()
			if err != nil {
				fmt.Println(err)
				ctx.StatusCode(iris.StatusBadRequest)
				return
			}

			for _, user := range list.Users {

				if auth.UserId == strconv.Itoa(user.ID) && auth.Password == user.Password {
					token := jwt.New(jwt.SigningMethodHS256)

					claims := token.Claims.(jwt.MapClaims)
					claims["name"] = user.Firstname + user.Lastname
					claims["admin"] = true
					claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

					t, err := token.SignedString([]byte("secret"))
					if err != nil {
						ctx.StatusCode(iris.StatusInternalServerError)
					}

					ctx.JSON(map[string]interface{}{
						"token":  t,
						"expire": claims["exp"],
					})

					ctx.StatusCode(iris.StatusOK)
				}
			}

			ctx.StatusCode(iris.StatusUnauthorized)

		})

	}

	return app

}
