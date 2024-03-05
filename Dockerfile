FROM golang:latest

WORKDIR /api

COPY ./go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o api

EXPOSE 8000

CMD ["./api"]