FROM golang:1.12 AS build

RUN mkdir -p /go/src/gin
WORKDIR /go/src/gin

ADD ./gin/ .

ENV GOPATH=/go
ENV GOROOT=/usr/local/go

ENV GO111MODULE=on
ENV GOPROXY=http://mirrors.aliyun.com/goproxy/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch AS prod

COPY --from=build /go/src/gin/gin .
CMD ["./gin"]
