ifndef $(GOROOT)
    GOROOT=$(shell go env GOROOT)
    export GOROOT
endif

compile:
	GOOS=js GOARCH=wasm go build -o assets/json.wasm cmd/wasm/main.go

copy_glue:
	cp ${GOROOT}/misc/wasm/wasm_exec.js assets/

run:
	STATIC_PATH="assets" go run cmd/server/main.go