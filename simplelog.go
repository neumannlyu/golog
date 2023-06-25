package golog

import (
    "fmt"
    "strings"

    "github.com/fatih/color"
)

type SimpleLog struct {
    // 日志等级
    Level int
    // 日志中的日期元素
    Data LogDate
    // 日志中的标签原色
    Tag TAG
    // 日志的格式化字符串
    FormatString string
    // 日志消息
    Msg MsgAttribute
}

// 新建一个simplelog
func NewSimpleLog() SimpleLog {
    var simple SimpleLog
    // 默认时间展示方式
    simple.Data.Format = "2006-01-02 15:04:05"
    simple.Data.Font = color.Bold
    simple.Data.Fgcolor = color.FgRed

    // 默认日志等级展示方式
    simple.Tag.Font = color.Bold
    simple.Tag.Bgcolor = color.BgBlue
    simple.Tag.Tag = "log"

    // 新建默认的日志对象
    simple.FormatString = "[&DT] [&TAG] "

    return simple
}

// 设置日志等级
func (simple *SimpleLog) SetLogLevel(level int) {
    simple.Level = level
}

// 更新元素
func (simple *SimpleLog) UpdateElement(newelement ILogElement) {
    switch element := newelement.(type) {
    case LogDate:
        simple.Data = element
    case TAG:
        simple.Tag = element
    }
}

// print a log
func (simplelog SimpleLog) Log(strs ...string) {
    format := simplelog.FormatString

    // 格式化Data
    format = strings.ReplaceAll(format, simplelog.Data.Flag(), simplelog.Data.ToString())

    // 格式化TAG
    format = strings.ReplaceAll(format, simplelog.Tag.Flag(), simplelog.Tag.ToString())

    // print
    fmt.Print(format)

    c := color.New()
    if simplelog.Msg.Bgcolor > 0 {
        c = c.Add(simplelog.Msg.Bgcolor)
    }
    if simplelog.Msg.Fgcolor > 0 {
        c = c.Add(simplelog.Msg.Fgcolor)
    }
    if simplelog.Msg.Font > 0 {
        c.Add(simplelog.Msg.Font)
    }
    for _, v := range strs {
        c.Print(v)
    }
    fmt.Println()
}
