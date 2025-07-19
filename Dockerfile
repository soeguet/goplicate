FROM golang:alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o /out/goplicate

ENTRYPOINT ["./out/goplicate"]