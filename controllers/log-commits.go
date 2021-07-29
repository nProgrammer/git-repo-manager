package controllers

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/andlabs/ui"
)

func LogCommits(logs *ui.Label, path string) [][]string {
	os.Chdir(path)
	out, err := exec.Command("git", "log").Output()
	if err != nil {
		panic(err)
	}

	ret := string(out)
	data := regexp.MustCompile("[ \n]").Split(string(ret), -1)
	commits := splitCommits(data)

	i := 0

	var text []string

	for i < len(commits) {
		text = append(text, "Commit HEX: "+commits[i].Hex+"\nAuthor: "+commits[i].Author+" "+commits[i].Email+
			"\nDate of commit: "+commits[i].Day+" "+commits[i].Month+" "+commits[i].Dint+" "+commits[i].Time+" "+
			commits[i].Year+"\nCommit message: "+commits[i].Message+"\n\n")
		i++
	}

	var cmts [][]string

	var cmt []string

	i = 0  // text[i]
	j := 0 // cmts[][j]
	for i < len(text) {
		if j < 5 {
			cmt = append(cmt, text[i])
			j++
			i++
			if i == len(text) {
				cmts = append(cmts, cmt)
			}
		} else if j == 5 {
			j = 0
			cmts = append(cmts, cmt)
		} else {
			fmt.Println("ERROR")
		}
	}
	return cmts
}
