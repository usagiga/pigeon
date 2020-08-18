xxx:
	@echo "Please select optimal option."

build:
	@go build -o pigeon .

clean:
	@rm -f ./pigeon
	@rm -f ./docker/__debug_bin

run:
	@go run .

debug:
	@air

test:
	@go test -v "./..."
