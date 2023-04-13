FROM golang:1.19.2-bullseye
RUN apt update -y
RUN apt upgrade -y
RUN mkdir /userhandle
COPY . /userhandle/
WORKDIR /userhandle
RUN go mod tidy
# for debugging
# CMD ["tail", "-f", "/dev/null"]
CMD ["go", "run", "."]
EXPOSE 5001
