## Purpose

This project is just a simple golang web project.

## How to build

You can build a docker image.
```shell
$ cd golang-web
$ docker build -t golang-web .
```

## How to run

```shell
$ docker run -p 8080:8080 -d golang-web:latest
```

## How to visit

You can enter the listed urls above in your browser.
+ [http://localhost:8080/start](http://localhost:8080/start)
+ [http://localhost:8080/stop](http://localhost:8080/stop)
+ [http://localhost:8080/health](http://localhost:8080/health)
