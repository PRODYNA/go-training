package snip

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type snipper struct {
	out string
	in  string
}

func New(out string, in string) *snipper {
	return &snipper{
		out: out,
		in:  in,
	}
}
func (s snipper) Snip() {

	if _, err := os.Stat(s.out); os.IsNotExist(err) {
		err = os.Mkdir(s.out, os.ModePerm)
		if err != nil {
			log.Fatal(fmt.Errorf("could not create %s: %v", s.out, err))
		}
	}

	err := CopyDirectory(s.in, s.out)
	if err != nil {
		log.Fatal(fmt.Errorf("could not copy %s to %s: %v", s.in, s.out, err))
	}

	entries, err := ioutil.ReadDir(s.out)
	if err != nil {
		log.Fatal(fmt.Errorf("could not read directory %s: %v", s.out, err))
	}

	for _, entry := range entries {
		removeSnippets(s.out, entry)
	}
}

func removeSnippets(path string, entry os.FileInfo) string {
	p := filepath.Join(path, entry.Name())
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(fmt.Errorf("could not read directory %s: %v", path, err))
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	skip := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "SNIP START") {
			skip = true
		}
		if strings.Contains(line, "SNIP END") {
			skip = false
			continue
		}

		if skip {
			continue
		}
		lines = append(lines, fmt.Sprintf("%s\n", line))
	}
	return strings.Join(lines, "")
}
