//https://mp.weixin.qq.com/s/cBjk4jOjX7mdsXuS4r-l-g

package main

// GOOS=js GOARCH=wasm
import (
	"fmt"
	"syscall/js"
)

func main() {
	btn := js.Global().Get("document").Call("getElementById", "test")

	signal := make(chan int)

	var callback js.Func
	callback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println(this)
		fmt.Println(args)
		fmt.Println("button clicked")
		return nil
	})

	btn.Set("innerHTML", "changed by go")
	btn.Call("addEventListener", "click", callback)

	<-signal
}
