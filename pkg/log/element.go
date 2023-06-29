package log

var _ElementPrefix string = "&(E*#@"
var _ElementSuffix string = "#*)"

type IElement interface {
    // get flag of elememt
    Flag() string
    // print
    ToColorString() string
}
