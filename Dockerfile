FROM golang:1.17

WORKDIR /go/src/github.com/rbozburun/StudentTeachAPI

COPY . .

RUN go get -v

RUN go build -o main .

CMD ["./main"]