-- users_tenants table
ALTER TABLE users_tenants
ADD CONSTRAINT fk_users_tenants_user_id FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE users_tenants
ADD CONSTRAINT fk_users_tenants_tenant_id FOREIGN KEY (tenant_id) REFERENCES tenants (id);

-- tenants table
ALTER TABLE tenants
ADD CONSTRAINT fk_tenants_product_id FOREIGN KEY (product_id) REFERENCES products (id);

-- products table
ALTER TABLE products
ADD CONSTRAINT fk_products_app_id FOREIGN KEY (app_id) REFERENCES apps (id);

ALTER TABLE products
ADD CONSTRAINT fk_products_tier_id FOREIGN KEY (tier_id) REFERENCES tiers (id);