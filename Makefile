doc:
	(cd /tmp && go get github.com/robertkrimen/godocdown/godocdown)
	godocdown client > README.md

.PHONY: docker test check staticcheck golint

docker:
	docker-compose run --service-ports client bash

test:
	gotest -v -cover ./...

golint:
	golint ./...

govet:
	go vet ./...

staticcheck:
	staticcheck ./...

check: test golint govet staticcheck