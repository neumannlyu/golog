package log

import (
    "fmt"
    "runtime"

    "github.com/fatih/color"
)

// common log object
var CommonLoggor CommonLog

type CommonLog struct {
    Level       int
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

// 设置日志等级
func (log *CommonLog) SetLogLevel(level int) {
    log.Level = level
}

// func (l CommonLog) UpdateElement(newelement ILogElement) {
//     switch newelement.(type) {
//     case LogDate:
//         for i, e := range l.Element {
//             switch e.(type) {
//             case LogDate:
//                 l.Element[i] = newelement
//             }
//         }
//     case TAG:
//         for i, e := range l.Element {
//             switch e.(type) {
//             case TAG:
//                 l.Element[i] = newelement
//             }
//         }
//     }
// }

// 运行时错误。当err不为空时，会提示错误信息，并且会打印函数的调用堆栈。
// 可以将上层的err直接传到这里，这里处理err。
// demo：
// _, err := function()
// if CatchError(err) {
// ......处理错误的代码
// }
// @param err
// @return bool。如果返回true，表明有错误；false没有发生错误。
func CatchError(err error) bool {
    // 没有发生错误。
    if err == nil {
        return false
    }

    // 发生错误。提示错误信息以及调用堆栈情况。

    log := NewCommonLog()
    // 该条日志对象必须显示
    log.Level = LOGLEVEL_ALL

    log.ErrorTag.Font = color.Bold
    log.ErrorTag.Tag = "CatchError"

    // print err
    log.Error(color.RedString(err.Error()))
    for i := 1; i < 1999; i++ {
        pc, file, line, ok := runtime.Caller(i)
        if !ok || file == "" || pc == 0 {
            break
        }
        // 格式：\t源码文件名:行数 函数名
        log.Error(color.RedString("%s:%d %s", file, line, runtime.FuncForPC(pc).Name()))
    }
    return true
}

//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
// 打印fatal级别的日志
func (commlog CommonLog) Fatal(msg ...string) {
    if commlog.Level > LOGLEVEL_FATAL {
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
    if commlog.Level > LOGLEVEL_ERROR {
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
    if commlog.Level > LOGLEVEL_WARN {
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
    if commlog.Level > LOGLEVEL_INFO {
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
    if commlog.Level > LOGLEVEL_DEBUG {
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
    if commlog.Level > LOGLEVEL_TRACE {
        return
    }
    fmt.Print(commlog.DateElement.ToColorString() + " " +
        commlog.ErrorTag.ToColorString() + " ")
    c := commlog.MsgAttr.GenColor()
    for _, m := range msg {
        c.Print(m)
    }
}
