FROM golang:1.21.5-alpine3.19 AS builder

ENV CGO_ENABLED=0 \
    GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /bin/app ./app/

FROM alpine:latest

RUN apk add --no-cache ca-certificates

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /bin/app /app/app

EXPOSE 8080

USER appuser

CMD [ "./app" ]
