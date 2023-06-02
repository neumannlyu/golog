package golog

import (
    "fmt"
    "strings"

    "github.com/fatih/color"
)

type StLog struct {
    Level    int
    Format   string
    FatalTag StLevelTag
    ErrorTag StLevelTag
    InfoTag  StLevelTag
    WarnTag  StLevelTag
    DebugTag StLevelTag
    TraceTag StLevelTag
    Element  []ILogElement
}

func NewDefaultLog() StLog {
    // 默认时间展示方式
    var time StTime
    time.Format = "2006-01-02 15:04:05"
    time.Font = color.Bold
    time.Fgcolor = color.FgRed

    // 默认日志等级展示方式
    var leveltag StLevelTag
    leveltag.Font = color.Bold
    leveltag.Bgcolor = color.BgBlue
    leveltag.tag = "INFO"

    // 新建默认的日志对象
    var log StLog
    log.Format = "[&DT] [&TAG]"
    log.Element = append(log.Element, time, leveltag)

    // 构造所有的标签
    // fatal
    log.FatalTag.tag = LOGTAG_FATAL
    log.FatalTag.Fgcolor = color.FgHiRed
    // error
    log.ErrorTag.tag = LOGTAG_ERROR
    log.ErrorTag.Fgcolor = color.FgHiRed
    // warn
    log.WarnTag.tag = LOGTAG_WARN
    log.WarnTag.Fgcolor = color.FgYellow
    // info
    log.InfoTag.tag = LOGTAG_INFO
    log.InfoTag.Fgcolor = color.FgCyan
    // debug
    log.DebugTag.tag = LOGTAG_DEBUG
    // trace
    log.TraceTag.tag = LOGTAG_TRACE
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
    case StTime:
        for i, e := range l.Element {
            switch e.(type) {
            case StTime:
                l.Element[i] = newelement
            }
        }
    case StLevelTag:
        for i, e := range l.Element {
            switch e.(type) {
            case StLevelTag:
                l.Element[i] = newelement
            }
        }
    }
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
