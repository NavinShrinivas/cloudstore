FROM golang:1.19.2-bullseye
RUN apt update -y
RUN apt upgrade -y
RUN apt install default-mysql-server -y
RUN mkdir /userhandle
COPY . /userhandle/
WORKDIR /userhandle
RUN go mod tidy
RUN mkdir -p /run/mysqld && chown -R mysql:mysql /run/mysqld
RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
EXPOSE 5001
