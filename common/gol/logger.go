package gol

// Info logger information message
func (l *Gol) Info(value ...interface{}) {
	l.logInfo.Println(value...)
}

// Infof logger information message with format mode
func (l *Gol) Infof(format string, a ...any) {
	l.logInfo.Printf(format, a)
}

// Warn logger light error(warn) info
func (l *Gol) Warn(value ...interface{}) {
	l.logWarning.Println(value...)
}

// Warnf logger light error(warn) info with format mode
func (l *Gol) Warnf(format string, a ...any) {
	l.logInfo.Printf(format, a)
}

// Err logger critical error info
func (l *Gol) Err(value ...interface{}) {
	l.logError.Println(value...)
}

// Errf logger critical error info with format mode
func (l *Gol) Errf(format string, a ...any) {
	l.logInfo.Printf(format, a)
}
