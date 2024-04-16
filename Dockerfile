FROM golang:1.22.1-alpine

RUN apk --no-cache add zip

WORKDIR /usr/src/app

# RUN go install git hub.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go install github.com/cosmtrek/air@latest

# WORKDIR /usr/src/app/cmd/server

ENTRYPOINT ["air", "-c", ".air.toml"]

# RUN go mod download

# CMD ["air", "-c", ".air.toml"]
