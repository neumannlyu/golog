package log

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
    Tag LogTag
    // 日志的格式化字符串
    FormatString string
    // 日志消息
    Msg MsgAttribute
}

// 设置日志等级
func (simple *SimpleLog) SetLogLevel(level int) {
    simple.Level = level
}

// 更新元素
func (simple *SimpleLog) UpdateElement(newelement IElement) {
    switch element := newelement.(type) {
    case LogDate:
        simple.Data = element
    case LogTag:
        simple.Tag = element
    }
}

// print a log
func (simplelog SimpleLog) Log(strs ...string) {
    format := simplelog.FormatString

    // 格式化Data
    format = strings.ReplaceAll(
        format, simplelog.Data.Flag(), simplelog.Data.ToColorString())

    // 格式化TAG
    format = strings.ReplaceAll(
        format, simplelog.Tag.Flag(), simplelog.Tag.ToColorString())

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
}

func (simplelog SimpleLog) Logln(strs ...string) {
    format := simplelog.FormatString
    fmt.Printf("format: %v\n", format)
    // 格式化Data
    format = strings.ReplaceAll(
        format, simplelog.Data.Flag(), simplelog.Data.ToColorString())

    // 格式化TAG
    format = strings.ReplaceAll(
        format, simplelog.Tag.Flag(), simplelog.Tag.ToColorString())

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
