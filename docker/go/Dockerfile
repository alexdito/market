FROM golang:1.19

WORKDIR /application

COPY ./backend ./

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go mod tidy

RUN go mod download

RUN go build -o bin/main cmd/app/main.go

RUN ["chmod", "+x", "./entrypoint.sh"]

ENTRYPOINT ["sh" ,"-c", "./entrypoint.sh"]

EXPOSE 8080
