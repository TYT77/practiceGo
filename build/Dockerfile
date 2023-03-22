FROM golang:1.19.1-alpine

RUN apk update && apk add git

WORKDIR /go/src/cmd/

CMD ["go", "mod", "init", "goDevEnv"]
# CMD ["go", "run", "hello.go"]


