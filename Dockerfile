FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./
COPY gofamlife gofamlife
RUN go mod download && go mod verify

COPY main.go main.go

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]