package read

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkIsProp(a []string) {
	if len(a) == 1 {
		fmt.Fprintf(os.Stderr, "Not valid prop: %s", a[0])
		os.Exit(1)
	}
}

func parse(line string) (string, string) {
	var ws = "\t "
	s := regexp.MustCompile(" ?: ?| ?= ?|\\s").Split(strings.TrimLeft(line, ws), 2)
	checkIsProp(s)
	k := strings.TrimLeft(s[0], ws)
	v := strings.TrimLeft(s[1], ws)
	return k, v
}

func skip(line string) bool {
	if len(line) == 0 {
		return true
	}

	comment, _ := regexp.MatchString(".*(#|!).*", line)
	if comment {
		return true
	}
	return false
}

func Props(file string) map[string]string {
	f, err := os.Open(file)
	defer f.Close()
	check(err)

	props := make(map[string]string)
	s := bufio.NewScanner(f)
	for s.Scan() {
		if skip(s.Text()) {
			continue
		}
		k, v := parse(s.Text())
		props[k] = v
	}
	return props
}