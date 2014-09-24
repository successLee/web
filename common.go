package web

import (
	"os"
	"path/filepath"
	"strings"
)

func in(key, list string) bool {
	if key == "" || list == "" {
		return false
	}
	for _, i := range strings.Split(list, ",") {
		if key == i {
			return true
		}
	}
	return false
}

func muxPath(str string) string {
	str = strings.ToLower(str)
	n := len(str)
	if str != "/" && str[n-1] == '/' {
		str = str[:n-1]
	}
	return str
}

func getFilePathList(path, suffix string) ([]string, error) {
	filePathList := make([]string, 0)
	err := filepath.Walk(path, filepath.WalkFunc(func(filepath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.HasSuffix(filepath, suffix) {
			filePathList = append(filePathList, filepath)
		}
		return nil
	}))
	return filePathList, err
}
