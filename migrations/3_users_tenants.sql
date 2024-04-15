CREATE TABLE IF NOT EXISTS users_tenants (
    user_id VARCHAR(36) PRIMARY KEY,
    tenant_id VARCHAR(36) NOT NULL,
    role VARCHAR(45) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
)