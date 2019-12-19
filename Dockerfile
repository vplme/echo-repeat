FROM golang:1.13.5-buster AS builder

WORKDIR /app

ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY . .

RUN go build

FROM alpine:3.10.3

COPY --from=builder /app/echo-repeat /echo-repeat

EXPOSE 8080

ENV PORT=8080

CMD ["/echo-repeat"]
