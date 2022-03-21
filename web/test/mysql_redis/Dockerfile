FROM golang:1.14 as builder
ENV GO111MODULE=on\
    CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64\
    GOPROXY="https://goproxy.cn,direct"
ARG dir=test
WORKDIR /$dir
COPY . .
RUN go build -o test .

FROM scratch as prod
ARG dir=test
WORKDIR /$dir
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/cert
COPY --from=builder /$dir .
CMD ["./test"]