package main

import (
	"net/http"
	"time"

	"github.com/kataras/iris"

	jwt "github.com/dgrijalva/jwt-go"

	. "GoRestAPI/user"
	. "GoRestAPI/config"

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


	// Gets all users
	// Method:   GET
	// Resource: this to get all all users
	app.Handle("GET", "/users", func(ctx context.Context) {
		results, err :=  user.GetUsers()
		if err != nil {
			ctx.JSON(context.Map{"response": err.Error()})
			ctx.StatusCode(iris.StatusNotFound)
			
		} else {
			ctx.StatusCode(iris.StatusOK) 
			ctx.JSON(results)
		}
	})

	// Gets a single user
	// Method:   GET
	// Resource: this to get all all users
	app.Handle("GET", "/users/{msisdn: id}", func(ctx context.Context) {
		msisdn := ctx.Params().Get("msisdn")
		if msisdn == "" {
			c.StatusCode(iris.StatusBadRequest)
		}

		result, err := user.GetUser()
		if err != nil {
			ctx.JSON(context.Map{"response": err.Error()})
			ctx.StatusCode(iris.StatusBadRequest)
		}
		ctx.JSON(result)
	})

	// Method:   POST
	// Resource: This is to sign up a new user
	app.Handle("POST", "/user", func(ctx context.Context) {
		params := &User{}
		err := ctx.ReadJSON(params)
		if err != nil {
			c.StatusCode(iris.StatusBadRequest)
			ctx.JSON(context.Map{"response": err.Error()})
		} else {

			user, err := user.CreateUser()
			if err != nil {
				ctx.JSON(context.Map{"response": err.Error()})
				ctx.StatusCode(iris.StatusBadRequest)
			} else {			
				ctx.StatusCode(iris.StatusOK)
			}
		}

	})

	// Method:   PUT
	// Resource: This is to update a user record
	app.Handle("PUT", "/users/{msisdn: string}", func(ctx context.Context) {
		msisdn := ctx.Params().Get("id")
		params := &User{}
		err := ctx.ReadJSON(params)
		if msisdn == "" {
			c.StatusCode(iris.StatusBadRequest)
		}else {

			user, err := user.UpdateUser()
			if err != nil {
				ctx.JSON(context.Map{"response": err.Error()})
				ctx.StatusCode(iris.StatusBadRequest)
			} else {			
				ctx.StatusCode(iris.StatusOK)
			}
		}	
	})

	
	// Method : GET
	// Resource : This is to login
	app.Handle("POST", "/login", func(ctx context.Context){

		auth := new(model.Auth)
		err := c.ReadJSON(auth)
		if err != nil {
			c.StatusCode(iris.StatusBadRequest)
			c.WriteString(err.Error())
			return
		}
	
		list := user.GetUsers()
	
		for _, id := range list {
			if auth.UserId == id, auth.Password ==  {
				token := jwt.New(jwt.SigningMethodHS256)
	
				claims := token.Claims.(jwt.MapClaims)
				claims["name"] = "Giancarlos"
				claims["admin"] = true
				claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		
				t, err := token.SignedString([]byte("secret"))
				if err != nil {
					c.StatusCode(iris.StatusInternalServerError)
					c.WriteString(err.Error())
				}
		
				c.JSON(map[string]interface{}{
					"token":  t,
					"expire": claims["exp"],
				})
	
				c.StatusCode(iris.StatusOK)
			}
		}
		
		c.StatusCode(iris.StatusUnauthorized)

	})

	return app

}


