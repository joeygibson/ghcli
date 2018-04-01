FROM library/golang:1.10-stretch

RUN go get -u github.com/golang/dep/...

WORKDIR /go/src/github.com/joeygibson/ghcli

ENTRYPOINT /go/src/github.com/joeygibson/ghcli/scripts/docker-builder

