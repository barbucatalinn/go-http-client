FROM golang:latest
RUN go get github.com/rakyll/gotest
RUN go get honnef.co/go/tools/cmd/staticcheck
RUN go get golang.org/x/lint/golint