-- Schema for employee_info matching the `account` struct in main.go
-- Run this with: psql -h localhost -p 5432 -U adarshpadhi -d employee_info -f db/schema.sql

-- Table 'users' mapping the account struct. Company fields are flattened.
CREATE TABLE IF NOT EXISTS employee (
  id                  INTEGER PRIMARY KEY,
  name                TEXT NOT NULL,
  username            TEXT NOT NULL UNIQUE,
  email               TEXT NOT NULL UNIQUE,
  company_name        TEXT,
  company_catchphrase TEXT,
  created_at          TIMESTAMPTZ DEFAULT now()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);


INSERT INTO users (id, name, username, email, company_name, company_catchphrase)
VALUES (1, 'Leanne Graham', 'Bret', 'Sincere@april.biz', 'Romaguera-Crona', 'Multi-layered client-server neural-net');