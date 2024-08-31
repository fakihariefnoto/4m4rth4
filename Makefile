usage: FORCE
	exit 1

FORCE:

start: FORCE
	@echo " >> building..."
	@mkdir -p log
	@go build ./cmd/billing
	@echo " >> run..."
	@./billing