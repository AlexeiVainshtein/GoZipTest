package main

import (
	"GoZipTest/Test"
	"fmt"
	"github.com/mholt/archiver"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	Test.Print(wd)
	var list []string
	filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if !info.IsDir() {
			return err
		}
		list = append(list, path)
		return nil
	})
	fmt.Println(list)
	archiver.Zip.Make("zipping.zip", []string{wd})
}
