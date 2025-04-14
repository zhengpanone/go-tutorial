//go:build js && wasm
// +build js,wasm

package main

import (
	_ "embed"
	"errors"
	"syscall/js"
	"wasm_package/utils"
)

//go:embed info.txt
var info []byte

func Readme(this js.Value, args []js.Value) any {
	return string(info)
}

func Add(this js.Value, args []js.Value) string {
	if len(args) == 0 {
		return utils.PackageResult(false, errors.New("参数不能为空!"), nil)
	}
	x := args[0].Float()
	y := args[1].Float()
	return utils.PackageResult(true, nil, x+y)
}

func Sub(this js.Value, args []js.Value) string {
	if len(args) == 0 {
		return utils.PackageResult(false, errors.New("参数不能为空!"), nil)
	}
	x := args[0].Float()
	y := args[1].Float()
	return utils.PackageResult(true, nil, x-y)

}

func Multi(this js.Value, args []js.Value) string {
	if len(args) == 0 {
		return utils.PackageResult(false, errors.New("参数不能为空!"), nil)
	}
	x := args[0].Float()
	y := args[1].Float()
	return utils.PackageResult(true, nil, x*y)

}

func Div(this js.Value, args []js.Value) string {
	if len(args) == 0 {
		return utils.PackageResult(false, errors.New("参数不能为空!"), nil)
	}
	x := args[0].Float()
	y := args[1].Float()
	return utils.PackageResult(true, nil, x/y)

}

func Mod(this js.Value, args []js.Value) string {
	if len(args) == 0 {
		return utils.PackageResult(false, errors.New("参数不能为空!"), nil)
	}
	x := args[0].Int()
	y := args[1].Int()
	return utils.PackageResult(true, nil, x%y)
}

func main() {
	js.Global().Set("readme", js.FuncOf(Readme))
	js.Global().Set("add", js.FuncOf(Add))
	// 阻塞挂起
	<-make(chan bool)
}
