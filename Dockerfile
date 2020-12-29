FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/HaHadaxigua/melancholy
COPY . $GOPATH/src/github.com/HaHadaxigua/melancholy
RUN go build .

EXPOSE 8003
ENTRYPOINT ["./melancholy"]

