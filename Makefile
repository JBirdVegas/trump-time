build: clean web/till_trump.wasm web/wasm_exec.js
deploy: build upload

# configure make
.PHONY:clean build upload deploy
.DEFAULT_GOAL := build
bin_PROGRAMS = till_trump

clean:
	@rm -f ./web/till_trump.wasm
	@rm -f ./web/wasm_exec.js

web/wasm_exec.js:
	@cp "${GOROOT}/misc/wasm/wasm_exec.js" web/.

web/till_trump.wasm:
	@GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o ./web/till_trump.wasm ./cmd/wasm/.

upload:
	@aws s3 cp ./web s3://tilltrump-web/website --recursive
