#######################################################
############## formats, lint, and tests ###############
#######################################################

.PHONY: fmt
fmt:
	@echo "----------------------------------------------------------------"
	@echo " ⚙️  Formatting code..."
	@echo "----------------------------------------------------------------"
	gofmt -s -w ./.

.PHONY: lint
lint:
	@echo "----------------------------------------------------------------"
	@echo " ⚙️  Linting code..."
	@echo "----------------------------------------------------------------"
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2 run -E gofmt --config=.golangci.yaml ./...
	@echo "Linting complete!"

.PHONY: test
test:
	@echo "----------------------------------------------------------------"
	@echo " ⚙️  Testing the code..."
	@echo "----------------------------------------------------------------"
	GOPRIVATE=${PRIVATE_REPOS} go test ./... -v
	@echo "Tests complete!"

#######################################################
########################  run #########################
#######################################################

.PHONY: run
run:
	@echo "Starting ${BINARY_NAME}..."
	GOPRIVATE=${PRIVATE_REPOS} go run ./main.go

