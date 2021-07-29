package views

import (
	"git-manager/controllers"
	"strings"

	"github.com/andlabs/ui"
)

func CommitsView() error {
	return ui.Main(func() {
		page := 0
		input := ui.NewEntry()
		nextPage := ui.NewButton("Next Page")
		prePage := ui.NewButton("Previous Page")
		button := ui.NewButton("Log repo from path")
		logs := ui.NewLabel("")
		errors := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter path to repo:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(prePage, false)
		box.Append(nextPage, false)
		box.Append(logs, false)
		box.Append(errors, false)
		window := ui.NewWindow("Git Manager", 800, 600, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			// logs = controllers.LogCommits(logs, input.Text())
			data := controllers.LogCommits(logs, input.Text())

			logs.SetText(strings.Join(data[page], ""))

			// greeting.SetText("Commit HEX: " + data[1] + "\nAuthor: " + data[3] + "\nEmail: " + data[4] + "\nDate of commit: " +
			// 	data[8] + " " + data[9] + " " + data[10] + " " + data[11] + " " + data[12] + "\nMessage: " + strings.Join(data[19:], " "))
		})
		nextPage.OnClicked(func(*ui.Button) {
			errors.SetText("")
			data := controllers.LogCommits(logs, input.Text())
			if page == len(data)-1 {
				errors.SetText("You've reached last page")
			} else {
				page++
			}
			logs.SetText(strings.Join(data[page], ""))
		})
		prePage.OnClicked(func(*ui.Button) {
			errors.SetText("")
			data := controllers.LogCommits(logs, input.Text())
			if page == 0 {
				errors.SetText("You've reached first page")
			} else {
				page--
			}
			logs.SetText(strings.Join(data[page], ""))
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
}
