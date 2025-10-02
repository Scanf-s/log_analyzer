package analyzer

import (
	"log_analyzer/internal/dto"
)

func LineCounter(entries []dto.LogEntry) int {
	// 로그 엔트리 개수 반환하는 함수
	// 현재는 간단하게 len()함수 사용하는데
	// 추후 필요 시 복잡한 로직 추가 가능
	return len(entries)
}

func LineCountByDate(entries []dto.LogEntry) map[string]int {
	// 일자별 로그 개수 세는 함수
	lineCountByDate := map[string]int{}

	for _, entry := range entries {
		timestamp := []rune(entry.Timestamp)
		lineCountByDate[string(timestamp[:10])]++
	}

	return lineCountByDate
}

func TotalLogLevelStats(entries []dto.LogEntry) (map[string]int, map[string]float32) {
	// 로그 레벨 별 개수 및 비율 계산 함수
	logLevelCount := map[string]int{}
	logLevelRatio := map[string]float32{}
	totalCount := LineCounter(entries)

	// 로그 레벨 별 개수 세기
	for _, entry := range entries {
		logLevelCount[entry.Level]++
	}

	// 비율 계산
	for level, count := range logLevelCount {
		logLevelRatio[level] = float32(count) / float32(totalCount)
	}

	return logLevelCount, logLevelRatio
}

func LogStatsByDateAndLevel(entries []dto.LogEntry) []map[string]interface{} {
	// 일자별 로그 레벨별 개수 및 비율 계산

	// 전체 엔트리 읽어서 일자별로 맵 만들기
	logByDate := map[string][]dto.LogEntry{}
	for _, entry := range entries {
		logByDate[entry.Timestamp[:10]] = append(logByDate[entry.Timestamp[:10]], entry)
	}

	// 일자별로 로그 레벨 통계 계산
	var logLevelStatsByDate []map[string]interface{}
	for date, entries := range logByDate {
		logLevelCount, logLevelRatio := TotalLogLevelStats(entries)
		for level, count := range logLevelCount {
			logLevelStatsByDate = append(logLevelStatsByDate, map[string]interface{}{
				"date":  date,
				"count": count,
				"level": level,
				"ratio": logLevelRatio[level] * 100,
			})
		}
	}

	return logLevelStatsByDate
}

func LogStatsByService(entries []dto.LogEntry) map[string]float32 {
	// 서비스 별 로그 발생 비율 계산 함수
	var serviceCount = map[string]int{}
	var logByService = map[string]float32{}
	totalCount := LineCounter(entries)

	for _, entry := range entries {
		serviceCount[entry.Service]++
	}

	for serviceName, count := range serviceCount {
		logByService[serviceName] = float32(count) / float32(totalCount) * 100
	}

	return logByService
}
