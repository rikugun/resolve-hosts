FROM golang:alpine as builder

ARG SRC_PATH=/go/src/github.com/rikugun/resolve-hosts
WORKDIR ${SRC_PATH}
ADD . .
RUN go env -w GOPROXY=https://goproxy.baidu.com,https://goproxy.cn,https://goproxy.io,direct
RUN go build -v -a -o resolve-hosts .

FROM alpine:3.12
LABEL maintainer=rikugun

WORKDIR /app
RUN mkdir static
COPY --from=bulder /go/src/github.com/rikugun/resolve-hosts/resolve-hosts .
COPY --from=bulder /go/src/github.com/rikugun/resolve-hosts/servers.txt .

EXPOSE 3000
CMD ["./resolve-hosts"]
