# syumia-workers-turso-integration

My attempt at integrating truso with syumai/workers-package. I have 2 branches net/http and echo. 

The net/http branch uses the net/http package to create a server while the echo branch uses the echo framework to create a server.

Both of these have their bugs and issues, so I made this repo to collab to fix them. Thank you for reading


## Pre-Requisites

Cloudflare Workers requires Wrangler to be installed. You can install Wrangler using npm:  
npm install wrangler --save-dev

Make sure to login in to wrangler in the CLI using the following command:  
npx wrangler login

(You need to have a Cloudflare account to login)

Tinygo is also required to build the workers. You can install tinygo using these docs:  
https://tinygo.org/getting-started/install/


## Installation

Clone the repo and run the following command to install the dependencies  
	go run github.com/syumai/workers/cmd/workers-assets-gen@v0.23.1

turn the main.go into wasm using the following command:  
	tinygo build -o ./build/app.wasm -target wasm -no-debug ./...

## Usage

To run the server using wrangler locally, use the following command:  
	make dev

To deploy the server using wrangler, use the following command:  
	make deploy

To build the server using go, use the following command:  
	go run main.go


## Contributing

Currently, both of these branches are experiencing their own issues, but I am closer to success with the net/http branch.  
**I'd massively you to work on the net/http branch, but any help is appreciated.**

## Errors

The net/http branch does not have any error when it is ran through go, but it is having an error through wrangler. The error is:  

Error: The script will never generate a response  
    at async Object.fetch file:///home/<user>/.npm/_npx/32026684e21afda6/node_modules/miniflare/dist/src/workers/core/entry.worker.js:1029:2

You can test this out for yourself by running the following commands:  
	make dev  
	curl http://localhost:8787  

I recommend you to stop wrangler by pressing X and look at the path to the log file that will hopefully help you to see what the error is.