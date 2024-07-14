package todo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type TodoStore struct {
	db *sql.DB
}

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

// TodoInsert represents the data required to insert a new Todo.
type TodoInsert struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// TodoUpdate represents the data fields that can be updated for Todo.
type TodoUpdate struct {
	Title     *string `json:"title,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
}

func NewTodoStore(db *sql.DB) *TodoStore {
	return &TodoStore{db: db}
}

func (s *TodoStore) Create(item TodoInsert) error {
	query := `INSERT INTO todos (title, completed) VALUES ($1, $2);`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, item.Title, item.Completed)
	if err != nil {
		return fmt.Errorf("could not create item: %v", err)
	}
	return nil
}

func (s *TodoStore) List(limit, offset int, filters *map[string]string) ([]Todo, error) {
	query := "SELECT * FROM todos"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Append filter condition
	if filters != nil {
		for k, v := range *filters {
			query += fmt.Sprintf(" WHERE %s LIKE '%%%s%%'", k, v)
		}
	}

	// Add pagination
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var items []Todo
	for rows.Next() {
		item := Todo{}
		err := rows.Scan(&item.ID, &item.Title, &item.Completed, &item.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return items, nil
}

func (s *TodoStore) Read(id int) (*Todo, error) {
	query := `SELECT * FROM todos WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := s.db.QueryRowContext(ctx, query, id)
	item := &Todo{}
	err := row.Scan(&item.ID, &item.Title, &item.Completed, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not read item: %v", err)
	}
	return item, nil
}

func (s *TodoStore) Update(id int, item TodoUpdate) error {
	query := `UPDATE todos SET  Title = $1, Completed = $2 WHERE id = $3`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := s.db.ExecContext(ctx, query, item.Title, item.Completed, id)
	if err != nil {
		return fmt.Errorf("could not update item: %v", err)
	}
	return nil
}

func (s *TodoStore) Delete(id int) error {
	query := `DELETE FROM todos WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not delete item: %v", err)
	}
	return nil
}
