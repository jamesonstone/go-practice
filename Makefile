.PHONY: day_one
day_one: TASK=day_one
day_one: run

.PHONY: run
run:
	@echo "Running $(TASK)..."
	@go run cmd/$(TASK)/main.go
