package utils

import (
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/betas-in/logger"
	"github.com/davecgh/go-spew/spew"
)

func Test() TestFunctions {
	return TestFunctions{
		logger: logger.NewCLILogger(6, 5),
	}
}

type TestFunctions struct {
	logger *logger.CLILogger
}

// OK fails the test if an err is not nil.
func (t TestFunctions) Nil(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		t.logger.Error("Test").Msgf("\n❌ FAILED at %s:%d", filepath.Base(file), line)
		t.logger.Error("Test").Msgf("Expected: <nil>, got: <%v>\n", err)
		tb.Fail()
	}
}

// Equals fails the test if exp is not equal to act.
func (t TestFunctions) Equals(tb testing.TB, expected, received interface{}) {
	if !reflect.DeepEqual(expected, received) {
		_, file, line, _ := runtime.Caller(1)
		t.logger.Error("Test").Msgf("\n❌ FAILED at %s:%d", filepath.Base(file), line)
		t.logger.Error("Test").Msgf("Expected and Received should have been equal, got:")
		t.logger.Error("Test").Msgf(diff.LineDiff(spew.Sdump(expected), spew.Sdump(received)))
		t.logger.Error("Test").Msgf("")
		tb.Fail()
	}
}

// Contains checks two strings
func (t TestFunctions) Contains(tb testing.TB, fullstring, substring string) {
	if !strings.Contains(fullstring, substring) {
		_, file, line, _ := runtime.Caller(1)
		t.logger.Error("Test").Msgf("\n❌ FAILED at %s:%d", filepath.Base(file), line)
		t.logger.Error("Test").Msgf("The string <%s> does not contain <%s>\n", fullstring, substring)
		tb.Fail()
	}
}
