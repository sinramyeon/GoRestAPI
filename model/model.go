package model

import (
	"GoRestAPI/config"
	"strconv"
)

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Age       string `json:"age"`
	Address   string `json:"address"`
	Password  string `json:"password"`

	Users []User
}

func GetUsers() (User, error) {

	var u User
	rows, err := config.SQL.Query(`SELECT * FROM "TESTUSER" ORDER BY "ID"`)
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

func CreateUser(u *User) error {

	firstname := u.Firstname
	lastname := u.Lastname
	email := u.Email
	age := u.Age
	address := u.Address
	password := u.Password

	sqlStatement := `INSERT INTO public."TESTUSER"(
	"FIRSTNAME", "LASTNAME", "EMAIL", "AGE", "ADDRESS", "PASSWORD")
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := config.SQL.Exec(sqlStatement, firstname, lastname, email, age, address, password)

	return err

}

func UpdateUser(u *User) error {

	id := u.ID
	firstname := u.Firstname
	lastname := u.Lastname
	email := u.Email
	age := u.Age
	address := u.Address
	password := u.Password

	sqlStatement := `
	UPDATE public."TESTUSER"
	SET "FIRSTNAME"=$2, "LASTNAME"=$3, "EMAIL"=$4, "AGE"=$5, "ADDRESS"=$6, "PASSWORD"=$7
	WHERE "ID"=$1;`
	_, err := config.SQL.Exec(sqlStatement, id, firstname, lastname, email, age, address, password)

	return err
}

func GetUser(id string) (User, error) {

	u := &User{}
	u.ID, _ = strconv.Atoi(id)
	err := config.SQL.QueryRow(`SELECT * FROM "TESTUSER" WHERE "ID"=$1;`, id).Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Age, &u.Email, &u.Address, &u.Password)

	return *u, err

}
