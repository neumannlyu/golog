package golog

import (
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
