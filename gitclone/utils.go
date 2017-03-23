package gitclone

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func subString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dir string) string {
	return subString(dir, 0, strings.LastIndex(dir, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func existFile(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func getGitRepoDirList(filePath string) (list []string) {
	filepath.Walk(filePath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil || err != nil {
				return err
			} else if f.IsDir() &&
				f.Name() == ".git" {
				list = append(list, getParentDirectory(path))
			}
			return nil
		})
	return
}

func isGitRepoDir(filePath string) bool {
	f, err := os.Stat(filePath)
	if f == nil {
	} else if err != nil && !os.IsExist(err) {
	} else if f.IsDir() {
		return existFile(path.Join(filePath, ".git"))
	}
	return false
}
