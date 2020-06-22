FROM golang:1.14.1

MAINTAINER Chimera Kang(chimerakang@gmail.com)

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

WORKDIR $GOPATH/src/github.com/chimerakang/simple_file_server/downloads

WORKDIR $GOPATH/src/github.com/chimerakang/simple_file_server/public
copy public/* ./
RUN ls -l

WORKDIR $GOPATH/src/github.com/chimerakang/simple_file_server
copy *.go ./
RUN go get ./...
RUN go build

EXPOSE 2077

ENTRYPOINT ["./simple_file_server"]

