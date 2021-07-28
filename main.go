package main

import (
	"git-manager/controllers"

	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		input := ui.NewEntry()
		button := ui.NewButton("Log repo from path")
		logs := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter path to repo:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(logs, false)
		window := ui.NewWindow("Git Manager", 600, 600, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			logs = controllers.LogCommits(logs, input.Text())
			// greeting.SetText("Commit HEX: " + data[1] + "\nAuthor: " + data[3] + "\nEmail: " + data[4] + "\nDate of commit: " +
			// 	data[8] + " " + data[9] + " " + data[10] + " " + data[11] + " " + data[12] + "\nMessage: " + strings.Join(data[19:], " "))
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
