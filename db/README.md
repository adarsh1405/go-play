Postgres DB setup for this project (database: employee_info)

This folder contains SQL to create a table matching the `account` struct in `main.go`.

Quick instructions (macOS, psql installed):

1. Ensure Postgres server is running and you have the user `USERNAME` (per your instruction).

2. Create the database and apply the schema:

If your user can create databases:

```sh
# create database
psql -U USERNAME -c "CREATE DATABASE employee_info;"
# apply schema
psql -U USERNAME -d employee_info -f db/schema.sql
```

If you prefer to create DB via createdb:

```sh
createdb -U USERNAME employee_info
psql -U USERNAME -d employee_info -f db/schema.sql
```

3. Verify the table exists:

```sh
psql -U USERNAME -d employee_info -c "\dt"
psql -U USERNAME -d employee_info -c "SELECT * FROM users LIMIT 5;"
```

Notes:
- The `users` table flattens the nested `Company` object into `company_name` and `company_catchphrase` columns.
- If your Postgres user requires a password or uses different host/port, add -h and -p options.
