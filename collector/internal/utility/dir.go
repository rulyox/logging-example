package utility

import (
	"os"
	"path"
	"path/filepath"
	"sync"
)

var currentDir string
var currentDirOnce sync.Once

func CurrentDir() string {
	currentDirOnce.Do(func() {
		exec, _ := os.Executable()
		execDir := filepath.Dir(exec)
		currentDir = path.Join(execDir, "..")
	})
	return currentDir
}
