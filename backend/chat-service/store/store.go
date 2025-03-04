package chatstore

import (
	"database/sql"
	"errors"

	"github.com/Atharv7901/E2EEChatApp/chat-service/db"
)

type ChatStoreInterface interface {
	GetAllUsers() ([]db.User, error)
	GetUserByID(id int) (*db.User, error)
}

type ChatStore struct {
	db *sql.DB
}

func NewChatStore(db *sql.DB) *ChatStore {
	return &ChatStore{
		db: db,
	}
}

func (s *ChatStore) GetAllUsers() ([]db.User, error) {
	rows, err := s.db.Query("SELECT id, firstName, lastName, email, createdAt, lastLoggedIn FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []db.User
	for rows.Next() {
		var user db.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.LastLoggedIn); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}

func (s *ChatStore) GetUserByID(id int) (*db.User, error) {
	var user db.User

	query := "SELECT id, firstName, lastName, email, createdAt, lastLoggedIn FROM users WHERE id = ?"
	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.LastLoggedIn)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, err
}
