// +build windows

package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/AllenDang/giu"
)

//go:embed sdelete64.exe
var sdelete64bytes []byte

var sdelete64path string

func shred(passes int, separate bool, name string, shredding *string) {
	stat, _ := os.Stat(name)
	if stat.IsDir() {
		// Walk the folder recursively
		filepath.Walk(name, func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			stat, _ := os.Stat(path)
			if stat.IsDir() {
				if stopShredding {
					return nil
				}

				t := 0
				files, _ := ioutil.ReadDir(path)
				for _, f := range files {
					if !f.IsDir() {
						t++
					}
				}
				shredDone += float32(t)
				shredUpdate(separate)
				shredding = strings.ReplaceAll(path, "\\", "/") + "/*"
				cmd := exec.Command(sdelete64path, "*", "-p", strconv.Itoa(passes))
				cmd.Dir = path
				cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				cmd.Run()
				giu.Update()
			}
			return nil
		})

		if !stopShredding {
			// sdelete64 doesn't delete the empty folder, so I'll do it manually
			os.RemoveAll(name)
		}
	} else {
		shredding = name
		cmd := exec.Command(sdelete64path, "*", "-p", strconv.Itoa(passes))
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()
		shredDone++
		shredUpdate(separate)
	}
}

func initializeShred() func() {
	// Create a temporary file to store sdelete64.exe
	sdelete64, _ := os.CreateTemp("", "sdelete64.*.exe")
	sdelete64path = sdelete64.Name()
	sdelete64.Write(sdelete64bytes)
	sdelete64.Close()
	cmd := exec.Command(s64deletepath, "/accepteula")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()
	return func() { os.Remove(sdelete64path) }
}
