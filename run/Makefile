root := $(dir $(shell pwd))

executable := example
dependencies := $(shell find $(root)src -name '*.go')

all: $(executable)

$(executable): $(dependencies)
	@cd $(root)src/$(1) && go build -o $(root)run/$@

test:
	@cd $(root)src && go test

clean:
	@rm -f $(executable)

.PHONY: all test clean
