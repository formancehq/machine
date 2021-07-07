generate_parser:
	cd script && rm -r parser && antlr4 -Dlanguage=Go -o parser NumScript.g4

test:
	go test -v -coverprofile=coverage.out -coverpkg=./... ./...

bench:
	go test -test.bench=. ./...