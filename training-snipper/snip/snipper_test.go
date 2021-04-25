package snip

import (
	"os"
	"testing"
)

func Test_removeSnippets(t *testing.T) {
	fi, _ := os.Stat("./test/test.txt")
	lines := removeSnippets("./test", fi)
	if lines != "line1\n\n\nline2\n" {
		t.Fatalf("line is wrong %s", lines)
	}
}
