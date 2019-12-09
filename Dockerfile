FROM golang:1.13-alpine AS builder

WORKDIR /go/src/apibin

COPY . .

RUN go build -o /go/bin/server -ldflags="-s -w" cmd/apibin/main.go

FROM alpine:3.10

# create non-root user and manage permissions
RUN addgroup -S app && adduser -S -g app app

COPY --from=builder --chown=app:app /go/bin/server /go/bin/server

USER app

CMD /go/bin/server

EXPOSE 8080