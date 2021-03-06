# Auth Service

This is the Auth service

Generated with

```
micro new github.com/Demoliang/go-micro/auth --namespace=go.micro.lgm --alias=auth --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.lgm.srv.auth
- Type: srv
- Alias: auth

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./auth-srv
```

Build a docker image
```
make docker
```