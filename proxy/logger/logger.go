package logger

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//Init --> Sets log level
func Init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	Warn("Setting log level:", zerolog.DebugLevel)
}

//Fatal --> Info logs
func Fatal(message string, args ...interface{}) {
	toLog := fmt.Sprintf(message, args...)
	log.Fatal().Msg(toLog)
}

//Error --> Info logs
func Error(message string, args ...interface{}) {
	toLog := fmt.Sprintf(message, args...)
	log.Error().Msg(toLog)
}

//Warn --> Info logs
func Warn(message string, args ...interface{}) {
	toLog := fmt.Sprintf(message, args...)
	log.Warn().Msg(toLog)
}

//Info --> Info logs
func Info(message string, args ...interface{}) {
	toLog := fmt.Sprintf(message, args...)
	log.Info().Msg(toLog)
}

//Debug --> Info logs
func Debug(message string, args ...interface{}) {
	toLog := fmt.Sprintf(message, args...)
	log.Debug().Msg(toLog)
}
