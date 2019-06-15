## How to compile wasm

### Requirements

`rustup update`
`rustup target add wasm32-unknown-unknown`

### Benches

Clone https://github.com/perlin-network/life

```
cd bench/cases/fib_recursive
cargo build --release
ls -l target/wasm32-unknown-unknown/release/fib_recursive.wasm
```

### Run Benches

Go to top level 

```
export GO111MODULE=on
go mod vendor
go build
./life bench/cases/fib_recursive/target/wasm32-unknown-unknown/release/fib_recursive.wasm
```

### Cosmos-wasm interface

`init` or `receive`
```json
{
    "signer": "hex addr",
    "coins": [{"ticker": "RGN", "value": "100.23"}],
    "msg": {/*your app-specific data*/}
}
```