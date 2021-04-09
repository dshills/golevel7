package commons

import (
	"bufio"
	"fmt"
	"io"
	"path/filepath"
	"regexp"
	"strings"
)

var hl7SplitToken = regexp.MustCompile("(\\r(\\n|\\x1c)+(\\n\\r)?MSH\\|\\^\\~\\\\\\&\\||$)")
var hl7FindStartToken = regexp.MustCompile("(MSH\\|\\^\\~\\\\\\&\\||$)")

const scanBufferSize = 10 * 1024 * 1024

// GetHl7Files finds all hl7 files in the current directory and returns the file names as a slice of strings
func GetHl7Files() (matches []string, err error) {
	pattern := "*.hl7"
	fileCnt := 0
	fmt.Println("")
	if matches, err = filepath.Glob(pattern); err == nil {
		for fileCnt, _ = range matches {
			fileCnt++
			if fileCnt == 1 || fileCnt%1000 == 0 {
				fmt.Printf("\rfound %v", fileCnt)
			}
		}
	}
	if fileCnt != 1 && fileCnt%1000 != 0 {
		fmt.Printf("\rfound %v", fileCnt)
	}
	fmt.Println("")
	return matches, err
}

// CrLfSplit implements a split function to deal with hl7 messages in a file, terminated by cr/lf and an optional second lf
func crLfSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 { // end of file
	} else {
		loc := hl7SplitToken.FindIndex(data) // found record delimiter
		if loc != nil || atEOF {
			nextLoc := hl7FindStartToken.FindIndex(data[1:])
			hl7RecPatch := []byte(strings.ReplaceAll(string(data[0:loc[0]]), "\r\n", "\r")) // cr/lf found after each segment, patch.
			return nextLoc[0] + 1, hl7RecPatch, nil
		}
	}
	return advance, token, err // no cr/lf found, either the end or get bigger data and look again
}

func NewBufScanner(r io.Reader) *bufio.Scanner {
	b := bufio.NewScanner(r)
	buf := make([]byte, scanBufferSize)
	b.Buffer(buf, scanBufferSize)
	b.Split(crLfSplit)
	return b
}
