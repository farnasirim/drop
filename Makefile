GIT ?= git
COMMIT := $(shell $(GIT) rev-parse HEAD)
VERSION ?= $(shell $(GIT) describe --abbrev=0 --tags 2>/dev/null)
ENVOY_IMAGE_NAME ?= quay.io/farnasirim/drop-envoy
DROP_IMAGE_NAME ?= quay.io/farnasirim/drop-server
IMAGE_VERSION ?= $(VERSION)
WEBPACK_DEV_FLAGS ?= --mode=development --env.NODE_ENV=local
WEBPACK_PROD_FLAGS ?= --mode=production --env.NODE_ENV=production -p

drop: grpc-go *.go */*.go */*/*.go
	go build github.com/farnasirim/drop/cmd/drop

grpc-go:
	protoc -I proto/ --go_out=plugins=grpc:proto/ proto/drop.proto 

grpc-js:
	protoc -I proto drop.proto \
	--js_out=import_style=commonjs:http/frontend \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:http/frontend

frontend-dev: grpc-js
	cd http/frontend && npx webpack $(WEBPACK_DEV_FLAGS) client.js

frontend-prod: grpc-js
	cd http/frontend && npx webpack $(WEBPACK_PROD_FLAGS) client.js

test: proto/drop.pb.go
	go test ./...

envoy-container:
	cd envoy; docker build -t $(ENVOY_IMAGE_NAME):$(VERSION) .

drop-container: drop
	docker build --force-rm -t $(DROP_IMAGE_NAME):$(VERSION) .

clean:
	rm drop

.phony: frontend grpc-go grpc-web envoy-contianer drop-container containers-push
