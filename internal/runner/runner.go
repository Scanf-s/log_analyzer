package runner

import (
	"log"
	"log_analyzer/internal/analyzer"
	"log_analyzer/internal/fileservice"
)

func Run(logPath string) {

	// 1. LogPath 유효성 검사
	files, err := fileservice.CheckDirNotEmpty(logPath)
	if err != nil {
		log.Fatal(err)
	}

	// 2. JSON 파일 읽어서 데이터구조 생성
	entries, err := fileservice.ReadLogFiles(logPath, files)
	if err != nil {
		log.Fatal(err)
	}

	// 3. 총 로그 개수 계산
	lineCount := analyzer.LineCounter(entries)
	log.Printf("Total lines: %d", lineCount)

	// 4. 일자별 로그 개수 계산
	lineCountByDate := analyzer.LineCountByDate(entries)
	for date, count := range lineCountByDate {
		log.Printf("Date: %s, Lines: %d", date, count)
	}

	// 5. 전체 로그 레벨 별 개수 및 비율 계산
	logLevelCount, logLevelRatio := analyzer.TotalLogLevelStats(entries)
	for level, count := range logLevelCount {
		log.Printf("Level: %s, Count: %d, Ratio: %f%%", level, count, logLevelRatio[level]*100)
	}

	// 6. 일자 별 로그 레벨 별 개수 및 비율 계산
	logStatsByDateAndLevel := analyzer.LogStatsByDateAndLevel(entries)
	for _, stats := range logStatsByDateAndLevel {
		log.Printf("Date: %s\tLevel: %s\tCount: %d\t Ratio: %f\t", stats["date"], stats["level"], stats["count"], stats["ratio"])
	}

	// 7. 서비스 별 로그 발생 비율 계산 함수
	logStatsByService := analyzer.LogStatsByService(entries)
	for service, ratio := range logStatsByService {
		log.Printf("Service: %s\tRatio: %f%%\t", service, ratio)
	}
}
