// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package todos

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int32
	Done        sql.NullBool
	Title       sql.NullString
	Description sql.NullString
	CreatedAt   time.Time
}