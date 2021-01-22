package get_linux_bash_list

import (
	"os/exec"
	"regexp"
)

type BashList struct {
	List []string
}

func GetBash() (bashList BashList, err error) {
	// init
	var bashText []byte
	var bashRegexp = regexp.MustCompile(" (\\w{,20}) \\[[\\-0-9a-zA-Z]")
	if bashText, err = exec.Command("help -d").Output(); err == nil {
		// Regular Expression Matching
		bashList.List = bashRegexp.FindAllString(string(bashText), -1)
	}
	return bashList, err
}

func (e BashList) GetList() []string {
	return e.List
}

func (e BashList) InBash(text string) bool {
	for _, v := range e.List {
		if text == v {
			return true
		}
	}
	return false
}
