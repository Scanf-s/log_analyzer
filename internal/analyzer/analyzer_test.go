package analyzer

import (
	"log_analyzer/internal/dto"
	"testing"
)

func TestLineCounter(t *testing.T) {
	// Given
	var entries []dto.LogEntry
	entries = append(entries, dto.LogEntry{})
	entries = append(entries, dto.LogEntry{})
	entries = append(entries, dto.LogEntry{})

	// When
	lineCount := LineCounter(entries)

	// Then
	if lineCount != len(entries) {
		t.Errorf("Expected line count to be 3, but got %d", lineCount)
	}
}
