package extractor

import (
	"fmt"
	"regexp"
	"testing"
)

func TestExtractTitle(t *testing.T) {
	r := regexp.MustCompile(`\[comment\]: <> \(.*?\)`)
	w := r.FindAllString("abc [comment]: <> (self, test) abcd [comment]: <> (test)", -1)

	fmt.Println(w)
}
