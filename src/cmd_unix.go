// +build linux darwin

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"

	"github.com/AllenDang/giu"
)

func shred(passes int, separate bool, name string, shredding *string) {
	stat, _ := os.Stat(name)
	if stat.IsDir() {
		var coming []string

		// Walk the folder recursively
		filepath.Walk(name, func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if stopShredding {
				return nil
			}
			stat, _ := os.Stat(path)
			if !stat.IsDir() {
				if len(coming) == 128 {
					// Use a WaitGroup to parallelize shredding
					var wg sync.WaitGroup
					for i, j := range coming {
						wg.Add(1)
						go func(wg *sync.WaitGroup, id int, j string) {
							defer wg.Done()
							runShredCommand(j, separate)
							giu.Update()
						}(&wg, i, j)
					}
					wg.Wait()
					coming = nil
				} else {
					coming = append(coming, path)
				}
			}
			return nil
		})
		for _, i := range coming {
			if stopShredding {
				break
			}
			go func(i string) {
				runShredCommand(i, separate)
				giu.Update()
			}(i)
		}
		if !stopShredding {
			os.RemoveAll(name)
		}
	} else { // The path is a file, not a directory, so just shred it
		runShredCommand(name, separate)
	}
}

func runShredCommand(name string, separate bool) {
	shredding = name
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command("shred", "-ufvz", "-n", strconv.Itoa(int(shredPasses)), name)
	} else {
		cmd = exec.Command("rm", "-rfP", name)
	}
	cmd.Run()
	shredDone++
	shredUpdate(separate)
}

func initializeShred() func() {
	return func() {}
}
