go mod tidy 

export $(cat .env | xargs)
echo 'drop database if exists '$DATABASE_NAME';' > sqlsetup.sql
echo 'create database '$DATABASE_NAME' CHARACTER SET utf8 COLLATE utf8_general_ci;' >> sqlsetup.sql
echo 'use '$DATABASE_NAME';' >> sqlsetup.sql
echo 'drop user if exists '$DATABASE_USERNAME'@'$DATABASE_HOST';' >> sqlsetup.sql
echo 'create user '$DATABASE_USERNAME'@'$DATABASE_HOST' IDENTIFIED BY '\'$DATABASE_PASSWORD\'';' >> sqlsetup.sql
echo 'grant all on *.* to '$DATABASE_USERNAME'@'$DATABASE_HOST';' >> sqlsetup.sql

echo 'insert into users (name, username, email, phone, password, user_type) values ("Admin", "admin", "admin@localhost", "0000000000", "'$ADMIN_PASSWORD_HASHED'", "admin");' >> sqlsetup.sql

mysql -u root < sqlsetup.sql
