FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod download && go mod verify
COPY . .
RUN chmod -R 777 soar
CMD ["go","run","main.go"]