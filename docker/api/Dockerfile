FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./api/server/cmd/http/main.go

RUN rm -rf /analyzer

EXPOSE 4513

ENTRYPOINT [ "/app/main" ]