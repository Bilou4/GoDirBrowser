# GoDirBrowser


## How to use it

### Installing

Make sure that Go (tested on `go version 1.17linux/amd64`) is installed on your computer.
```sh
go mod tidy
go build -ldflags "-w -s" -o GoDirBrowser.out
./GoDirBrowser.out # The server is running on localhost:8080.
```



## Framework used

* [gin-framework](https://github.com/gin-gonic/gin)
