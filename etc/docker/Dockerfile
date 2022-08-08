FROM docker.io/golang:1.19.0-alpine3.16
WORKDIR /build
RUN apk --no-cache add git
RUN go install github.com/ahmetb/govvv@v0.3.0
COPY . /build/.
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags "$(govvv -flags)" email2matrix-server.go

FROM docker.io/alpine:3.16.1
WORKDIR /
RUN apk --no-cache add ca-certificates
COPY --from=0 /build/email2matrix-server .
CMD ["./email2matrix-server"]
