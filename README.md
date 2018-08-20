# Echo server

## Description

Echo server is very simple HTTP server that will listen for incoming HTTP
requests, print them to standard output and reply with a `204` response.

## Install

```bash
go install github.com/gjacquet/echo-server
```

## Run

Assuming your `$GOPATH/bin` is in your path, simply run:

```bash
echo-server
```
