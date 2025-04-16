package logger

const (
	AppLogger         = "AppLogger"
	KeyHttpRequest    = "httpRequest"
	KeyRequestMethod  = "requestMethod"
	KeyRequestUrl     = "requestUrl"
	KeyRemoteIp       = "remoteIp"
	KeyUserAgent      = "userAgent"
	KeyReferer        = "referer"
	KeyProtocol       = "protocol"
	KeyLatency        = "latency"
	KeyStatus         = "status"
	KeyUser           = "user"
	sourceLocationKey = "logging.googleapis.com/sourceLocation"
	attrSeverity      = "severity"
	attrMessage       = "message"
	levelDebug        = "DEBUG"
	levelInfo         = "INFO"
	levelWarn         = "WARNING"
	levelError        = "ERROR"
	levelDefault      = "DEFAULT"
)
