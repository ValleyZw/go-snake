# build stage
FROM golang:1.12.0-alpine3.9 AS builder
WORKDIR /src
RUN set -x \
        && echo "https://mirrors.ustc.edu.cn/alpine/v3.8/main" > /etc/apk/repositories \
        && echo "https://mirrors.ustc.edu.cn/alpine/v3.8/community" >> /etc/apk/repositories \
        && apk add --no-cache git \
        && echo "192.30.253.112 github.com" > /etc/hosts \
        && echo "151.101.185.194 github.global.ssl.fastly.net" >> /etc/hosts \
        && rm -rf /var/cache/apk/*

COPY . .

# removing debug informations and compile only for linux target and disabling cross compilation.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
 go build -a -installsuffix cgo -ldflags="-w -s" -o go-snake .

# final stage
FROM scratch 
WORKDIR /app
COPY --from=builder /src/go-snake .
CMD  ["./go-snake"]