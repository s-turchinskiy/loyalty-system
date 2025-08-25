sudo su - postgres
psql
CREATE DATABASE loyaltysystem;
GRANT CREATE ON SCHEMA public TO my_database_user;
Grant all privileges on database loyaltysystem to my_database_user;
# CREATE SCHEMA IF NOT EXISTS loyaltysystem;

#go generate ./internal/models/models.go
