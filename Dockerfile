FROM golang:alpine AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
ARG ENV=prod
WORKDIR /apps
COPY . .
RUN go mod download
RUN go build -ldflags="-s -w" -o /apps/chat main.go

FROM alpine
ENV APP_ENV prod
COPY --from=builder /apps/chat /apps/chat
COPY --from=builder /apps/*.yml /apps/etc
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates tzdata tzdata bash bash-doc bash-completion
RUN chmod 755 /apps/chat
ENV TZ Asia/Shanghai
WORKDIR /apps
EXPOSE 7788
ENTRYPOINT /apps/chat /apps/etc/config.${APP_ENV}.yml