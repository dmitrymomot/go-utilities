package log

// M is a short way to add map[string]interface{}
type M map[string]interface{}

var defaultLogger Logger

func init() {
	if defaultLogger == nil {
		SetDefault(NewConsoleLogger(0))
	}
}

// SetDefault logger
func SetDefault(l Logger) {
	defaultLogger = l
}

// StandardLogger returns current logger instance
func StandardLogger() Logger {
	return defaultLogger
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	StandardLogger().Print(v...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	StandardLogger().Printf(format, v...)
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	StandardLogger().Println(v...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	StandardLogger().Fatal(v...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	StandardLogger().Fatalf(format, v...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	StandardLogger().Fatalln(v...)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	StandardLogger().Panic(v...)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	StandardLogger().Panicf(format, v...)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	StandardLogger().Panicln(v...)
}

// Debug level log
func Debug(v ...interface{}) {
	StandardLogger().Debug(v...)
}

// Debugf level log with arguments formating
func Debugf(format string, v ...interface{}) {
	StandardLogger().Debugf(format, v...)
}

// Info level log
func Info(v ...interface{}) {
	StandardLogger().Info(v...)
}

// Infof level log with arguments formating
func Infof(format string, v ...interface{}) {
	StandardLogger().Infof(format, v...)
}

// Warn level log
func Warn(v ...interface{}) {
	StandardLogger().Warn(v...)
}

// Warnf level log with arguments formating
func Warnf(format string, v ...interface{}) {
	StandardLogger().Warnf(format, v...)
}

// Error level log
func Error(v ...interface{}) {
	StandardLogger().Error(v...)
}

// Errorf level log with arguments formating
func Errorf(format string, v ...interface{}) {
	StandardLogger().Errorf(format, v...)
}

// WithFields adds attributes to log
func WithFields(fields map[string]interface{}) Logger {
	return StandardLogger().WithFields(fields)
}
