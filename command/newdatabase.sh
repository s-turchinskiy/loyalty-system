sudo su - postgres
psql
CREATE DATABASE loyaltysystem;
GRANT CREATE ON SCHEMA public TO my_database_user;
Grant all privileges on database loyaltysystem to my_database_user;
# CREATE SCHEMA IF NOT EXISTS loyaltysystem;

#go generate ./internal/models/models.go ./internal/models/accrualservice.go
#cd /home/stanislav/go/loyalty-system/cmd/accrual && ./accrual_linux_amd64 -a ":8081"
