package golog

// 7:slient/off：完全不输出信息
// 6:fatal：导致程序退出，输出程序退出时的错误信息
// 5:error：错误信息
// 4:warn：警告信息
// 3:info：用于输出信息
// 2:debug：输出调试信息
// 1:trace： 程序运行时的跟踪信息
// 0:all： 所有信息
var (
    LOGLEVEL_SLIENT int    = 7
    LOGLEVEL_FATAL  int    = 6
    LOGLEVEL_ERROR  int    = 5
    LOGLEVEL_WARN   int    = 4
    LOGLEVEL_INFO   int    = 3
    LOGLEVEL_DEBUG  int    = 2
    LOGLEVEL_TRACE  int    = 1
    LOGLEVEL_ALL    int    = 0
    LOGTAG_FATAL    string = "FATAL"
    LOGTAG_ERROR    string = "ERROR"
    LOGTAG_WARN     string = "WARN"
    LOGTAG_INFO     string = "INFO"
    LOGTAG_DEBUG    string = "DEBUG"
    LOGTAG_TRACE    string = "TRACE"
)

// 全局变量：日志等级
var _g_LogLevel int = LOGLEVEL_ALL

// common log object
var _g_CommonLoggor CommonLog
var _ElementPrefix string = "&(E*#@"
var _ElementSuffix string = "#*)"
var (
    // Unified log date format
    UnifiedLogDataFormat string = "2006-01-02 15:04:05"
    // unified log format
    UnifiedLogFormat string
    // unified log level
    UnifiedLogLevel int = LOGLEVEL_ALL
)

func init() {
    _g_CommonLoggor = NewCommonLog()
}
