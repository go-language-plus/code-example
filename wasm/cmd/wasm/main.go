package main

import "fmt"

// GOOS=js GOARCH=wasm go build -o ../../assets/main.wasm
func main() {
	fmt.Println("Hello, WebAssembly!")
}
