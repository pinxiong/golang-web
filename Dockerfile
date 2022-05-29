FROM golang:1.16.4-alpine3.13 AS build

MAINTAINER "Xiong, Pin" "pinxiongcn@foxmail.com"

WORKDIR /go/src/project/
COPY . .
RUN go mod download; \
    go mod verify; \
    go build -o /bin/project/golang-web src/main.go;

FROM alpine AS final
COPY --from=build /bin/project/ /bin/project/
EXPOSE 8080
ENTRYPOINT ["/bin/project/golang-web"]


#FROM golang:1.18.1
#
#MAINTAINER "Xiong, Pin" "pinxiongcn@foxmail.com"
#
#RUN apt-get update -y; \
#    apt install -y build-essential; \
#    apt-get install -y vim; \
#    apt-get install -y gcc;
#
#WORKDIR /go/src/project/
#COPY . .
#RUN go mod download; \
#    go mod verify; \
#    go build -o /bin/project/golang-web src/main.go;
#
#EXPOSE 8080
#ENTRYPOINT ["/bin/project/golang-web"]