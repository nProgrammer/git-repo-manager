package controllers

import (
	"fmt"
	"git-manager/models"
	"regexp"
	"strings"
)

func SplitCommits(data []string) []models.Commit {
	a := regexp.MustCompile("commit").Split(strings.Join(data, " "), -1)
	var commits []models.Commit
	i := 1
	for i < len(a) {
		var commit models.Commit
		b := regexp.MustCompile("[ \n]").Split(string(a[i]), -1)
		commit.Hex = b[1]
		commit.Author = b[3]
		commit.Email = b[4]
		commit.Day = b[7]
		commit.Month = b[8]
		commit.Dint = b[9]
		commit.Time = b[10]
		commit.Year = b[11]
		commit.Message = strings.Join(b[18:], "")
		commits = append(commits, commit)
		i++
	}
	fmt.Println(commits)
	return commits
}
