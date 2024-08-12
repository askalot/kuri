FROM golang:1.22

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

# Cache dependencies, and only redownload if they change in subsequent builds
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
