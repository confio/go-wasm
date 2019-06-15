## How to compile wasm

### Requirements

```
rustup update
rustup default nightly
rustup target add wasm32-unknown-unknown
```

### Compile example wasm

```shell
cd examples/greet
sh ./build.sh
ls -l build
```

### Run go tests

From top level, call `go test -v .`

### Cosmos-wasm interface

`init` or `receive`
```json
{
    "signer": "hex addr",
    "coins": [{"ticker": "RGN", "value": "100.23"}],
    "msg": {/*your app-specific data*/}
}
```