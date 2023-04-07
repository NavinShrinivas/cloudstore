FROM golang:1.19.2-bullseye
RUN apt update -y
RUN apt upgrade -y
RUN apt install default-mysql-server -y
RUN ls /etc/init.d/
RUN mkdir /products
COPY . /products/
WORKDIR /products
RUN go mod tidy
RUN mkdir -p /run/mysqld && chown -R mysql:mysql /run/mysqld
CMD ["mysqld"]
CMD ["mysql","-u","root","<","./sqlsetup.sql"]
CMD ["go","run","."]
