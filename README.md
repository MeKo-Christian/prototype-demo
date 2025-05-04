# prototype-demo

This is a simple demo coded with the [prototype](https://github.com/gonutz/prototype) library. Just an old-school parallax demo with some scrolling text and a catchy tune.

![screenshot](https://raw.github.com/meko-christian/prototype-demo/main/screenshot.jpg)

## Try it out

👉 **[Run the demo in your browser](https://meko-christian.github.io/prototype-demo/)**

> Requires a modern browser with WebAssembly support (most do).

## 🚀 Building

This project uses [just](https://github.com/casey/just) as a task runner. You can install it via:

```bash
cargo install just
```

Then, you can build the project with:

```bash
just            # Builds both web and native versions
just build-sdl  # Builds the SDL version (native)
just build-web  # Builds the WebAssembly version
```

Binaries will be placed in the `bin` directory.

## 🕹️ Running

### Native (SDL)

```bash
just run-sdl
```

### Web (WASM)

```bash
just run-web
```

This will start a local HTTP server on `http://localhost:8080`, serving the demo.wasm file.
