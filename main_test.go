package golog

import (
    "testing"
)

func TestLog(t *testing.T) {
    // 1. 设置日志等级
    SetLogLevel(LOGLEVEL_ALL)
    // Test simple log
    simplelog := NewSimpleLog()
    simplelog.Logln("这是一个测试的日志。")
    Trace("hello ", "pop\n")
    GetCommonLog().TraceTag.Tag = "New Tag"
    Trace("水电费水电费")
}
