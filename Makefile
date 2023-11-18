.PHONY: build-CompressUrlFunction
build-CompressUrlFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(ARTIFACTS_DIR)/bootstrap -tags lambda.norpc .

.PHONY: test
test:
	go test -v ./...
