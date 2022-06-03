FROM 120717539064.dkr.ecr.us-west-2.amazonaws.com/golang:1.16-alpine AS build

MAINTAINER "Xiong, Pin" "pinxiongcn@foxmail.com"

WORKDIR /go/src/project/
COPY . .
RUN go mod download; \
    go mod verify; \
    go build -o /bin/project/golang-web src/main.go;

FROM 120717539064.dkr.ecr.us-west-2.amazonaws.com/golang:1.16-alpine AS final
COPY --from=build /bin/project/ /bin/project/
EXPOSE 8080
ENTRYPOINT ["/bin/project/golang-web"]