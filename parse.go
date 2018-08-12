package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	cmdRe   = regexp.MustCompile("(-?[0-9.]+)\\s*(.+)\\s+in\\s+(.+)")
	optRe   = regexp.MustCompile("-([a-z]+)")
	stripRe = regexp.MustCompile(",")
)

func parseCmd(s string) (amount float64, fromStr, toStr string, err error) {
	s = stripRe.ReplaceAllString(s, "")
	mg := cmdRe.FindStringSubmatch(s)
	if len(mg) != 4 {
		return amount, fromStr, toStr, fmt.Errorf("incomplete command")
	}
	fromStr = mg[2]
	toStr = mg[3]

	q, err := strconv.ParseFloat(mg[1], 6)
	if err != nil {
		return amount, fromStr, toStr, fmt.Errorf("failed to parse value: %s", mg[1])
	}
	amount = q

	return amount, fromStr, toStr, nil
}

// parse out options from a given string
func parseOpts(s string) (ns string, opts []string) {

	matches := optRe.FindAllString(s, -1)
	for _, m := range matches {
		for _, ch := range m[1:] {
			opts = append(opts, string(ch))
		}
		s = strings.Replace(s, m, "", 1)
	}

	return strings.Trim(s, " "), opts
}
