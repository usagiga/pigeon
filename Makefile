MOCKGEN=$(shell go env GOPATH)/bin/mockgen

xxx:
	@echo "Please select optimal option."

build:
	@go build -o pigeon .

clean:
	@rm -f ./pigeon

run:
	@go run .

gen_mock:
	@$(MOCKGEN) -source=./infra/infra.go -destination=./util/mock/mock_infra/infra.go
	@$(MOCKGEN) -source=./domain/domain.go -destination=./util/mock/mock_domain/domain.go

test:
	@make gen_mock
	@go test -v "./..."
