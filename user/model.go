package model

import (
	"GoRestAPI/config"

	"github.com/kataras/iris/context"
)

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	Password  string `json:"password"`

	Users []User
}

func GetUsers() (User, error) {

	var u User
	rows, err := config.SQL.Query(`SELECT * FROM "USER" ORDER BY "ID"`)
	defer rows.Close()

	if err == nil {
		for rows.Next() {
			err = rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Age, &u.Address, &u.Password)
			if err == nil {
				row :=
					User{ID: u.ID,
						Firstname: u.Firstname,
						Lastname:  u.Lastname,
						Email:     u.Email,
						Age:       u.Age,
						Address:   u.Address,
						Password:  u.Password,
					}
				u.Users = append(u.Users, row)
			}
		}

	}

	return u, err

}

func CreateUser(ctx context.Context) (User, error) {

	u := &User{}
	err := ctx.ReadJSON(u)

	firstname := u.Firstname
	lastname := u.Lastname
	email := u.Email
	age := u.Email
	address := u.Address
	password := u.Password

	err = config.SQL.QueryRow(`
		INSERT INTO public."USER"(
			"FIRSTNAME", "LASTNAME", "EMAIL", "AGE", "ADDRESS", "PASSWORD")
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		firstname, lastname, email, age, address, password).Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Age, &u.Address, &u.Password)

	return *u, err

}
