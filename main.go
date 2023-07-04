package golog

import (
    "fmt"
    "runtime"

    "github.com/fatih/color"
)

// 新建一个simplelog
func NewSimpleLog() SimpleLog {
    var simple SimpleLog
    // 默认时间展示方式
    simple.Data = UnifiedLogData
    simple.Tag = UnifiedLogTag
    // 新建默认的日志对象
    simple.FormatString = UnifiedLogFormatString
    return simple
}

func NewCommonLog() CommonLog {
    var commonLog CommonLog
    // 默认时间展示方式
    commonLog.DateElement = UnifiedLogData
    // 默认日志排列布局
    commonLog.Format = UnifiedLogFormatString

    // 构造所有的标签
    // fatal
    commonLog.FatalTag.FormatString = LOGTAG_FATAL
    commonLog.FatalTag.Fgcolor = color.FgHiRed
    // error
    commonLog.ErrorTag.FormatString = LOGTAG_ERROR
    commonLog.ErrorTag.Fgcolor = color.FgRed
    // warn
    commonLog.WarnTag.FormatString = LOGTAG_WARN
    commonLog.WarnTag.Fgcolor = color.FgYellow
    // info
    commonLog.InfoTag.FormatString = LOGTAG_INFO
    commonLog.InfoTag.Fgcolor = color.FgCyan
    // debug
    commonLog.DebugTag.FormatString = LOGTAG_DEBUG
    // trace
    commonLog.TraceTag.FormatString = LOGTAG_TRACE
    return commonLog
}

// 获取通用日志对象,通常都使用这个对象进行自定义配置
func GetCommonLog() *CommonLog {
    return &_g_CommonLoggor
}

// 设置日志显示等级。默认等级为7（ALL）
func SetLogLevel(level int) int {
    _g_LogLevel = level
    return _g_LogLevel
}

// 打印fatal级别的日志
func Fatal(msg ...string) {
    _g_CommonLoggor.Fatal(msg...)
}

// 打印error级别的日志
func Error(msg ...string) {
    _g_CommonLoggor.Error(msg...)
}

// 打印error级别的日志
func Warn(msg ...string) {
    _g_CommonLoggor.Warn(msg...)
}

// 打印info级别的日志
func Info(msg ...string) {
    _g_CommonLoggor.Info(msg...)
}

// 打印debug级别的日志
func Debug(msg ...string) {
    _g_CommonLoggor.Debug(msg...)
}

// 打印trace级别的日志
func Trace(msg ...string) {
    _g_CommonLoggor.Trace(msg...)
}

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

    tmplog := NewCommonLog()
    tmplog.ErrorTag.Font = color.Bold
    tmplog.ErrorTag.FormatString = "CatchError"

    // print err
    tmplog.Error(color.RedString(err.Error()))
    fmt.Println()
    for i := 1; i < 1999; i++ {
        pc, file, line, ok := runtime.Caller(i)
        if !ok || file == "" || pc == 0 {
            break
        }
        // 格式：\t源码文件名:行数 函数名
        tmplog.Error(color.RedString("%s:%d %s",
            file, line, runtime.FuncForPC(pc).Name()))
        fmt.Println()
    }
    return true
}
