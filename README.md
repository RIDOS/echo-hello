# Echo-hello

[![Go](https://github.com/RIDOS/echo-hello/actions/workflows/go.yml/badge.svg)](https://github.com/RIDOS/echo-hello/actions/workflows/go.yml)

This project showcases a service built using the [Echo](https://echo.labstack.com/) framework and utilizes [ZeroLog](https://github.com/rs/zerolog) for logging.


## How It Works

Project use port `:8080` and use endpoint:
```bash
curl http://localhost:8080/status
```

This endpoint output count of days until the start of the year `2025`.

Also, we can use endpoint with header like this:
```bash
curl -H "User-Role: Admin" http://localhost:8080/status
```

And then we can see in server console output:
```bash
⇨ http server started on [::]:8080
5:47AM WRN red button user detected
```

> This warning indicates that a user with the "admin" role accessed the endpoint.

## Getting Started

### Running the Application Locally

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

### Running the Application with Docker

Build your image:
```bash
docker build -t echo-hello .
```

Start container:
```bash
docker run -p 8080:8080 echo-hello
```

## Links

I would like [spatecon](https://github.com/spatecon) for their instructional video and guidance, which greatly contributed to the development of this application.
