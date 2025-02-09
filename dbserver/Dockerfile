FROM golang:1.18-alpine as builder

RUN go version
ENV GOPATH=/

COPY go.mod .
COPY go.sum .

COPY . .

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

EXPOSE 8080

CMD ["./cmd/main"]