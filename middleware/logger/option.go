package logger

// Option function option
type Option func(*logConfig)

// Path set logPath
// if is zero will print,or write file
func Path(logPath string) Option {
	return func(_logCfg *logConfig) {
		_logCfg.LogPath = logPath
	}
}

// Compress compress log
func Compress(compress bool) Option {
	return func(_logCfg *logConfig) {
		_logCfg.Compress = compress
	}
}

// Level set log level default info
func Level(level string) Option {
	return func(_logCfg *logConfig) {
		_logCfg.LogLevel = level
	}
}

// MaxSize Log Max Size
func MaxSize(size int) Option {
	return func(_logCfg *logConfig) {
		_logCfg.MaxSize = size
	}
}

// MaxAge log store day
func MaxAge(age int) Option {
	return func(_logCfg *logConfig) {
		_logCfg.MaxAge = age
	}
}

// MaxBackups total store log
func MaxBackups(backup int) Option {
	return func(_logCfg *logConfig) {
		_logCfg.MaxBackups = backup
	}
}
