# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

# directory for subsequent commands
WORKDIR /app

#recreate current working directory and cd to it
COPY . .

RUN go build -o main /app/cmd/main-server/main.go

EXPOSE 4000

CMD [ "/app/main" ]