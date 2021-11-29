# GoDirBrowser

File sharing (upload/download) application over HTTP, written in Go.

## How to use it

### Installing

Make sure that Go (tested on `go version 1.17linux/amd64`) is installed on your computer.
```sh
go mod tidy
go build -ldflags "-w -s" -o GoDirBrowser.out
./GoDirBrowser.out # The server is running on localhost:8080.
```

### Usage

```sh
-directory string
    Serve from another directory
-password string
    Password protect the page
-port int
    Serve from another port than 8080 (default 8080)
-ssl
    Use an SSL connection
```

## TODO

+ Embed HTML files in the final binary
+ TLS - do not depend on static file
+ Github release - generate multiple architectures for targets

## Framework used

* [gin-framework](https://github.com/gin-gonic/gin)
