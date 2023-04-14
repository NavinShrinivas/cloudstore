FROM golang:1.19.2-bullseye
RUN apt update -y
RUN apt upgrade -y
RUN mkdir /orders
COPY . /orders/
WORKDIR /orders
RUN go mod tidy
CMD ["go","run","."]
EXPOSE 5003
