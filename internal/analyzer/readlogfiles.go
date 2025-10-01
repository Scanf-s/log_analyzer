package analyzer

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"log_analyzer/internal/dto"
)

func ReadLogFiles(logPath string, files []os.DirEntry) ([]dto.LogEntry, error) {
	var entries []dto.LogEntry

	if len(files) == 0 {
		return nil, errors.New("no log files found")
	}

	for idx, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".log") {
			log.Printf("Reading %dth log file: %s\n", idx+1, file.Name())

			// 파일 전체 읽기
			filePath := filepath.Join(logPath, file.Name())
			fileEntries, err := processFile(filePath)
			if err != nil {
				log.Printf("Error processing file %s: %s\n", file.Name(), err)
				continue
			}
			entries = append(entries, fileEntries...) // 슬라이스 병합
		} else {
			continue
		}
	}

	return entries, nil
}

func processFile(filePath string) ([]dto.LogEntry, error) {
	var entries []dto.LogEntry

	logFile, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open log file %s: %s\n", filePath, err)
		return nil, err
	}
	defer func(logFile *os.File) { // 함수 종료시 파일 닫기
		_ = logFile.Close()
	}(logFile)

	scanner := bufio.NewScanner(logFile) // 한줄씩 읽어주는 스캐너
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	// Scanner의 기본 버퍼 크기는 64KB라서 로그 파일 한줄이 64KB를 넘으면 에러 발생
	// 따라서 충분히 큰 버퍼를 갖는 버퍼 설정

	for scanner.Scan() { // 파일 끝까지 한줄씩 읽으면서
		var entry dto.LogEntry
		line := strings.TrimSpace(scanner.Text()) // 앞뒤 공백 제거
		if line == "" {                           // 빈 줄이면 무시
			continue
		}

		err := json.Unmarshal([]byte(line), &entry) // JSON 파싱해서 LogEntry 구조체에 매핑
		if err != nil {
			log.Printf("Failed to parse log entry in file %s: %s\n", filePath, err)
			continue
		}
		entries = append(entries, entry) // 파싱된 로그 엔트리를 슬라이스에 추가
	}

	err = scanner.Err() // 스캐너 에러 체크
	if err != nil {
		log.Printf("Scanner error in file %s: %s\n", filePath, err)
	}

	return entries, nil // 모든 로그 엔트리 반환
}
