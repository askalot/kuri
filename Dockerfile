FROM golang:1.22-alpine3.20

#RUN mkdir /usr/src/app

WORKDIR /usr/src/app

# Cache dependencies, and only redownload if they change in subsequent builds
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/server ./cmd/main.go

EXPOSE 3000

CMD ["server"]
