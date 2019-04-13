package main

import (
	"fmt"
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		name := ui.NewEntry()
		button := ui.NewButton("测试")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("输入姓名："), false)
		box.Append(name, false)
		box.Append(button, false)
		box.Append(greeting, false)

		//创建window窗口,并设置长度
		window := ui.NewWindow("first application.", 600, 500, false)
		window.SetChild(box)
		button.OnClicked(func(button *ui.Button) {
			fmt.Println("get name:", name.Text())
			greeting.SetText("Hello," + name.Text() + "!")
		})
		window.OnClosing(func(window *ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
