package log

import (
    "time"

    "github.com/fatih/color"
)

type LogDate struct {
    id           int
    FormatString string
    Font         color.Attribute
    Bgcolor      color.Attribute
    Fgcolor      color.Attribute
}

func (t LogDate) Flag() string {
    return _ElementPrefix + "Date" + _ElementSuffix
}

func (t LogDate) ToColorString() string {
    c := color.New()
    if t.Font >= 0 {
        c.Add(t.Font)
    }
    if t.Bgcolor >= 0 {
        c.Add(t.Bgcolor)
    }
    if t.Fgcolor >= 0 {
        c.Add(t.Fgcolor)
    }
    return c.Sprint(time.Now().Format(t.FormatString))
}
