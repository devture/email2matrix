FROM docker.io/golang:1.17.6-alpine3.15
WORKDIR /build
RUN apk --no-cache add git
RUN go install github.com/ahmetb/govvv@v0.3.0
COPY . /build/.
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags "$(govvv -flags)" email2matrix-server.go

FROM docker.io/alpine:3.15.0
WORKDIR /
RUN apk --no-cache add ca-certificates
COPY --from=0 /build/email2matrix-server .
CMD ["./email2matrix-server"]
