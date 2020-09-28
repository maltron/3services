FROM golang:alpine AS builder
WORKDIR /B
ENV GOPATH=/B
ADD . /B
RUN apk --no-cache add git && \
    cd /B && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /B/serviceB .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /B/serviceB .
CMD ["./serviceB"]
