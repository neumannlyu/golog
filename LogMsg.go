package golog

import "github.com/fatih/color"

type MsgAttribute struct {
    Font    color.Attribute
    Bgcolor color.Attribute
    Fgcolor color.Attribute
}

func (m MsgAttribute) GenColor() color.Color {
    return *color.New(m.Font, m.Bgcolor, m.Fgcolor)
}
