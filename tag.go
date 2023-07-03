package golog

import "github.com/fatih/color"

type LogTag struct {
    Tag     string
    Font    color.Attribute
    Bgcolor color.Attribute
    Fgcolor color.Attribute
}

func (tag LogTag) Flag() string {
    return _ElementPrefix + "TAG" + _ElementSuffix
}

func (tag LogTag) ToColorString() string {
    c := color.New()
    if tag.Font >= 0 {
        c.Add(tag.Font)
    }
    if tag.Bgcolor >= 0 {
        c.Add(tag.Bgcolor)
    }
    if tag.Fgcolor >= 0 {
        c.Add(tag.Fgcolor)
    }
    return c.Sprint(tag.Tag)
}
