FROM golang:onbuild

ADD . /go/src/github.com/chrismar035/sudoku-api

RUN go install github.com/chrismar035/sudoku-api

ENV REDIS_ADDR "redis:6379"

ENTRYPOINT /go/bin/sudoku-api

EXPOSE 8080
