docker run --rm -ti \
	-v  ${PWD}:/work \
	-w /work \
	golang:1.12.7 \
	sh -c 'go get -u -v github.com/ahmetb/govvv && go build -a -ldflags "$(govvv -flags)" email2matrix-server.go'