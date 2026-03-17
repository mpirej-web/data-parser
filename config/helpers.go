package data_parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func getExecutablePath() (string, error) {
	exe, err := os.Executable()
	if err!= nil {
		return "", err
	}
	return filepath.Abs(exe)
}

func getExecutableDir() string {
	exe, err := os.Executable()
	if err!= nil {
		log.Fatal(err)
	}
	return filepath.Dir(exe)
}

func tablePrinter(data [][]string) {
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	tw.SetHeader([]string{"Column 1", "Column 2", "Column 3"})
	tw.Append(data)
	tw.Render()
}

func getCallerInfo(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%s:%d", file, line)
}

func trimStr(s string) string {
	return strings.TrimSpace(s)
}

func isDirEmpty(dir string) bool {
	files, err := os.ReadDir(dir)
	if err!= nil {
		return false
	}
	return len(files) == 0
}