package dump

import (
	"os"
	"path/filepath"
)

var (
	file = ".startdict.txt"
	// WordSeperator is a banner for words
	WordSeperator = []byte("========")
)

func OpenFile() (*os.File, error) {
	home := os.Getenv("HOME")
	absf := filepath.Join(home, file)
	return os.OpenFile(absf, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}
