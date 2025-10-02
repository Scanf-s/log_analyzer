package fileservice

import (
	"errors"
	"os"
	"strings"
)

func CheckDirNotEmpty(logPath string) ([]os.DirEntry, error) {
	var fileCount = 0
	files, err := os.ReadDir(logPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".log") { // 디렉토리가 아니면서 .log 파일만 카운트
			fileCount++
		}
	}
	if fileCount == 0 {
		return nil, errors.New("no log files found")
	}

	return files, nil
}
