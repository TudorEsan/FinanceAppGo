FROM golang:1.18-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GIN_MODE=release 

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY . .

RUN go build -o app.out

EXPOSE 8080

CMD ["./app.out"]



