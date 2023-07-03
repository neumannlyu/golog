package golog

import (
    "fmt"
    "runtime"

    "github.com/fatih/color"
    "github.com/neumannlyu/golog/pkg/log"
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

// 新建一个simplelog
func NewSimpleLog() log.SimpleLog {
    var simple log.SimpleLog
    // 默认时间展示方式
    simple.Data.FormatString = log.UnifiedLogDataFormat
    simple.Tag.Tag = "simplelog"
    // 新建默认的日志对象
    simple.FormatString = simple.Data.Flag() + " " + simple.Tag.Flag() + " "
    return simple
}

// 获取通用日志对象,通常都使用这个对象进行自定义配置
func GetCommonLog() *log.CommonLog {
    return &log.CommonLoggor
}

// 设置日志显示等级。默认等级为7（ALL）
func SetLogLevel(level int) int {
    return log.SetLogLevel(level)
}

// 打印fatal级别的日志
func Fatal(msg ...string) {
    log.CommonLoggor.Fatal(msg...)
}

// 打印error级别的日志
func Error(msg ...string) {
    log.CommonLoggor.Error(msg...)
}

// 打印error级别的日志
func Warn(msg ...string) {
    log.CommonLoggor.Warn(msg...)
}

// 打印info级别的日志
func Info(msg ...string) {
    log.CommonLoggor.Info(msg...)
}

// 打印debug级别的日志
func Debug(msg ...string) {
    log.CommonLoggor.Debug(msg...)
}

// 打印trace级别的日志
func Trace(msg ...string) {
    log.CommonLoggor.Trace(msg...)
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

    tmplog := log.NewCommonLog()
    tmplog.ErrorTag.Font = color.Bold
    tmplog.ErrorTag.Tag = "CatchError"

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
