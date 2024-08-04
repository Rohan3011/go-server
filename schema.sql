CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    done BOOLEAN DEFAULT false,
    title TEXT,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
)   