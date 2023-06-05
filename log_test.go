package golog

import (
	"errors"
	"testing"

	"github.com/fatih/color"
)

func TestLog(t *testing.T) {
    log := NewDefaultLog()
    log.Level = 7
    var mytime LogTime
    mytime.Fgcolor = color.FgRed
    mytime.Format = "[[2006-01-02 15:04:05]]"
    log.UpdateElement(mytime)

    var tag LogLevel
    tag.Bgcolor = color.BgBlack
    tag.Tag = "MyINFO"
    log.UpdateElement(tag)
    log.Logln(color.CyanString("msg"))
    log.Fatal("fatal\n")
    log.Error("error\n")
    log.Warn("warn\n")
    log.Info("info\n")
    log.Debug("debug\n")
    log.Trace("trace\n")

    CheckError(errors.New("Test Error."))
}
