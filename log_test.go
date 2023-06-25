package golog

import (
    "testing"

    "github.com/fatih/color"
)

func TestLog(t *testing.T) {
    // Test simple log
    simplelog := NewSimpleLog()
    simplelog.Level = 7
    simplelog.Msg.Bgcolor = color.BgCyan
    simplelog.Msg.Fgcolor = color.FgYellow
    simplelog.Tag.Font = color.Bold
    simplelog.Tag.Tag = "MyTag"
    simplelog.Log("这是一个测试的日志。")
    // log := NewDefaultLog()
    // log.Level = 7
    // var mytime LogDate
    // mytime.Fgcolor = color.FgRed
    // mytime.Format = "[[2006-01-02 15:04:05]]"
    // log.UpdateElement(mytime)

    // var tag TAG
    // tag.Bgcolor = color.BgBlack
    // tag.Tag = "MyINFO"
    // log.UpdateElement(tag)
    // log.Logln(color.CyanString("msg"))
    // log.Fatal("fatal\n")
    // log.Error("error\n")
    // log.Warn("warn\n")
    // log.Info("info\n")
    // log.Debug("debug\n")
    // log.Trace("trace\n")

    // CheckError(errors.New("Test Error."))
}
