.PHONY: gen run test

gen:
	@go run ./gen -day $(DAY)

run:
	@go run ./days/day$(DAY) -part $(PART)

test:
	@go test -v -count=1 ./days/day$(DAY)
