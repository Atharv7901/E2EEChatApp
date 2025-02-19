package authstore

import (
	"database/sql"
	"errors"

	"github.com/Atharv7901/E2EEChatApp/db"
)

type AuthStore struct {
	db *sql.DB
}

type AuthStoreInterface interface {
	GetUserByEmail(email string) (*db.User, error)
	GetUserByID(userID int) (*db.User, error)
	CreateUser(user *db.User) (int, error)
}

func NewAuthStore(db *sql.DB) *AuthStore {
	return &AuthStore{
		db: db,
	}
}

func (s *AuthStore) GetUserByEmail(email string) (*db.User, error) {
	user := db.User{}
	err := s.db.QueryRow("SELECT id, firstName, lastName, email, createdAt, lastLoggedIn FROM users WHERE email = ?", email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.LastLoggedIn)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (s *AuthStore) GetUserByID(userID int) (*db.User, error) {
	user := db.User{}
	err := s.db.QueryRow("SELECT id, firstName, lastName, email, createdAt, lastLoggedIn FROM users where id = ?", userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.LastLoggedIn)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, err
}

func (s *AuthStore) CreateUser(user *db.User) (int, error) {
	//check existing user
	var existingID int
	err := s.db.QueryRow("SELECT id FROM users WHERE email = ?", user.Email).Scan(&existingID)
	if err != sql.ErrNoRows {
		return 0, errors.New("email already registered")
	}

	//create new user
	res, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?,?,?,?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userID, _ := res.LastInsertId()
	return int(userID), nil
}
