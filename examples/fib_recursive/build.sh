#!/bin/bash

rm -r build || true
mkdir build

cargo build --release

find ./target/wasm32-unknown-unknown/release/ -name "*.wasm" -exec cp "{}" build/ ";"
