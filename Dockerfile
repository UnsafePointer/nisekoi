FROM golang:1.9.2-alpine3.7 as builder
RUN apk update && apk add --no-cache git && \
    go get github.com/Ruenzuo/nisekoi

FROM alpine:3.7
RUN apk update && apk --no-cache add ca-certificates
COPY --from=builder /go/bin/nisekoi /bin/nisekoi
ENTRYPOINT ["/bin/nisekoi"]
CMD ["--help"]
