# Boogeyman

[![][goreportcard-svg]][goreportcard] 
[![][CodeFactor]](https://www.codefactor.io/repository/github/khanhtc1202/boogeyman)
[![][Build Status]](https://travis-ci.org/khanhtc1202/boogeyman)

[Build Status]: https://travis-ci.org/khanhtc1202/boogeyman.svg?branch=master
[CodeFactor]: https://www.codefactor.io/repository/github/khanhtc1202/boogeyman/badge
[goreportcard]: https://goreportcard.com/report/github.com/khanhtc1202/boogeyman
[goreportcard-svg]: https://goreportcard.com/badge/github.com/khanhtc1202/boogeyman

A simple program that help you get search results from multi search engines instead of just from google.

## What can it does

This program searches through multi search engines and returns search results under some of strategies:

> Top

Return top result of each search engines. 

> Cross Matching

Return matched results cross through multi search engines. (Appeared in 2 or more search engines)

> All (with limit 20)

Return all :)

On ver `1.2`, search engines list:

1. Ask
2. Bing
3. Google
4. Yandex (remove on ver 1.1.1)

This list will be updated by far ~

## The design

The program's architecture implemented under `the clean architecture` design. More info go [here](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html).

![boogeymain design](public/boogeyman_design.jpg)

## Usage

In case exec file you downloaded's name is `boogeyman`.

Sample full params command

```bash
$ ./boogeyman -e=bing -s=top -k="some anything"
```

Type `-h` to get help. Return value be like

```$xslt
Usage of ./boogeyman:
  -e string
        search engine(s): google | bing | ask | yandex | all (default "all")
  -k string
        search (query) string (default "bar")
  -s string
        result show strategy: top | cross | all (default "all")
  -v    show application version
  -version
        show application version
```

Sample result

![query keyword cross search engines](public/sample.png)

## Run on local

Boogeyman development environment requires: 

1. Golang (1.9.2 or higher). Install go [here](https://golang.org/doc/install).
2. dep (Go dependency management tool). Install go [here](https://github.com/golang/dep).
3. go-bindata (Go generate bindata from template file). Install go [here](https://github.com/jteeuwen/go-bindata).

Run by `go`

```bash
$ go run main.go
```

or check [Makefile](https://github.com/khanhtc1202/boogeyman/blob/master/Makefile) for building bin on your local.

## Contribution

All contributions will be welcome in this project.

## Download

For linux x64 : [download](public/boogeyman-linux-64)

For MacOS x64 : [download](public/boogeyman-darwin-64)

## License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.
