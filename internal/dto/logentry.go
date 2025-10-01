package dto

type LogEntry struct {
	Timestamp string `json:"timestamp"` // 로그 발생 시간
	Level     string `json:"level"`     // 로그 레벨 (INFO, WARN, ERROR, DEBUG, ...)
	Service   string `json:"service"`   // 어떤 서비스에 대한 로그인지
	Message   string `json:"message"`   // 로그 메세지
}
