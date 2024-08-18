FROM golang:1.21.7-alpine3.19 AS builder

RUN go install github.com/grpc-ecosystem/grpc-health-probe@v0.4.25

WORKDIR /src/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./cmd/server

FROM alpine:3.19

ARG USER=user

RUN apk add --update sudo curl

RUN adduser -D $USER \
        && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
        && chmod 0440 /etc/sudoers.d/$USER

# install as root
COPY --from=builder /go/bin/grpc-health-probe /usr/bin/
RUN chmod +x /usr/bin/grpc-health-probe

USER $USER

# install as user
COPY --from=builder --chown=$USER:$USER /bin/app /usr/bin/
RUN chmod +x /usr/bin/app

CMD ["/usr/bin/app" , "run"]
