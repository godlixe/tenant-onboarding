CREATE TABLE IF NOT EXISTS tenants (
    id VARCHAR(36) PRIMARY KEY,
    product_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    subdomain VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    resource_information JSON,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
)