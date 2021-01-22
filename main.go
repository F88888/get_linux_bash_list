package get_linux_bash_list

import (
	"os/exec"
	"regexp"
)

type BashList struct {
	Err  error
	List []string
}

func GetBash() (bashList BashList) {
	// init
	var bashText []byte
	var bashRegexp = regexp.MustCompile(" (\\w{,20}) \\[[\\-0-9a-zA-Z]")
	if bashText, bashList.Err = exec.Command("help -d").Output(); bashList.Err == nil {
		// Regular Expression Matching
		bashList.List = bashRegexp.FindAllString(string(bashText), -1)
	}
	return bashList
}

func (e BashList) GetList() ([]string, error) {
	return e.List, e.Err
}

func (e BashList) InBash(text string) (bool, error) {
	isTrue := false
	if e.Err == nil {
		for _, v := range e.List {
			if text == v {
				isTrue = true
				break
			}
		}
	}
	return isTrue, e.Err
}
