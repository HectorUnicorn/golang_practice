package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	INDENT = "    "
	INDENT_SEP = "│   "
	MIDDLE_INDICATOR = "├──"
	END_INDICATOR = "└──"
)

var (
	rootPath string
)

func init() {
	flag.StringVar(&rootPath, "p", "", "The path of target directory.")
}

func showFiles(basePath string, prefix string, middle_inds string, end_inds string, separator string, showAll bool) error {
	base, err := os.Open(basePath)
	// defer File.Close(base)
	if err != nil {
		return err
	}
	subs, err := base.Readdir(-1)
	if err != nil {
		return err
	}
	for i:=0; i<len(subs); i++ {
		v := subs[i]
		fi := v.(os.FileInfo)
		fp := fi.Name()
		if strings.HasPrefix(fp, ".") && !showAll {
			continue
		}
		if fi.IsDir() {
			absFp := path.Join(basePath, fp)
			if err != nil {
				return err
			}
			isLast := false
			if i == (len(subs) - 1) {
				fmt.Printf("%s\n", separator+end_inds+fp)
				isLast = true
			} else {
				fmt.Printf("%s\n", separator+middle_inds+fp)
			}
			sep := INDENT_SEP
			if isLast {
				sep = INDENT
			}
			err = showFiles(absFp, prefix, middle_inds, end_inds, separator+sep, showAll)
			if err != nil {
				return err
			}
		} else {
			if i == (len(subs) - 1) {
				fmt.Printf("%s\n", separator+end_inds+fp)
			} else {
				fmt.Printf("%s\n", separator+middle_inds+fp)
			}
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(rootPath) == 0 {
		defaultPath, err := os.Getwd()
		if err != nil {
			fmt.Println("GetwdError:", err)
		}
		rootPath = defaultPath
	}
	fmt.Printf("%s\n", rootPath)
	err := showFiles(rootPath, INDENT_SEP, MIDDLE_INDICATOR, END_INDICATOR, "", false)
	if err != nil {
		fmt.Println("showFilesError:", err)
	}
}
