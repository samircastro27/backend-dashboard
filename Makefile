run-server:
	@echo "Running server..."
	source ./.env && dapr run -f .

build-server:
	@echo "Building server deletion worker..."
	go build -o bin/deletionworker cmd/deletionworker/main.go

.PHONY: test
test:
	go test -v -test.short -p=1 -run "^*__Unit" ./test/... -coverprofile=coverage.out

.PHONY: coverage
coverage: test
	go tool cover -func=coverage.out

.PHONY: coverage-html
coverage-html: test
	go tool cover -html=coverage.out


brew install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
