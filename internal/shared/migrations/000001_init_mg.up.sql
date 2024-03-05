CREATE TABLE users (
    id int unsigned auto_increment PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INTEGER,
    created_at datetime NULL,
    updated_at datetime NULL,
    deleted_at datetime NULL
);
