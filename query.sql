-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY created_at;

-- name: CreateTodo :one
INSERT INTO todos (
  title, description, done
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
  set title = $2,
  description = $3,
  done = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todos
WHERE id = $1
RETURNING *;