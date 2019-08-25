package log

// Logger interface
type Logger interface {
	// Print calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Print.
	Print(v ...interface{})

	// Printf calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Printf.
	Printf(format string, v ...interface{})

	// Println calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Println.
	Println(v ...interface{})

	// Fatal is equivalent to Print() followed by a call to os.Exit(1).
	Fatal(v ...interface{})

	// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
	Fatalf(format string, v ...interface{})

	// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
	Fatalln(v ...interface{})

	// Panic is equivalent to Print() followed by a call to panic().
	Panic(v ...interface{})

	// Panicf is equivalent to Printf() followed by a call to panic().
	Panicf(format string, v ...interface{})

	// Panicln is equivalent to Println() followed by a call to panic().
	Panicln(v ...interface{})

	// Debug level log
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	// Info level log
	Info(v ...interface{})
	Infof(format string, v ...interface{})

	// Warning level log
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	// Error level log
	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	WithFields(map[string]interface{}) Logger
}
