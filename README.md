# Boogeyman

A simple program that help you get search results from multi search engines instead of just from google.

## What can it does

This program searches through multi search engines and return search result under some of strategies.

> Top

Return top result of each search engines. 

> Cross Matching

Return matched result cross through multi search engines. (Appeared in 2 hay more search engines)

> All (with limit 20)

Return all :)

On ver `1.0`, search engines list:

1. Ask
2. Bing
3. Google

This list will be updated by far ~

## The design

The program's architecture implemented under `the clean architecture` design. More info go [here](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html).

![boogeymain design](public/boogeyman_design.jpg)

## Usage

In case exec file you downloaded's name is `boogeyman`.

Sample full params command

```bash
$ ./boogeyman -engine=bing -strategy=top -keyword=some anything
```

Type `-h` to get help. Return value be like

```$xslt
Usage of ./public/boogeyman-darwin-64:
  -engine string
        search engine: google | bing | ask | all (default "all")
  -keyword string
        search (query) string (default "bar")
  -strategy string
        result show strategy: top | cross | all (default "all")
```

## Run on local

Boogeyman development environment requires: 

1. Golang (1.9.2 or higher). Install go [here](https://golang.org/doc/install).
2. dep (Go dependency management tool). Install go [here](https://github.com/golang/dep).

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
