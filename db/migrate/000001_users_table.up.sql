CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);

