package database

import (
	"database/sql"
	"log"

	"github.com/google/uuid"

	"kaffino/internal/coffeeshop"
)

func (s *service) GetUser(email string) (coffeeshop.User, error) {
	query := `
		SELECT id, email, username, subscriber
		FROM users
		WHERE email = $1
	`
	var user coffeeshop.User
	err := s.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Username, &user.Subscriber)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return coffeeshop.User{}, nil // Return an empty User struct and a nil error
		}
		log.Println("Error retrieving user:", err)
		return coffeeshop.User{}, err
	}

	return user, nil
}
func (s *service) GetUserID(email string) (string, error) {
	// Check if the user exists
	var userID string
	query := `
		SELECT id
		FROM users
		WHERE email = $1
	`
	err := s.db.QueryRow(query, email).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Not user found... Creating new user.......")
			// User doesn't exist, create a new user
			userID, err = s.createUser(email)
			if err != nil {
				log.Println("Error creating user:", err)
				return "", err
			}
			return userID, nil
		} else {
			log.Println("Error retrieving user:", err)
			return "", err
		}
	}

	// User exists, return the existing user ID
	return userID, nil
}

// createUser creates a new user in the database.
func (s *service) createUser(email string) (string, error) {
	// Generate a new UUID for the user ID
	userID := uuid.New().String()

	// Insert the user into the database
	query := `
		INSERT INTO users (id, email)
		VALUES ($1, $2)
	`
	_, err := s.db.Exec(query, userID, email)
	if err != nil {
		log.Println("Database query error:", err)
		return "", err
	}

	return userID, nil
}
