FROM golang
WORKDIR /app/src/go-rest-example
ENV GOPATH=/app
COPY . /app/src/go-rest-example/
RUN go mod tidy
RUN go build -o main .
CMD ["./main"]