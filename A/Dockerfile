FROM golang:alpine AS builder
WORKDIR /A
ENV GOPATH=/A
ADD . /A
RUN apk --no-cache add git && \
    cd /A && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /A/serviceA .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /A/serviceA .
CMD ["./serviceA"]
