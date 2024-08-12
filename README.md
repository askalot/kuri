# kuri

## Requires

* [Make](https://www.gnu.org/software/make/)
* [Docker](https://www.docker.com/)

## Uses

* [Go](https://go.dev/)
* [Air](https://github.com/air-verse/air)
* [GoDotEnv](https://pkg.go.dev/github.com/joho/godotenv)
* [chi](https://go-chi.io/)

## Quickstart

Check build commands with:
```
make usage
```

Run:
```
make
```

## Setup

```
make setup
```

## Run

Update `ADMIN_USERNAME` and `ADMIN_PASSWORD` in `.env`.

Start the server on `PORT` defined in `.env`:

```
make serveContainer
```

Default URL is http://localhost:3000.

## Build

```
make buildContainer
```

## Clean

```
make clean
```
