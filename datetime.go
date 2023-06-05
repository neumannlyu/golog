package golog

import (
    "time"

    "github.com/fatih/color"
)

type LogTime struct {
    Format  string
    Font    color.Attribute
    Bgcolor color.Attribute
    Fgcolor color.Attribute
}

func (t LogTime) Flag() string {
    return "&DT"
}

func (t LogTime) ToString() string {
    c := color.New()
    if t.Font > 0 {
        c.Add(t.Font)
    }
    if t.Bgcolor > 0 {
        c.Add(t.Bgcolor)
    }
    if t.Fgcolor > 0 {
        c.Add(t.Fgcolor)
    }
    return c.Sprint(time.Now().Format(t.Format))
}
