package utils

import (
	"bufio"
	"os"

	"github.com/betas-in/logger"
)

type CLIFunctions struct{}

func CLI() CLIFunctions {
	return CLIFunctions{}
}

func (c CLIFunctions) Question(clog *logger.CLILogger, where, format string, v ...interface{}) string {
	reader := bufio.NewReader(os.Stdin)

	clog.Announce(where).Msgf(format, v...)
	text, err := reader.ReadString('\n')
	if err != nil {
		clog.Error(where).Msgf("could not read input: %v", err)
	}
	return text
}
