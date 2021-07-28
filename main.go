package main

import (
	"git-manager/controllers"
	"os/exec"
	"regexp"

	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		input := ui.NewEntry()
		button := ui.NewButton("LOG")
		logs := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(logs, false)
		window := ui.NewWindow("Hello", 600, 600, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			out, err := exec.Command("git", "log").Output()
			if err != nil {
				panic(err)
			}

			ret := string(out)
			data := regexp.MustCompile("[ \n]").Split(string(ret), -1)
			commits := controllers.SplitCommits(data)

			i := 0

			var text string

			for i < len(commits) {
				text = text + "Commit HEX: " + commits[i].Hex + "\nAuthor: " + commits[i].Author + " " + commits[i].Email +
					"\nDate of commit: " + commits[i].Day + " " + commits[i].Month + " " + commits[i].Dint + " " + commits[i].Time + " " +
					commits[i].Year + "\nCommit message: " + commits[i].Message + "\n\n"
				i++
			}
			logs.SetText(text)
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
