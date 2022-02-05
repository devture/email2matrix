FROM docker.io/golang:1.17.6-alpine3.15
WORKDIR /build
RUN apk --no-cache add git
COPY . /build/.
RUN go install github.com/ahmetb/govvv@v0.3.0
RUN GOOS=linux GOARCH=arm64 go build -ldflags "$(govvv -flags)" email2matrix-server.go

FROM docker.io/alpine:3.15.0
WORKDIR /
RUN apk --no-cache add ca-certificates
COPY --from=0 /build/email2matrix-server .
CMD ["./email2matrix-server"]
