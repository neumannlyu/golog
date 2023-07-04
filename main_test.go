package golog

import (
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
    // 1. 设置日志等级
    SetLogLevel(LOGLEVEL_ALL)
    // Test simple log
    simplelog := NewSimpleLog()
    simplelog.Logln("这是一个测试的日志。")
    Trace("hello ", "pop\n")
    GetCommonLog().TraceTag.FormatString = "NewTag"
    Trace("hello\n")
    Fatal("fatal\n")
    Error("error\n")
    Debug("debug\n")
    Info("info\n")
    Trace("trace\n")
    CatchError(errors.New("newerror"))
    NewCommonLog().Error("new err\n")
}
