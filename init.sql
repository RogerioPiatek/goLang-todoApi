-- docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:latest

-- Create database
CREATE DATABASE api_todo;

-- Create user
CREATE USER user_todo;

-- Set user password
ALTER USER user_todo WITH ENCRYPTED PASSWORD '1122';

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE api_todo TO user_todo;

-- Connect to the database
\c api_todo;

-- Create table
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR,
    description TEXT,
    done BOOLEAN DEFAULT FALSE
);

-- Grant privileges on tables and sequences
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO user_todo;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO user_todo;
