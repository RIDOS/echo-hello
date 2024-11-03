# Echo-hello

[![Go](https://github.com/RIDOS/echo-hello/actions/workflows/go.yml/badge.svg)](https://github.com/RIDOS/echo-hello/actions/workflows/go.yml)

In this project we can see service written in framework [Echo](https://echo.labstack.com/) and use [ZeroLog](https://github.com/rs/zerolog).

How it works? Project use port `:8080` and use endpoint:
```bash
curl http://localhost:8080/status
```

This endpoint output count days before start `2025` lear.

Also, we can use endpoint with header like this:
```bash
curl -H "User-Role: Admin" http://localhost:8080/status
```

And then we can see in server console output:
```bash
⇨ http server started on [::]:8080
5:47AM WRN red button user detected
```

## Getting Started

### With Default

Clone this repo:
```bash
git clone https://github.com/RIDOS/echo-hello
```

Install dependencies:
```bash
go mod download
```

Build app:
```bash
go build echo-hello
```

Run:
```bash
./echo-hello
```

### With Docker

Build your image:
```bash
docker build -t echo-hello .
```

Start container:
```bash
docker run -p 8080:8080 echo-hello
```

## Links

Project written with help [spatecon](https://github.com/spatecon) and I really thanks with this experience.
