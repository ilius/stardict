package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ilius/stardict/pkg/parser"
)

type Dictionary struct {
	*parser.Dictionary

	ifoFile  *os.File
	idxFile  *os.File
	dictFile *os.File
}

func (d *Dictionary) Close() {
	d.ifoFile.Close()
	d.idxFile.Close()
	d.dictFile.Close()
}

func NewDictionary(basePath string) (*Dictionary, error) {
	ifoFile, err := os.Open(basePath + ".ifo")
	if err != nil {
		return nil, err
	}
	idxFile, err := os.Open(basePath + ".idx")
	if err != nil {
		return nil, err
	}
	dictFile, err := os.Open(basePath + ".dict.dz")
	if err != nil {
		return nil, err
	}
	pdic, err := parser.NewDictionary(ifoFile, idxFile, dictFile)
	if err != nil {
		return nil, err
	}
	return &Dictionary{
		Dictionary: pdic,
		ifoFile:    ifoFile,
		idxFile:    idxFile,
		dictFile:   dictFile,
	}, nil
}

type DictionaryList []*Dictionary

func (l DictionaryList) Close() {
	for _, dic := range l {
		dic.Close()
	}
}

func Open(rootPath string) (DictionaryList, error) {
	dicList := DictionaryList{}
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".ifo") {
			return nil
		}
		basePath := path[:len(path)-4]
		// if os.Stat(basePath + ".dict.dz")
		dic, err := NewDictionary(basePath)
		if err != nil {
			fmt.Printf("error while opening %v: %v\n", path, err)
			return nil
		}
		dicList = append(dicList, dic)
		return nil
	})
	return dicList, nil
}
