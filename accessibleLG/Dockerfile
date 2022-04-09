# syntax=docker/dockerfile:1

FROM golang:1.17.8-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy

COPY . ./

RUN go build -o /accessibleLG

EXPOSE 9003

CMD [ "/accessibleLG" ]
