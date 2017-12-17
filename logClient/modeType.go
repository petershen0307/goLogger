package logClient

type logDisplayMode int

// ModeSetting to configure how to display log string
var ModeSetting logDisplayMode = ModePipe

// mode setting value
const (
	ModePipe logDisplayMode = iota
	ModePrint
)
