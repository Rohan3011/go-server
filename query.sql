-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY name;

-- name: CreateTodo :one
INSERT INTO todos (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
  set name = $2,
  bio = $3
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todos
WHERE id = $1
RETURNING *;