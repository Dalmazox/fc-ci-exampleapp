FROM golang:1.22.3-alpine3.20

WORKDIR /app

COPY . .

RUN go build -o math

CMD [ "./math" ]