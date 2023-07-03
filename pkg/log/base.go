package log

// var _ElementPrefix string = "&(E*#@"
// var _ElementSuffix string = "#*)"
// 全局变量：日志等级
var _g_LogLevel int = LOGLEVEL_ALL
var (
    // Unified log date format
    UnifiedLogDataFormat string = "2006-01-02 15:04:05"
    // unified log format
    UnifiedLogFormat string
    // unified log level
    UnifiedLogLevel int = LOGLEVEL_ALL
)

func init() {
    CommonLoggor = NewCommonLog()
}

func SetLogLevel(level int) int {
    _g_LogLevel = level
    return _g_LogLevel
}
