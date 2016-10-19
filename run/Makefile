root := $(dir $(shell pwd))

executable := example
dependencies := $(shell find $(root)src -name '*.go')

export GOMAXPROCS=1

build: $(executable)

test:
	@cd $(root)src && go test

run: $(executable)
	@./$(executable) -c $(executable).json

clean:
	@rm -f $(executable)

$(executable): $(dependencies)
	@cd $(root)src/$(1) && go build -o $(root)run/$@

.PHONY: build test run clean
