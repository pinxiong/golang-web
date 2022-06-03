FROM public.ecr.aws/docker/library/golang:1.18.3-alpine3.15 AS build

MAINTAINER "Xiong, Pin" "pinxiongcn@foxmail.com"

WORKDIR /go/src/project/
COPY . .
RUN go mod download; \
    go mod verify; \
    go build -o /bin/project/golang-web src/main.go;

FROM public.ecr.aws/docker/library/golang:1.18.3-alpine3.15 AS final
COPY --from=build /bin/project/ /bin/project/
EXPOSE 8080
ENTRYPOINT ["/bin/project/golang-web"]