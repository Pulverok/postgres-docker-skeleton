FROM golang:1.20-alpine

WORKDIR /schema_generator
COPY . .
RUN go build -o /bin/schema_generator

CMD ["/bin/schema_generator"]
