FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./analyzer/cmd/server/main.go

RUN rm -rf /api

EXPOSE 50051

ENTRYPOINT [ "/app/main" ]