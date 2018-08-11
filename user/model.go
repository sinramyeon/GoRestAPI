package model

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Address   string `json:"address"`

	Users []User
}



func GetUsers(req *http.Request) (User, error) {

	var u User
	rows, err := config.SQL.Query(`SELECT * FROM "USER" ORDER BY "ID"`)
	defer rows.Close()

	if err == nil {
		for rows.Next() {
			err = rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Age, &u.Address)
			if err == nil {
				row :=
					User{ID: r.UniqueID,
						Name:       r.Name,
						PrepTime:   r.PrepTime,
						Difficulty: r.Difficulty,
						Vegetarian: r.Vegetarian}
				r.Recipes = append(r.Recipes, row)
			}
		}

	}

	return r, err

}