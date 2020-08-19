package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ilius/stardict/pkg/parser"
)

type Dictionary struct {
	*parser.Dictionary
}

func (d *Dictionary) Close() {
}

func LoadDictionary(basePath string) (*Dictionary, error) {
	ifoFile, err := os.Open(basePath + ".ifo")
	if err != nil {
		return nil, fmt.Errorf("failed to open .ifo: %w", err)
	}
	defer ifoFile.Close()
	idxFile, err := os.Open(basePath + ".idx")
	if err != nil {
		return nil, fmt.Errorf("failed to open .idx: %w", err)
	}
	defer idxFile.Close()
	var dict io.ReadCloser
	dict, err = os.Open(basePath + ".dict")
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to open .dict: %w", err)
		}
		dictDz, err := os.Open(basePath + ".dict.dz")
		if err != nil {
			return nil, fmt.Errorf("failed to open .dict.dz: %w", err)
		}
		defer dictDz.Close()
		dict, err = gzip.NewReader(dictDz)
		if err != nil {
			return nil, fmt.Errorf("error in gzip.NewReader: %w", err)
		}
	}
	if dict == nil {
		return nil, fmt.Errorf("dict == nil")
	}
	defer dict.Close()

	pdic, err := parser.LoadDictionary(ifoFile, idxFile, dict)
	if err != nil {
		return nil, err
	}
	return &Dictionary{
		Dictionary: pdic,
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
		dic, err := LoadDictionary(basePath)
		if err != nil {
			fmt.Printf("error while opening %v: %v\n", path, err)
			return nil
		}
		dicList = append(dicList, dic)
		return nil
	})
	return dicList, nil
}
