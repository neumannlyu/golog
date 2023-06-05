package golog

import (
    "github.com/fatih/color"
)

type LogLevel struct {
    Tag     string
    Font    color.Attribute
    Bgcolor color.Attribute
    Fgcolor color.Attribute
}

func (lt LogLevel) Flag() string {
    return "&TAG"
}

func (lt LogLevel) ToString() string {
    c := color.New()
    if lt.Font > 0 {
        c.Add(lt.Font)
    }
    if lt.Bgcolor > 0 {
        c.Add(lt.Bgcolor)
    }
    if lt.Fgcolor > 0 {
        c.Add(lt.Fgcolor)
    }
    return c.Sprint(lt.Tag)
}
