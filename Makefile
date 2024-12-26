# Project Paths
PB_DIR = pb
SRC_DIR = src
TS_DIR = $(SRC_DIR)/ts
DIST_DIR = dist
HTML_SRC = index.html
HTML_DIST = $(DIST_DIR)/index.html

# Files
PROTO_FILE = game.proto
WASM_FILE = $(SRC_DIR)/main.wasm
WASM_EXEC = $(SRC_DIR)/wasm_exec.js

# Tools
PROTOC = protoc
LUAC = luac
GO = go
NPM = npm

# Default target
all: proto wasm build

# Compile Protocol Buffers for Go and Javascript
proto:
	@echo "Compiling Protocol Buffers..."
	$(PROTOC) -I=$(PB_DIR) \
		--go_out=paths=source_relative:$(PB_DIR) \
		--ts_out=paths=source_relative:$(TS_DIR) \
		$(PROTO_FILE)
	mv $(TS_DIR)/game.ts $(TS_DIR)/game_pb.mts

# Build Go WebAssembly
wasm:
	@echo "Building Go WebAssembly..."
	cp "$(shell $(GO) env GOROOT)/misc/wasm/wasm_exec.js" $(WASM_EXEC)
	GOOS=js GOARCH=wasm $(GO) build -o $(WASM_FILE)

build:
	@echo "Building with vite..."
	$(NPM) run build

# Serve the application
serve:
	@echo "Serving application with Vite..."
	$(NPM) run serve

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(DIST_DIR)
	rm -f $(TS_DIR)/game_pb.mts
	rm -f $(PB_DIR)/*.pb.go
	rm -f $(WASM_FILE)

# Rebuild everything
rebuild: clean all

.PHONY: proto wasm build serve clean
