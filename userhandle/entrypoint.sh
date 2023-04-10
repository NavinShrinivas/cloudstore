#!/bin/bash
service mariadb start
./setup.sh
go run /userhandle
