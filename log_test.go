package golog

import (
    "testing"

    "github.com/fatih/color"
)

func TestLog(t *testing.T) {
    log := NewDefaultLog()
    log.Level = 7
    var mytime StTime
    mytime.Fgcolor = color.FgRed
    mytime.Format = "[[2006-01-02 15:04:05]]"
    log.UpdateElement(mytime)

    var tag StLevelTag
    tag.Bgcolor = color.BgBlack
    tag.tag = "INFOF"
    log.UpdateElement(tag)
    log.Println(color.CyanString("  dsfdg"))
    log.Fatal("kkkkk\n")
    log.Println(color.CyanString("xxxdsfdg"))
    log.Error("error\n")
    log.Println(color.CyanString("xxxdsfdg"))
    log.Warn("warn\n")
    log.Info("info\n")
    log.Debug("debug\n")
    log.Trace("trace\n")
    log.Println(color.CyanString("xxxdsfdg"))
}
