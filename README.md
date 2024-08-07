# kuri

## Requires

* [Go](https://go.dev/)
* [Make](https://www.gnu.org/software/make/)
* [Air](https://github.com/air-verse/air)

## Uses

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
make serve
```

Default URL is http://localhost:3000.

## Build

```
make build
```

## Clean

```
make clean
```
