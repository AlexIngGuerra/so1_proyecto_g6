FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE 5010

CMD ["go","run","main.go"]