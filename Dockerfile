FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod download && go mod verify

COPY . .
RUN cd soar && \
mkdir result && mkdir sql && \
wget https://github.com/XiaoMi/soar/releases/download/0.11.0/soar.linux-amd64 && \
chmod -R 777 ./

CMD ["go","run","main.go"]