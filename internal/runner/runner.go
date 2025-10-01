package runner

import (
	"errors"
	"log"
	"os"
	"strings"

	"log_analyzer/internal/analyzer"
)

func Run(logPath string) {

	// 1. LogPath 유효성 검사
	files, err := CheckDirNotEmpty(logPath)
	if err != nil {
		log.Fatal(err)
	}

	// 2. JSON 파일 읽어서 데이터구조 생성
	entries, err := analyzer.ReadLogFiles(logPath, files)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		log.Printf("%+v\n", entry)
	}

	lineCount := analyzer.LineCounter(entries)
	log.Printf("Total lines: %d", lineCount)
}

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
