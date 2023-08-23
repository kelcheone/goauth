package models

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) CreateUser() error {
	query := `insert into users(name, email, password) values($1, $2, $3)`

	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

// get user

func (user *User) GetUser() error {
	query := `select name, email from users where id=$1`

	return db.QueryRow(query, user.ID).Scan(&user.Name, &user.Email)
}

// check email

func (user *User) CheckEmail() bool {
	// get password from db
	query := `select id, password from users where email=$1`

	row := db.QueryRow(query, user.Email).Scan(&user.ID, &user.Password)

	return row == nil
}
