FROM golang:1.12.8-alpine3.10
WORKDIR /build
RUN apk --no-cache add git
COPY . /build/.
RUN go get -u -v github.com/ahmetb/govvv
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags "$(govvv -flags)" email2matrix-server.go

FROM alpine:3.10
WORKDIR /
RUN apk --no-cache add ca-certificates
COPY --from=0 /build/email2matrix-server .
CMD ["./email2matrix-server"]
