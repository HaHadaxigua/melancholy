# Github: https://github.com/HaHadaxigua
ARG GOVERSION=latest
# FROM 语句下载指定的映像，并创建基于此映像的新容器。
FROM golang:latest AS build

ENV GOPROXY https://goproxy.cn,direct
ENV GOOS=linux
ENV GO111MODULE=on

# WORKDIR 命令在容器中设置当前工作目录，供以下命令使用
WORKDIR /app
#COPY 命令将文件从主计算机复制到容器。 第一个参数 (myapp_code) 是主计算机上的文件或文件夹。 第二个参数 (.) 指定文件或文件夹的名称来充当容器中的目标位置。 在本例中，目标位置是当前工作目录 (/app)。
COPY . .

RUN rm -f melancholy && go run cmd/melancholy.go -no-upgrade build syncthing

EXPOSE 8003
ENTRYPOINT ["./melancholy"]

