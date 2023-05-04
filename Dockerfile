FROM golang:1.19.2-bullseye

LABEL authors="Litmus"

WORKDIR /app

COPY . .

RUN go mod dowload

RUN go build -o /godocker

EXPOSE 8080

CMD ["/godocker --server=http"]

