FROM golang

WORKDIR /go/src/github.com/grzegorz-bielski/ws-mockend
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["ws-mockend"]

EXPOSE 3000