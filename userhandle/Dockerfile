FROM golang:1.19.2-bullseye
RUN apt update -y
RUN apt upgrade -y
RUN apt install default-mysql-server -y
RUN ls /etc/init.d/
RUN mkdir /userhandle
COPY . /userhandle/
COPY ./entrypoint.sh /usr/local/bin/
WORKDIR /userhandle
RUN go mod tidy
RUN mkdir -p /run/mysqld && chown -R mysql:mysql /run/mysqld
RUN chmod +x /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
