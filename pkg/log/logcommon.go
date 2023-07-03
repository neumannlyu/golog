package log

import (
    "fmt"

    "github.com/fatih/color"
)

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

// common log object
var CommonLoggor CommonLog

type CommonLog struct {
    DateElement LogDate
    Format      string
    MsgAttr     MsgAttribute
    FatalTag    LogTag
    ErrorTag    LogTag
    InfoTag     LogTag
    WarnTag     LogTag
    DebugTag    LogTag
    TraceTag    LogTag
    // ElementSet  []IElement
}

func NewCommonLog() CommonLog {
    var commonLog CommonLog
    // 默认时间展示方式
    commonLog.DateElement.FormatString = UnifiedLogDataFormat

    // 默认日志排列布局
    // date tag
    commonLog.Format = commonLog.DateElement.Flag() +
        " " + commonLog.TraceTag.Flag() + " "

    // 构造所有的标签
    // fatal
    commonLog.FatalTag.Tag = LOGTAG_FATAL
    commonLog.FatalTag.Fgcolor = color.FgHiRed
    // error
    commonLog.ErrorTag.Tag = LOGTAG_ERROR
    commonLog.ErrorTag.Fgcolor = color.FgRed
    // warn
    commonLog.WarnTag.Tag = LOGTAG_WARN
    commonLog.WarnTag.Fgcolor = color.FgYellow
    // info
    commonLog.InfoTag.Tag = LOGTAG_INFO
    commonLog.InfoTag.Fgcolor = color.FgCyan
    // debug
    commonLog.DebugTag.Tag = LOGTAG_DEBUG
    // trace
    commonLog.TraceTag.Tag = LOGTAG_TRACE
    return commonLog
}

//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
// 打印fatal级别的日志
func (commlog CommonLog) Fatal(msg ...string) {
    if _g_LogLevel > LOGLEVEL_FATAL {
        return
    }
    fmt.Print(commlog.DateElement.ToColorString() + " " +
        commlog.FatalTag.ToColorString() + " ")
    c := commlog.MsgAttr.GenColor()
    for _, m := range msg {
        c.Print(m)
    }
}

// 打印error级别的日志
func (commlog CommonLog) Error(msg ...string) {
    if _g_LogLevel > LOGLEVEL_ERROR {
        return
    }
    fmt.Print(commlog.DateElement.ToColorString() + " " +
        commlog.ErrorTag.ToColorString() + " ")
    c := commlog.MsgAttr.GenColor()
    for _, m := range msg {
        c.Print(m)
    }
}

// 打印error级别的日志
func (commlog CommonLog) Warn(msg ...string) {
    if _g_LogLevel > LOGLEVEL_WARN {
        return
    }
    fmt.Print(commlog.DateElement.ToColorString() + " " +
        commlog.WarnTag.ToColorString() + " ")
    c := commlog.MsgAttr.GenColor()
    for _, m := range msg {
        c.Print(m)
    }
}

// 打印info级别的日志
func (commlog CommonLog) Info(msg ...string) {
    if _g_LogLevel > LOGLEVEL_INFO {
        return
    }
    fmt.Print(commlog.DateElement.ToColorString() + " " +
        commlog.InfoTag.ToColorString() + " ")
    c := commlog.MsgAttr.GenColor()
    for _, m := range msg {
        c.Print(m)
    }
}

// 打印debug级别的日志
func (commlog CommonLog) Debug(msg ...string) {
    if _g_LogLevel > LOGLEVEL_DEBUG {
        return
    }
    fmt.Print(commlog.DateElement.ToColorString() + " " +
        commlog.DebugTag.ToColorString() + " ")
    c := commlog.MsgAttr.GenColor()
    for _, m := range msg {
        c.Print(m)
    }
}

// 打印trace级别的日志
func (commlog CommonLog) Trace(msg ...string) {
    if _g_LogLevel > LOGLEVEL_TRACE {
        return
    }
    fmt.Print(commlog.DateElement.ToColorString() + " " +
        commlog.TraceTag.ToColorString() + " ")
    c := commlog.MsgAttr.GenColor()
    for _, m := range msg {
        c.Print(m)
    }
}
