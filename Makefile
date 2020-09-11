xxx:
	@echo "Please select optimal option."

build:
	@go build -o pigeon .

clean:
	@rm -f ./pigeon

run:
	@go run .

test:
	@go test -v "./..."
