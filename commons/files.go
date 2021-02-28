package commons

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var hl7SplitToken = regexp.MustCompile("\\r\\n\\n?")

// GetHl7Files finds all hl7 files in the current directory and returns the file names as a slice of strings
func GetHl7Files() ([]string, error) {
	var matches []string
	pattern := "*.hl7"

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	} else if len(matches) == 0 {
		return nil, fmt.Errorf("No files found")
	}
	return matches, nil
}

// CrLfSplit implements a split function to deal with hl7 messages in a file, terminated by cr/lf and an optional second lf
func CrLfSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 { // end of file
	} else {
		loc := hl7SplitToken.FindIndex(data) // found record delimiter
		if loc != nil || atEOF {
			return loc[1], data[0:loc[0]], nil
		}
	}
	return advance, token, err // no cr/lf found, either the end or get bigger data and look again
}
