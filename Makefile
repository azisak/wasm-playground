ifndef $(GOROOT)
    GOROOT=$(shell go env GOROOT)
    export GOROOT
endif


compile_go:
	GOOS=js GOARCH=wasm go build -o assets/json.wasm cmd/wasm/main.go

compile_c:
	emcc cmd/wasm/main.c -o assets/wasm_c_exec.js -s "EXTRA_EXPORTED_RUNTIME_METHODS=['cwrap']"

compile: compile_go compile_c

copy_glue:
	cp ${GOROOT}/misc/wasm/wasm_exec.js assets/wasm_go_exec.js

run:
	STATIC_PATH="assets" go run cmd/server/main.go