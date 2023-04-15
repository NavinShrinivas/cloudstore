FROM golang:1.19.2-bullseye
RUN apt update -y
RUN apt upgrade -y
RUN mkdir /reviews
COPY . /reviews/
WORKDIR /reviews
RUN go mod tidy
CMD ["go","run","."]
EXPOSE 5004
