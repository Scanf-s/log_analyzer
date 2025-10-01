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
