package golog

type IElement interface {
    // get flag of elememt
    Flag() string
    // print
    ToColorString() string
}
