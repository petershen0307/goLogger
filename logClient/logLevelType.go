package logClient

// LogLevel is for log priority level
type LogLevel int

// copy from syslog Priority
const (
	LevelEmerg LogLevel = iota
	LevelAlert
	LevelCrit
	LevelErr
	LevelWarning
	LevelNotice
	LevelInfo
	LevelDebug
)

func (l *LogLevel) toString() string {
	var val string
	switch *l {
	case LevelEmerg:
		val = "[Emerg]"
	case LevelAlert:
		val = "[Alert]"
	case LevelCrit:
		val = "[Crit]"
	case LevelErr:
		val = "[Err]"
	case LevelWarning:
		val = "[Warning]"
	case LevelNotice:
		val = "[Notice]"
	case LevelInfo:
		val = "[Info]"
	case LevelDebug:
		val = "[Debug]"
	}
	return val
}
