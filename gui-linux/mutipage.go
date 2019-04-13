package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainWin *ui.Window

func makePage3() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	label := ui.NewLabel("其他费用(万元)")
	//hbox.Append(label, true)
	grid.Append(label,
		1, 0, 1, 1,
		true, ui.AlignCenter, true, ui.AlignCenter)

	grid.Append(ui.NewLabel("修理费"),
		0, 1, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)
	grid.Append(ui.NewEntry(),
		1, 1, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)
	grid.Append(ui.NewButton("估算"),
		2, 1, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)

	grid.Append(ui.NewLabel("其他管理费"),
		0, 2, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)
	grid.Append(ui.NewEntry(),
		1, 2, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)
	grid.Append(ui.NewButton("估算"),
		2, 2, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)

	grid.Append(ui.NewLabel("其他营业费"),
		0, 3, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)
	grid.Append(ui.NewEntry(),
		1, 3, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)
	grid.Append(ui.NewButton("估算"),
		2, 3, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)
	return vbox
}

func makePage2() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	label := ui.NewLabel("人员工资及福利费(万元)")
	grid.Append(label,
		0, 0, 1, 1,
		true, ui.AlignCenter, true, ui.AlignCenter)

	grid.Append(ui.NewButton("估算"),
		0, 1, 1, 1,
		true, ui.AlignCenter, true, ui.AlignCenter)

	rb := ui.NewRadioButtons()
	rb.Append("技术人员")
	rb.Append("管理人员")
	rb.Append("基础人员")
	grid.Append(rb,
		0, 2, 1, 1,
		true, ui.AlignCenter, true, ui.AlignCenter)
	return vbox

}

func makePage1() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	label := ui.NewLabel("初投资(万元)")
	grid.Append(label,
		1, 0, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)

	label = ui.NewLabel("土建费(万元)")
	entry := ui.NewEntry()

	grid.Append(label,
		0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)
	var attributedString string
	entry.OnChanged(func(entry *ui.Entry) {
		attributedString = entry.Text()
	})
	bt := ui.NewButton("终值")
	bt.OnClicked(func(button *ui.Button) {
		ui.MsgBox(mainWin, "test value", attributedString)
	})
	grid.Append(bt,
		2, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)
	bt = ui.NewButton("净现值")
	grid.Append(bt,
		3, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	label = ui.NewLabel("设备费(万元)")
	entry = ui.NewEntry()

	grid.Append(label,
		0, 2, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 2, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)
	bt = ui.NewButton("估值")
	grid.Append(bt,
		2, 2, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	rb := ui.NewRadioButtons()
	rb.Append("供能设备")
	rb.Append("电气设备")
	rb.Append("辅助设备")
	grid.Append(rb,
		1, 3, 1, 1,
		false, ui.AlignCenter, false, ui.AlignCenter)

	//entryForm := ui.NewForm()
	//entryForm.SetPadded(true)
	//
	//entryForm.Append("Entry", ui.NewEntry(), false)
	//grid.Append(entryForm,
	//	1, 4, 1, 1,
	//	false, ui.AlignCenter, false, ui.AlignCenter)
	//grid.Append(ui.NewButton("估算"),
	//	2, 4, 1, 1,
	//	false, ui.AlignCenter, false, ui.AlignCenter)

	return vbox
}

func makeBasicControlsPage() ui.Control {
	// 单选框
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	rb := ui.NewRadioButtons()
	rb.Append("全寿命周期冷热定价")
	rb.Append("任一年的冷热定价")
	vbox.Append(rb, false)
	vbox.Append(ui.NewHorizontalSeparator(), false)

	// 三个小页面
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	hbox.Append(makePage1(), true)
	hbox.Append(ui.NewVerticalSeparator(), false)
	hbox.Append(makePage2(), true)
	hbox.Append(ui.NewVerticalSeparator(), false)
	hbox.Append(makePage3(), true)
	vbox.Append(hbox, true)
	return vbox
}

func makeNumbersPage() ui.Control {

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	spinbox := ui.NewSpinbox(0, 100)
	slider := ui.NewSlider(0, 100)
	pbar := ui.NewProgressBar()
	spinbox.OnChanged(func(*ui.Spinbox) {
		slider.SetValue(spinbox.Value())
		pbar.SetValue(spinbox.Value())
	})
	slider.OnChanged(func(*ui.Slider) {
		spinbox.SetValue(slider.Value())
		pbar.SetValue(slider.Value())
	})
	vbox.Append(spinbox, false)
	vbox.Append(slider, false)
	vbox.Append(pbar, false)

	ip := ui.NewProgressBar()
	ip.SetValue(-1)
	vbox.Append(ip, false)

	group = ui.NewGroup("Lists")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	cbox := ui.NewCombobox()
	cbox.Append("Combobox Item 1")
	cbox.Append("Combobox Item 2")
	cbox.Append("Combobox Item 3")
	vbox.Append(cbox, false)

	ecbox := ui.NewEditableCombobox()
	ecbox.Append("Editable Item 1")
	ecbox.Append("Editable Item 2")
	ecbox.Append("Editable Item 3")
	vbox.Append(ecbox, false)

	rb := ui.NewRadioButtons()
	rb.Append("Radio Button 1")
	rb.Append("Radio Button 2")
	rb.Append("Radio Button 3")
	vbox.Append(rb, false)

	return hbox
}

func makeDataChoosersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	vbox.Append(ui.NewDatePicker(), false)
	vbox.Append(ui.NewTimePicker(), false)
	vbox.Append(ui.NewDateTimePicker(), false)
	vbox.Append(ui.NewFontButton(), false)
	vbox.Append(ui.NewColorButton(), false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	button := ui.NewButton("Open File")
	entry := ui.NewEntry()
	entry.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainWin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry.SetText(filename)
	})
	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	button = ui.NewButton("Save File")
	entry2 := ui.NewEntry()
	entry2.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.SaveFile(mainWin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry2.SetText(filename)
	})
	grid.Append(button,
		0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry2,
		1, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	msggrid := ui.NewGrid()
	msggrid.SetPadded(true)
	grid.Append(msggrid,
		0, 2, 2, 1,
		false, ui.AlignCenter, false, ui.AlignStart)

	button = ui.NewButton("Message Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBox(mainWin,
			"This is a normal message box.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	button = ui.NewButton("Error Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBoxError(mainWin,
			"This message box describes an error.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		1, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	return hbox
}

func setupUI() {
	mainWin = ui.NewWindow("linux gui ", 640, 400, true)
	mainWin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainWin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainWin.SetChild(tab)
	mainWin.SetMargined(true)

	tab.Append("固定成本", makeBasicControlsPage())
	tab.SetMargined(0, true)

	tab.Append("可变成本", makeNumbersPage())
	tab.SetMargined(1, true)

	tab.Append("营业收入", makeDataChoosersPage())
	tab.SetMargined(2, true)

	mainWin.Show()
}

func main() {
	err := ui.Main(setupUI)
	if err != nil {
		panic(err)
	}
}
