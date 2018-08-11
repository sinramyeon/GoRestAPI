package main

import (
	"github.com/kataras/iris"

	"github.com/dgrijalva/jwt-go"
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


> Iris
> SQLite
> go-REST
> go-CORS
> jwt-go
*/

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}

func AuthHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.Writef("%s", user.Signature)
}

func main() {

	app := iris.New()

}
