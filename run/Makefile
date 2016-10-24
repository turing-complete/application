root := $(dir $(shell pwd))
input := assets
output := results

executable := example
problem := rebecca
dependencies := $(shell find $(root)src -name '*.go')

export GOMAXPROCS=1

build: $(executable)

test:
	@cd $(root)src && go test

run: $(executable)
	@./$(executable) -c $(input)/$(problem).json -o $(output)/$(problem).h5

clean:
	@rm -f $(executable) $(output)/*.h5

$(executable): $(dependencies)
	@cd $(root)src/$(1) && go build -o $(root)run/$@

.PHONY: build test run clean
