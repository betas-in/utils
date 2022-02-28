package utils

import (
	"time"

	"github.com/betas-in/logger"
)

type TimeFunctions struct{}

func Time() TimeFunctions {
	return TimeFunctions{}
}

func Elapsed(log *logger.Logger, what string, logging bool) func() {
	start := time.Now()
	return func() {
		if logging {
			log.Debug(what).Msgf("%s\n", time.Since(start))
		}
	}
}
