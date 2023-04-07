#!/bin/bash

service mariadb start
mysql -u root < /userhandle/sqlsetup.sql
go run /userhandle
