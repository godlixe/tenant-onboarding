CREATE TABLE IF NOT EXISTS tiers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
)