## WebAssembly

```bash
cd cmd/wasm

# 编译 .wasm
GOOS=js GOARCH=wasm go build -o ../../assets/main.wasm

# Copy go 官方给我们准备好的胶水 JS 文件
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ../../assets/
```

启动 HTTP 服务：
```bash
go run cmd/server/main.go -dir assets/
```
打开 `127.0.0.1:2345` 查看效果。

也可以直接用 goexec 来启动服务：
```bash
# install goexec: go get -u github.com/shurcooL/goexec
goexec 'http.ListenAndServe(`:2345`, http.FileServer(http.Dir(`assets/`)))'
```
