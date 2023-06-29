package log

// var _ElementPrefix string = "&(E*#@"
// var _ElementSuffix string = "#*)"
var (
    // Unified log date format
    UnifiedLogDataFormat string = "2006-01-02 15:04:05"
    // unified log format
    UnifiedLogFormat string
    // unified log level
    UnifiedLogLevel int = LOGLEVEL_ALL
)

// type ILogElement interface {
//     ToString() string
//     Flag() string
// }

func init() {
    CommonLoggor = NewCommonLog()
}
