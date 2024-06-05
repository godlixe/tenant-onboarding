CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
CREATE TABLE IF NOT EXISTS organizations (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    subdomain VARCHAR(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
CREATE TABLE IF NOT EXISTS users_organizations (
    user_id VARCHAR(36) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
    organization_id VARCHAR(36) REFERENCES organizations (id) ON UPDATE CASCADE ON DELETE CASCADE,
    role VARCHAR(45) NOT NULL
);
CREATE TABLE IF NOT EXISTS tenants (
    id VARCHAR(36) PRIMARY KEY,
    product_id VARCHAR(36) NOT NULL,
    organization_id VARCHAR(36) NOT NULL,
    name VARCHAR(45) NOT NULL,
    status VARCHAR(50) NOT NULL,
    resource_information JSON,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
CREATE TABLE IF NOT EXISTS users_tenants (
    user_id VARCHAR(36) PRIMARY KEY,
    tenant_id VARCHAR(36) NOT NULL,
    role VARCHAR(45) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
CREATE TABLE IF NOT EXISTS apps (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
CREATE TABLE IF NOT EXISTS tiers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(36) PRIMARY KEY,
    app_id INT NOT NULL,
    tier_id INT NOT NULL,
    deployment_schema JSON NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
CREATE TABLE IF NOT EXISTS infrastructures (
    id VARCHAR(36) PRIMARY KEY,
    product_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    deployment_model VARCHAR(255) NOT NULL,
    user_count INT NOT NULL,
    user_limit INT NOT NULL,
    metadata JSON NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);
-- users_tenants table
ALTER TABLE users_tenants
ADD CONSTRAINT fk_users_tenants_user_id FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE users_tenants
ADD CONSTRAINT fk_users_tenants_tenant_id FOREIGN KEY (tenant_id) REFERENCES tenants (id);
-- tenants table
ALTER TABLE tenants
ADD CONSTRAINT fk_tenants_product_id FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE tenants
ADD CONSTRAINT fk_tenants_organization_id FOREIGN KEY (organization_id) REFERENCES organizations (id);
-- products table
ALTER TABLE products
ADD CONSTRAINT fk_products_app_id FOREIGN KEY (app_id) REFERENCES apps (id);
ALTER TABLE products
ADD CONSTRAINT fk_products_tier_id FOREIGN KEY (tier_id) REFERENCES tiers (id);
-- infrastructures table
ALTER TABLE infrastructures
ADD CONSTRAINT fk_infrastructures_product_id FOREIGN KEY (product_id) REFERENCES products (id);