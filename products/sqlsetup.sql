drop database if exists cloudstore;
create database cloudstore;
use cloudstore
drop user if exists 'storeuser'@'localhost';
create user 'storeuser'@'localhost' IDENTIFIED BY 'pass1234';
grant all on *.* to 'storeuser'@'localhost';
