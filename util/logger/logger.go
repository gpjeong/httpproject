package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"time"
)

// LoggerServer 로깅 관련 정보
type LoggerServer struct {
	Level LogLevel
	// 콘솔 로깅 활성화
	ConsoleLoggingEnabled bool
	// EncodeLogsAsJson 은 로그 프레임워크 로그 JSON을 만듭니다.
	EncodeLogsAsJson bool
	// FileLoggingEnabled 는 프레임워크가 파일에 로그하도록 합니다.
	// 이 값이 false이면 아래 필드를 건너뛸 수 있습니다!
	FileLoggingEnabled bool
	// Directory 파일 로깅이 활성화된 경우 로그인할 디렉터리
	Directory string
	// Filename 은 디렉토리 안에 위치할 로그 파일의 이름입니다.
	Filename string
	// MaxSize 롤링되기 전 로그 파일의 최대 크기(MB)
	MaxSize int
	// MaxBackups 보관할 최대 롤링 파일 수
	MaxBackups int
	// MaxAge 로그 파일을 보관할 최대 기간(일)
	MaxAge int
}

const (
	LogLevelDisable   = iota
	LogLevelEmergency //logger Panic
	LogLevelAlert     //logger XXX == Panic
	LogLevelCritical  //logger Fatal
	LogLevelError
	LogLevelWarning
	LogLevelNotice //logger XXX == WARING
	LogLevelInfo
	LogLevelDebug
	LogLevelDefault //logger XXX == DEBUG
	LogLevelTrace
)

type LogLevel int

type Logger struct {
	logger   *zerolog.Logger
	logLevel LogLevel
}

var Log zerolog.Logger

var LogLevelToString = [...]string{
	"DISABLE",
	"EMERGENCY",
	"ALERT",
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
	"DEFAULT",
	"TRACE",
}

func (l LogLevel) Name() string { return LogLevelToString[l] }

func GetLogLevelFromString(name string) LogLevel {
	level := LogLevelDisable
	for i, v := range LogLevelToString {
		if v == name {
			level = i
			break
		}
	}
	return LogLevel(level)
}

// func New(setLevel LogLevel) *Logger {
func New(logConfig LoggerServer) *Logger {
	setLevel := logConfig.Level
	level := zerolog.InfoLevel
	switch setLevel {
	case LogLevelDisable:
		level = zerolog.Disabled
	case LogLevelEmergency:
		fallthrough
	case LogLevelAlert:
		level = zerolog.PanicLevel
	case LogLevelCritical:
		level = zerolog.FatalLevel
	case LogLevelError:
		level = zerolog.ErrorLevel
	case LogLevelWarning:
		fallthrough
	case LogLevelNotice:
		level = zerolog.WarnLevel
	case LogLevelInfo:
		level = zerolog.InfoLevel
	case LogLevelDebug:
		fallthrough
	case LogLevelDefault:
		level = zerolog.DebugLevel
	case LogLevelTrace:
		level = zerolog.TraceLevel
	default:
	}

	//GCP Level
	zerolog.LevelFieldName = "severity"

	zerolog.LevelPanicValue = "EMERGENCY"
	//"ALERT"
	zerolog.LevelFatalValue = "CRITICAL"
	zerolog.LevelErrorValue = "ERROR"
	zerolog.LevelWarnValue = "WARNING"
	//"NOTICE"
	zerolog.LevelInfoValue = "INFO"
	zerolog.LevelDebugValue = "DEBUG"
	//"DEFAULT"

	zerolog.LevelTraceValue = "TRACE"

	zerolog.SetGlobalLevel(level)
	// origin
	//logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	//////// vm 구성으로 인한 로그 파일 생성 로직으로 변경 시작 ////////
	consoleLoggingEnabled := logConfig.ConsoleLoggingEnabled
	fileLoggingEnabled := logConfig.FileLoggingEnabled
	encodeLogsAsJson := logConfig.EncodeLogsAsJson
	directory := logConfig.Directory
	fileName := logConfig.Filename
	maxBackups := logConfig.MaxBackups
	maxSize := logConfig.MaxSize
	maxAge := logConfig.MaxAge

	var writers []io.Writer

	if consoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}
	if fileLoggingEnabled {
		writers = append(writers, newRollingFile(logConfig))
	}
	mw := io.MultiWriter(writers...)

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", fileLoggingEnabled).
		Bool("jsonLogOutput", encodeLogsAsJson).
		Str("logDirectory", directory).
		Str("fileName", fileName).
		Int("maxSizeMB", maxSize).
		Int("maxBackups", maxBackups).
		Int("maxAgeInDays", maxAge).
		Msg("logging configured")

	Log = logger
	//////// vm 구성으로 인한 로그 파일 생성 로직으로 변경 종료 ////////

	return &Logger{
		logger:   &logger,
		logLevel: setLevel,
	}
}

func newRollingFile(logConfig LoggerServer) io.Writer {
	directory := logConfig.Directory
	fileName := logConfig.Filename
	maxBackups := logConfig.MaxBackups
	maxSize := logConfig.MaxSize
	maxAge := logConfig.MaxAge
	if err := os.MkdirAll(directory, 0744); err != nil {
		log.Error().Err(err).Str("path", directory).Msg("can't create log directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   path.Join(directory, fileName),
		MaxBackups: maxBackups, // files
		MaxSize:    maxSize,    // megabytes
		MaxAge:     maxAge,     // days
	}
}

func (l *Logger) GetHandle() *zerolog.Logger {
	return l.logger
}

// Output duplicates the global logger and sets w as its output.
func (l *Logger) Output(w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}

// With creates a child logger with the field added to its context.
func (l *Logger) With() zerolog.Context {
	return l.logger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func (l *Logger) Level(level zerolog.Level) zerolog.Logger {
	return l.logger.Level(level)
}

// Sample returns a logger with the s sampler.
func (l *Logger) Sample(s zerolog.Sampler) zerolog.Logger {
	return l.logger.Sample(s)
}

// Hook returns a logger with the h Hook.
func (l *Logger) Hook(h zerolog.Hook) zerolog.Logger {
	return l.logger.Hook(h)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Err(err error) *zerolog.Event {
	return l.logger.Err(err).Caller()
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Trace() *zerolog.Event {
	return l.logger.Trace().Caller()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug().Caller()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info().Caller()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn().Caller()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error().Caller()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal().Caller()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic().Caller()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) {
	l.logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

// Emergency starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Emergency() *zerolog.Event {
	return l.logger.Panic()
}

// Critical starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Critical() *zerolog.Event {
	return l.logger.Fatal()
}

func (l *Logger) Alert() *zerolog.Event {
	if l.logLevel >= LogLevelAlert {
		return l.logger.Log().Str("severity", "ALERT")
	} else {
		return l.logger.Panic()
	}
}

func (l *Logger) Notice() *zerolog.Event {
	if l.logLevel >= LogLevelNotice {
		return l.logger.Log().Str("severity", "NOTICE")
	} else {
		return l.logger.Info()
	}
}

func (l *Logger) Default() *zerolog.Event {
	if l.logLevel >= LogLevelDefault {
		return l.logger.Log().Str("severity", "DEFAULT")
	} else {
		return l.logger.Trace()
	}
}
