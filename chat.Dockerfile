FROM golang:1.22.5-alpine AS builder

COPY . /github.com/andredubov/chat
WORKDIR /github.com/andredubov/chat

RUN go mod download && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/chat ./cmd/chat/main.go

FROM alpine:3.20

WORKDIR /root
COPY --from=builder /github.com/andredubov/chat/bin/chat .

CMD ["./chat"]