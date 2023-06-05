package golog

import (
    "fmt"
    "runtime"
    "strings"

    "github.com/fatih/color"
)

type StLog struct {
    Level    int
    Format   string
    FatalTag LogLevel
    ErrorTag LogLevel
    InfoTag  LogLevel
    WarnTag  LogLevel
    DebugTag LogLevel
    TraceTag LogLevel
    Element  []ILogElement
}

func NewDefaultLog() StLog {
    // 默认时间展示方式
    var time LogTime
    time.Format = "2006-01-02 15:04:05"
    time.Font = color.Bold
    time.Fgcolor = color.FgRed

    // 默认日志等级展示方式
    var leveltag LogLevel
    leveltag.Font = color.Bold
    leveltag.Bgcolor = color.BgBlue
    leveltag.Tag = "INFO"

    // 新建默认的日志对象
    var log StLog
    log.Format = "[&DT] [&TAG]"
    log.Element = append(log.Element, time, leveltag)

    // 构造所有的标签
    // fatal
    log.FatalTag.Tag = LOGTAG_FATAL
    log.FatalTag.Fgcolor = color.FgHiRed
    // error
    log.ErrorTag.Tag = LOGTAG_ERROR
    log.ErrorTag.Fgcolor = color.FgHiRed
    // warn
    log.WarnTag.Tag = LOGTAG_WARN
    log.WarnTag.Fgcolor = color.FgYellow
    // info
    log.InfoTag.Tag = LOGTAG_INFO
    log.InfoTag.Fgcolor = color.FgCyan
    // debug
    log.DebugTag.Tag = LOGTAG_DEBUG
    // trace
    log.TraceTag.Tag = LOGTAG_TRACE
    return log
}

func (l StLog) Println(strs ...string) {
    log := l.Format
    for _, e := range l.Element {
        log = strings.ReplaceAll(log, e.Flag(), e.ToString())
    }
    fmt.Print(log)
    for _, v := range strs {
        fmt.Print(v)
    }
    fmt.Println()
}

func (l StLog) UpdateElement(newelement ILogElement) {
    switch newelement.(type) {
    case LogTime:
        for i, e := range l.Element {
            switch e.(type) {
            case LogTime:
                l.Element[i] = newelement
            }
        }
    case LogLevel:
        for i, e := range l.Element {
            switch e.(type) {
            case LogLevel:
                l.Element[i] = newelement
            }
        }
    }
}

// 运行时错误。当err不为空时，会提示错误信息，并且会打印函数的调用堆栈。
// 可以将上层的err直接传到这里，这里处理err。
// demo：
// _, err := function()
// if RunTimeError(err) {
// ......处理错误的代码
// }
// @param err
// @return bool。如果返回true，表明有错误；false没有发生错误。
func CheckError(err error) bool {
    // 没有发生错误。
    if err == nil {
        return false
    }

    // 发生错误。提示错误信息以及调用堆栈情况。
    var logtime LogTime
    logtime.Format = "[2006-01-02 15:04:05]"
    logtime.Fgcolor = color.FgRed
    var errtag LogLevel
    errtag.Tag = LOGTAG_ERROR
    errtag.Bgcolor = color.BgRed
    errtag.Fgcolor = color.FgHiRed

    log := NewDefaultLog()
    log.Level = 7
    log.Format = logtime.Flag() + " " + errtag.Flag() + " "
    log.UpdateElement(logtime)
    log.UpdateElement(errtag)

    log.Println(color.RedString(err.Error()))
    for i := 1; i < 1999; i++ {
        pc, file, line, ok := runtime.Caller(i)
        if !ok || file == "" || pc == 0 {
            break
        }
        // 格式：\t源码文件名:行数 函数名
        log.Println(color.RedString("%s:%d %s", file, line, runtime.FuncForPC(pc).Name()))
    }
    return true
}

//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
// 打印fatal级别的日志
func (l StLog) Fatal(msg ...string) {
    if l.Level < LOGLEVEL_FATAL {
        return
    }
    // 用新的副本操作，防止修改原本设置的tag格式
    tmp := l
    // 切片需要新建对象，否则仍会修改原本对象的值
    tmp.Element = make([]ILogElement, len(l.Element))
    copy(tmp.Element, l.Element)
    tmp.UpdateElement(tmp.FatalTag)

    // 格式化字符串
    log := l.Format
    for _, e := range tmp.Element {
        log = strings.ReplaceAll(log, e.Flag(), e.ToString())
    }

    // 输出
    fmt.Print(log)
    for _, v := range msg {
        fmt.Print(v)
    }
}

// 打印error级别的日志
func (l StLog) Error(msg ...string) {
    if l.Level < LOGLEVEL_ERROR {
        return
    }
    // 用新的副本操作，防止修改原本设置的tag格式
    tmp := l
    // 切片需要新建对象，否则仍会修改原本对象的值
    tmp.Element = make([]ILogElement, len(l.Element))
    copy(tmp.Element, l.Element)
    tmp.UpdateElement(tmp.ErrorTag)

    // 格式化字符串
    log := l.Format
    for _, e := range tmp.Element {
        log = strings.ReplaceAll(log, e.Flag(), e.ToString())
    }

    // 输出
    fmt.Print(log)
    for _, v := range msg {
        fmt.Print(v)
    }
}

// 打印error级别的日志
func (l StLog) Warn(msg ...string) {
    if l.Level < LOGLEVEL_WARN {
        return
    }
    // 用新的副本操作，防止修改原本设置的tag格式
    tmp := l
    // 切片需要新建对象，否则仍会修改原本对象的值
    tmp.Element = make([]ILogElement, len(l.Element))
    copy(tmp.Element, l.Element)
    tmp.UpdateElement(tmp.WarnTag)

    // 格式化字符串
    log := l.Format
    for _, e := range tmp.Element {
        log = strings.ReplaceAll(log, e.Flag(), e.ToString())
    }

    // 输出
    fmt.Print(log)
    for _, v := range msg {
        fmt.Print(v)
    }
}

// 打印info级别的日志
func (l StLog) Info(msg ...string) {
    if l.Level < LOGLEVEL_INFO {
        return
    }
    // 用新的副本操作，防止修改原本设置的tag格式
    tmp := l
    // 切片需要新建对象，否则仍会修改原本对象的值
    tmp.Element = make([]ILogElement, len(l.Element))
    copy(tmp.Element, l.Element)
    tmp.UpdateElement(tmp.InfoTag)

    // 格式化字符串
    log := l.Format
    for _, e := range tmp.Element {
        log = strings.ReplaceAll(log, e.Flag(), e.ToString())
    }

    // 输出
    fmt.Print(log)
    for _, v := range msg {
        fmt.Print(v)
    }
}

// 打印debug级别的日志
func (l StLog) Debug(msg ...string) {
    if l.Level < LOGLEVEL_DEBUG {
        return
    }
    // 用新的副本操作，防止修改原本设置的tag格式
    tmp := l
    // 切片需要新建对象，否则仍会修改原本对象的值
    tmp.Element = make([]ILogElement, len(l.Element))
    copy(tmp.Element, l.Element)
    tmp.UpdateElement(tmp.DebugTag)

    // 格式化字符串
    log := l.Format
    for _, e := range tmp.Element {
        log = strings.ReplaceAll(log, e.Flag(), e.ToString())
    }

    // 输出
    fmt.Print(log)
    for _, v := range msg {
        fmt.Print(v)
    }
}

// 打印trace级别的日志
func (l StLog) Trace(msg ...string) {
    if l.Level < LOGLEVEL_TRACE {
        return
    }
    // 用新的副本操作，防止修改原本设置的tag格式
    tmp := l
    // 切片需要新建对象，否则仍会修改原本对象的值
    tmp.Element = make([]ILogElement, len(l.Element))
    copy(tmp.Element, l.Element)
    tmp.UpdateElement(tmp.TraceTag)

    // 格式化字符串
    log := l.Format
    for _, e := range tmp.Element {
        log = strings.ReplaceAll(log, e.Flag(), e.ToString())
    }

    // 输出
    fmt.Print(log)
    for _, v := range msg {
        fmt.Print(v)
    }
}
