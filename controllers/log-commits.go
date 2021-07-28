package controllers

import (
	"os/exec"
	"regexp"

	"github.com/andlabs/ui"
)

func LogCommits(logs *ui.Label) *ui.Label {
	out, err := exec.Command("git", "log").Output()
	if err != nil {
		panic(err)
	}

	ret := string(out)
	data := regexp.MustCompile("[ \n]").Split(string(ret), -1)
	commits := splitCommits(data)

	i := 0

	var text string

	for i < len(commits) {
		text = text + "Commit HEX: " + commits[i].Hex + "\nAuthor: " + commits[i].Author + " " + commits[i].Email +
			"\nDate of commit: " + commits[i].Day + " " + commits[i].Month + " " + commits[i].Dint + " " + commits[i].Time + " " +
			commits[i].Year + "\nCommit message: " + commits[i].Message + "\n\n"
		i++
	}
	logs.SetText(text)

	return logs
}
