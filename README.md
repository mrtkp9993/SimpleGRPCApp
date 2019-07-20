# Simple GRPC App

[![Go Report Card](https://goreportcard.com/badge/github.com/mrtkp9993/SimpleGRPCApp)](https://goreportcard.com/report/github.com/mrtkp9993/SimpleGRPCApp)

Simple GRPC example app with Golang.

## Usage

1. Run server.
2. Run client.
3. Type ```pi```, ```e```, ... in client app.

## Generating certificates

```
$ openssl genrsa -out cert/server.key 2048
$ openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650
$ openssl req -new -sha256 -key cert/server.key -out cert/server.csr
$ openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650
```

*Source*: [Link](https://www.digitalocean.com/community/tutorials/openssl-essentials-working-with-ssl-certificates-private-keys-and-csrs)

## To-Do

- [ ] Auth
- [x] Secure Channel (SSL/TLS)
- [ ] REST gateway 
