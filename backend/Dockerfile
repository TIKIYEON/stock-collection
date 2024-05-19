FROM golang:1.22.3-alpine3.19 
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main .
RUN ls -la /app

CMD ["./main"]
