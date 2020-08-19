package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dicDir := path.Join(homeDir, ".stardict", "dic")
	dicList, err := Open(dicDir)
	defer dicList.Close()
	if err != nil {
		panic(err)
	}
	for _, word := range os.Args[1:] {
		transCount := 0
		for _, dic := range dicList {
			transList := dic.GetFormatedMeaning(word)
			if len(transList) == 0 {
				continue
			}
			if transCount > 0 {
				fmt.Printf("\n")
			}
			transCount += len(transList)
			fmt.Printf(
				"--> query %#v from %s\n",
				word,
				dic.GetBookName(),
			)
			for _, trans := range transList {
				fmt.Println(trans)
			}
		}
	}
}
