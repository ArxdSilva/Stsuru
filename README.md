[![Build Status](https://travis-ci.org/ArxdSilva/Stsuru.svg?branch=master)](https://travis-ci.org/ArxdSilva/Stsuru)
[![Go Report Card](https://goreportcard.com/badge/github.com/arxdsilva/Stsuru)](https://goreportcard.com/badge/github.com/arxdsilva/Stsuru)

# Stsuru

- **Simple** link 'shortener';
- Written in [Go](http://golang.org);

## Introduction
Implementation of a simple URL shortener in Golang. Intended to `hash` & `display` sortened URL's in a 'pure' Golang's server. It uses Gorilla's mux to handle server requests.

## Instalation

### Stable
If you need the latest stable version I recomend you to get our latest [**Release - 3.0.0.3**](https://github.com/ArxdSilva/Stsuru/releases).

### Go Get
- _If you enjoy, dont forget to **Star** the repository! :)_

The easiest way is to install with go get (**needed Golang 1.7 or later installed**):
```shell
$ go get -u github.com/arxdsilva/Stsuru
```
### Git Clone
As an alternative you can use also `git clone`:
- SSH
```shell
$ git clone git@github.com:ArxdSilva/Stsuru.git
```
- HTTPS
```shell
$ git clone https://github.com/ArxdSilva/Stsuru.git
```

## Usage
```shell
$ cd (PATH)/github.com/arxdsilva/Stsuru
```
### Tests
```shell
$ go test -v ./...
```
### Start server
```shell
$ go run main.go
```
Open your **browser** and type:
[`localhost:8080`](http://localhost:8080/)
### Building packages
```shell
$ go build
```


## LICENSE
Check our [MIT](https://github.com/ArxdSilva/Stsuru/blob/master/LICENSE) license file for more info.

## Credits
Made by **[@arxdsilva](https://twitter.com/arxdsilva)** with great help of [Tsuru team](https://github.com/tsuru/tsuru) & contributors!

## Extra
Want to implement a free PaaS & that is Open source? Check [Tsuru](https://github.com/tsuru/tsuru)!
