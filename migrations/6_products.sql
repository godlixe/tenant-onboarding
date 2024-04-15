CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(36) PRIMARY KEY,
    app_id INT NOT NULL,
    tier_id INT NOT NULL,
    deployment_schema JSON NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
)