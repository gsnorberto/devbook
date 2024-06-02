package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {

	statement, error := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}

	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastInsertedID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedID), nil
}

func (repository Users) GetAllUsers() ([]models.User, error) {
	rows, error := repository.db.Query("SELECT id, name, nick, email, password FROM users")
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		error := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password)
		if error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	if error = rows.Err(); error != nil {
		return nil, error
	}

	return users, nil
}
