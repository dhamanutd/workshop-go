-- Create extension for UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    created_at BIGINT DEFAULT (EXTRACT(EPOCH FROM now()) * 1000),
    updated_at BIGINT DEFAULT (EXTRACT(EPOCH FROM now()) * 1000)
);

-- Create products table
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    price FLOAT8 NOT NULL,
    created_at BIGINT DEFAULT (EXTRACT(EPOCH FROM now()) * 1000),
    updated_at BIGINT DEFAULT (EXTRACT(EPOCH FROM now()) * 1000)
);

-- Create orders table
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    product_id UUID,
    amount FLOAT8 NOT NULL,
    status VARCHAR NOT NULL,
    created_at BIGINT DEFAULT (EXTRACT(EPOCH FROM now()) * 1000),
    updated_at BIGINT DEFAULT (EXTRACT(EPOCH FROM now()) * 1000),
    CONSTRAINT fk_user
        FOREIGN KEY (user_id) 
        REFERENCES users (id) 
        ON DELETE SET NULL,
    CONSTRAINT fk_product
        FOREIGN KEY (product_id) 
        REFERENCES products (id) 
        ON DELETE SET NULL
);
