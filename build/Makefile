root := $(shell dirname "$(shell pwd)")
source := ${root}/src
build := ${root}/build
input := ${build}/assets
output := ${build}/results

executable := example
problem := rebecca
dependencies := $(shell find "${source}" -name '*.go')

export GOMAXPROCS=1

build: ${executable}

test:
	@cd "${source}" && go test ./...

clean: flush
	@rm -f ${executable}

flush:
	@rm -rf "${output}/${problem}"*

${problem}: ${output}/${problem}.h5

${output}/%.h5: ${executable}
	@./$< -c "${input}/$*.json" -o "$@"

${executable}: ${dependencies}
	@cd "${source}" && go build -o "${build}/$@"

.PHONY: build test clean flush ${problem}
