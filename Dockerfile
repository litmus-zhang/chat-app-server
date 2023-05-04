FROM golang:1.19

LABEL authors="Litmus"

WORKDIR /app



COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /godocker

EXPOSE 8080

CMD ["/godocker --server=http"]

