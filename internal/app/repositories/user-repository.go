package repositories

import (
	"fmt"
	"khodza/rest-api/internal/app/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	query := "SELECT * FROM USERS"
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) CreateUser(user models.User) (models.User, error) {
	query := "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id, username, email"
	var createdUser models.User
	err := r.db.Get(&createdUser, query, user.Username, user.Email)
	if err != nil {
		return models.User{}, err
	}
	return createdUser, nil
}

func (r *UserRepository) GetUser(userID int) (models.User, error) {

	var user models.User
	query := "SELECT * FROM users WHERE id = $1"
	err := r.db.Get(&user, query, userID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(userID int, user models.User) (models.User, error) {
	// Start building query
	updateQuery := "UPDATE users SET"
	params := []interface{}{}
	paramCount := 1

	if user.Username != "" {
		updateQuery += fmt.Sprintf(" username = $%d,", paramCount)
		params = append(params, user.Username)
		paramCount++
	}

	if user.Email != "" {
		updateQuery += fmt.Sprintf(" email = $%d,", paramCount)
		params = append(params, user.Email)
		paramCount++
	}

	// Add updated_at column update
	updateQuery += " updated_at = CURRENT_TIMESTAMP,"

	// Remove the trailing comma and space from the update query
	updateQuery = strings.TrimSuffix(updateQuery, ",")

	if len(params) == 0 {
		// Retrieve the updated user from the database
		var updatedUser models.User
		getQuery := "SELECT * FROM users WHERE id = $1"
		err := r.db.Get(&updatedUser, getQuery, userID)
		if err != nil {
			return models.User{}, err
		}

		return updatedUser, nil
	}

	updateQuery += fmt.Sprintf(" WHERE id = $%d", paramCount)
	params = append(params, userID)

	//executing query
	_, err := r.db.Exec(updateQuery, params...)
	if err != nil {
		return models.User{}, err
	}

	// Retrieve the updated user from the database
	var updatedUser models.User
	getQuery := "SELECT * FROM users WHERE id = $1"
	err = r.db.Get(&updatedUser, getQuery, userID)
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func (r *UserRepository) DeleteUser(userID int) error {
	query := "DELETE FROM users WHERE id = $1"

	_, err := r.db.Exec(query, userID)

	if err != nil {
		return err
	}
	// _=result
	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	return err
	// }

	// if rowsAffected == 0 {
	// 	return fmt.Errorf("user with given ID not found")
	// }
	return nil
}
