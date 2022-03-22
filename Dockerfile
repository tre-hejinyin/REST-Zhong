FROM golang:1.17.3-buster AS debug
WORKDIR /app
COPY . .
ENV GOPROXY=https://goproxy.cn,direct
RUN go install github.com/cosmtrek/air@v1.15.1
CMD air

FROM golang:1.17.3-buster AS build
WORKDIR /app
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
COPY go.* ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -mod=readonly -v -o server

# k8s、cloud-run用ビルド --target deploy
FROM alpine:3.12 AS deploy
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN apk add --no-cache ca-certificates
COPY --from=build /app/server /server
CMD ["/server"]
