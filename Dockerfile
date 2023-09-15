FROM golang:1.20.6

WORKDIR /app

COPY . /app

RUN go build -o forum

EXPOSE 8888

CMD ["./forum"] & echo "Web application is running at http://localhost:8080/"
