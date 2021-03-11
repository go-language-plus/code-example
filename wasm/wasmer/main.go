package main

import (
	"fmt"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	// 手动编写 WAT 文本
	wasmBytes := []byte(`
	(module
	  (type (func (param i32 i32) (result i32)))
	  (func (type 0)
	    local.get 0
	    local.get 1
	    i32.add)
	  (export "sum" (func 0)))
`)

	// 创建 wasmer 引擎
	engine := wasmer.NewEngine()

	// 创建一个存储空间
	store := wasmer.NewStore(engine)

	// 编译 wasm 模块
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		fmt.Println("Failed to compile module:", err)
	}

	// 导入一个空的导入对象
	importObject := wasmer.NewImportObject()

	// 初始化 WebAssembly 模块到对象中
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		panic(fmt.Sprintln("Failed to instantiate the module:", err))
	}

	// 从对象中导出定义的函数
	sum, err := instance.Exports.GetFunction("sum")
	if err != nil {
		panic(fmt.Sprintln("Failed to get the `add_one` function:", err))
	}

	// 使用导出的函数
	result, err := sum(1, 2)
	if err != nil {
		panic(fmt.Sprintln("Failed to call the `add_one` function:", err))
	}

	// 打印结果
	fmt.Println("Results of `sum`:", result)

	// Output:
	// Results of `sum`: 3
}
