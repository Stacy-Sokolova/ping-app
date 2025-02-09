FROM golang:1.23.2

RUN go version
ENV GOPATH=/

COPY go.mod .
COPY go.sum .

COPY . .

CMD ["/cmd/maing"]