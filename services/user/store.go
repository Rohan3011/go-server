package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rohan3011/go-server/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	query := `
		SELECT * FROM users WHERE email = $1;
	`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	row := s.db.QueryRowContext(ctx, query, email)

	user := new(types.User)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with email: %s", email)
		}
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {

	query := `
	SELECT * FROM users WHERE id = ?;
`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	row := s.db.QueryRowContext(ctx, query, id)

	user := new(types.User)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with id: %d", id)
		}
		return nil, err
	}

	return user, nil
}

func (s *Store) CreateUser(user types.User) error {

	query := `
	INSERT INTO users (firstname, lastname, email, password)
	VALUES ($1, $2, $3, $4)
`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := s.db.ExecContext(ctx, query, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
