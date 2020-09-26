xxx:
	@echo "Please select optimal option."

build:
	@go build -o pigeon .

clean:
	@rm -f ./pigeon

run:
	@go run .

gen_mock:
	@mockgen -source=./infra/infra.go -destination=./util/mock/mock_infra/infra.go
	@mockgen -source=./domain/domain.go -destination=./util/mock/mock_domain/domain.go

test:
	@make gen_mock
	@go test -v "./..."
