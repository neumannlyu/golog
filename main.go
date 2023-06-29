package golog

import (
    "runtime"

    "github.com/fatih/color"
    "github.com/neumannlyu/golog/pkg/log"
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
    // 该条日志对象必须显示
    tmplog.Level = log.LOGLEVEL_ALL

    tmplog.ErrorTag.Font = color.Bold
    tmplog.ErrorTag.Tag = "CatchError"

    // print err
    tmplog.Error(color.RedString(err.Error()))
    for i := 1; i < 1999; i++ {
        pc, file, line, ok := runtime.Caller(i)
        if !ok || file == "" || pc == 0 {
            break
        }
        // 格式：\t源码文件名:行数 函数名
        tmplog.Error(color.RedString("%s:%d %s",
            file, line, runtime.FuncForPC(pc).Name()))
    }
    return true
}
