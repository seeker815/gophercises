package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type sampleFile struct {
	fp   string
	name string
}

func main() {
	// Get wd and sample path
	wdPath, err := os.Getwd()
	check(err, "Failed to read current dir path")
	samplePath := filepath.Join(wdPath, "sample")

	dirList, e := iterateDir(samplePath)
	check(e, "failed to iterate through sample path and list files")

	renameFiles(dirList)
}

func iterateDir(fp string) ([]sampleFile, error) {
	// Iterate through the samplePath
	fileList := []sampleFile{}
	e := filepath.Walk(fp, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileList = append(fileList, sampleFile{
				fp:   path,
				name: info.Name(),
			})
		}
		return nil
	})
	if e != nil {
		return nil, e
	}
	return fileList, nil
}

func renameFiles(fileList []sampleFile) {
	r, _ := regexp.Compile("birthday_00[1-9].txt")
	for _, file := range fileList {
		if r.MatchString(file.name) {
			fmt.Println(file.fp)
			fs := strings.Split(strings.SplitAfter(file.name, "_")[1], ".")
			ns := fs[0] + "-" + "birthday.txt"

			np := r.ReplaceAllString(file.fp, ns)
			err := os.Rename(file.fp, np)
			if err != nil {
				check(err, fmt.Sprintf("Failed to rename the file: %s", file.fp))
			}
		}
	}
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
