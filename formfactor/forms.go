package formfactor

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func formatHTMLBreaks(z string) string {
	return strings.Replace(z, "\n", "<br>", -1)
}
func ConstructEmail(x string) string {
	htfirst, err := ioutil.ReadFile("index1.txt")
	if err != nil {
		fmt.Println(err)
		return "Error Parsing HTML - PT1"
	}
	partOne := string(htfirst)
	htsecond, err2 := ioutil.ReadFile("index2.txt")
	if err != nil {
		fmt.Println(err2)
		return "Error Parsing - HTML - PT2"
	}
	partTwo := string(htsecond)
	z := formatHTMLBreaks(x)
	return partOne + z + partTwo
}
