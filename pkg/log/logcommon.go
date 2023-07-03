package log

import (
	"fmt"

	"github.com/fatih/color"
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
