-- +goose Up

--Companies-----------------------------------------------------------------------

CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    status BOOLEAN NOT NULL,
    image_url TEXT NOT NULL DEFAULT ''

);

--Models--------------------------------------------------------------------------

CREATE TABLE models(
    id UUID PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id),
    name TEXT NOT NULL UNIQUE,
    image_url TEXT NOT NULL DEFAULT ''
);



--Categories -----------------------------------------------------------------------

CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    image TEXT
);

--Brands-------------------------------------------------------------------------

CREATE TABLE brands (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    image TEXT NOT NULL
);

-- Customers ----------------------------------------------------------------------------

CREATE TABLE customers(
    id UUID PRIMARY KEY,
    customer_company_name VARCHAR(255) NOT NULL,
    contact_person VARCHAR(255) NOT NULL ,
    mobile VARCHAR(20) UNIQUE NOT NULL,
    type VARCHAR(255) NOT NULL,
    customer_designation VARCHAR(100),
    address TEXT,
    flat VARCHAR(100),
    street VARCHAR(100),
    city VARCHAR(100),
    state VARCHAR(100),
    pincode VARCHAR(20),
    payment_mode VARCHAR(20),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);


-- +goose Down

DROP TABLE customers;
DROP TABLE categories;
DROP TABLE brands;
DROP TABLE models;
DROP TABLE companies;