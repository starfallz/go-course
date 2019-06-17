package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const minNumberOfLetters = 3

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	var result []string
	for _, val := range vals {
		result = append(result, convertString(val))
	}
	return result
}

func convertString(val string) string {
	stringLength := len(val)
	if stringLength > minNumberOfLetters {
		return strings.Join([]string{val[:1], strconv.Itoa(stringLength - 2), val[stringLength-1 : stringLength]}, "")
	}
	return val
}
