NAMESPACE = `echo privyTest`

format:
	$(shell go env GOPATH)/bin/gofumpt -w .
	$(shell go env GOPATH)/bin/gci -w .
	$(shell go env GOPATH)/bin/goimports -w .

generate-mock:
	$(shell go env GOPATH)/bin/mockgen -source repository/interface.go -destination mock/repository/interface_mock.go
	$(shell go env GOPATH)/bin/mockgen -source service/interface.go -destination mock/service/interface_mock.go
	$(shell go env GOPATH)/bin/mockgen -source integration/logging/logging.go -destination mock/integration/logging/logging_mock.go

test:
	go clean -testcache ./...
	go test -tags dynamic -v `go list ./...` -coverpkg=./service/... -coverprofile=coverage.out -cover -failfast


swag:
	swag init

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

