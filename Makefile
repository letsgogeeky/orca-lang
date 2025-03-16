test-lexer:
	cd ./orca && go test -v ./lexer

test: test-lexer
	@echo "Done"

run:
	cd ./orca && go run main.go
