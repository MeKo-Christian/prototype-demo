# Output folder
output := "bin"

# Default target
default:
    just build-web
    just build-sdl

# WebAssembly build
build-web:
    mkdir -p {{output}}
    GOOS=js GOARCH=wasm go build -o {{output}}/demo.wasm

# Native SDL build
build-sdl:
    mkdir -p {{output}}
    go build -o {{output}}/demo

# Clean up binaries
clean:
    rm -rf {{output}}

# Run SDL version (for dev)
run-sdl:
    just build-sdl
    ./{{output}}/demo

# Run WebAssembly version (for dev)
run-web:
    just build-web
    cd {{output}}
    python3 -m http.server 8080