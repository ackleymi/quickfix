all: lint test

clean:
	rm -rf gen

generate-dev: clean
	mkdir -p gen; cd gen; go run ../cmd/generate-fix/generate-fix.go -prefix=github.com/quickfixgo/quickfix/gen ../spec/*.xml

fmt:
	go fmt `go list ./... | grep -v quickfix/gen`
	gofmt -s -w *.go

lint:
	golangci-lint run -c .golangci.yml

test:
	go test -v -cover . ./datadictionary ./internal

ci-test:
	go test -v -failfast -race -covermode=atomic -coverprofile=coverage.txt . ./datadictionary ./internal

build-all: 
	go build -v `go list ./...`

build-acceptance-server: 
	cd _test; go build -o echo-server

build: build-all build-acceptance-server

fix40:
	cd _test; ./runat.sh $@.cfg 5001 "definitions/server/$@/*.def"
fix41:
	cd _test; ./runat.sh $@.cfg 5002 "definitions/server/$@/*.def"
fix42:
	cd _test; ./runat.sh $@.cfg 5003 "definitions/server/$@/*.def"
fix43:
	cd _test; ./runat.sh $@.cfg 5004 "definitions/server/$@/*.def"
fix44:
	cd _test; ./runat.sh $@.cfg 5005 "definitions/server/$@/*.def"
fix50:
	cd _test; ./runat.sh $@.cfg 5006 "definitions/server/$@/*.def"
fix50sp1:
	cd _test; ./runat.sh $@.cfg 5007 "definitions/server/$@/*.def"
fix50sp2:
	cd _test; ./runat.sh $@.cfg 5008 "definitions/server/$@/*.def"

ACCEPT_SUITE=fix40 fix41 fix42 fix43 fix44 fix50 fix50sp1 fix50sp2 
accept: $(ACCEPT_SUITE)

.PHONY: clean all lint fmt generate-dev test ci-test accept build build-acceptance-server build-all $(ACCEPT_SUITE)
