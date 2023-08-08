package golog

import (
    "fmt"
)

type CommonLog struct {
    DateElement LogDate
    FatalTag    LogTag
    ErrorTag    LogTag
    InfoTag     LogTag
    WarnTag     LogTag
    DebugTag    LogTag
    TraceTag    LogTag
    Format      string
    MsgAttr     MsgAttribute
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
        fmt.Print(c.Sprint(m))
    }
    fmt.Println()
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
        fmt.Print(c.Sprint(m))
    }
    fmt.Println()
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
        fmt.Print(c.Sprint(m))
    }
    fmt.Println()
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
        fmt.Print(c.Sprint(m))
    }
    fmt.Println()
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
        fmt.Print(c.Sprint(m))
    }
    fmt.Println()
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
        fmt.Print(c.Sprint(m))
    }
    fmt.Println()
}
