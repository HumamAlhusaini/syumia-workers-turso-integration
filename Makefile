.PHONY: dev
dev:
	npx wrangler dev

.PHONY: build
build:
#This command constructs build/ shim.mjs wasm_exec.js worker.mjs
	go run github.com/syumai/workers/cmd/workers-assets-gen@v0.23.1
#
#run this
	
tinygo build -o ./build/app.wasm -target wasm -no-debug ./...
.PHONY: deploy
deploy:
	npx wrangler deploy
