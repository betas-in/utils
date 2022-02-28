package utils

import (
	"testing"
)

func TestTesting(t *testing.T) {
	var err error
	Test().Nil(t, err)

	// err = fmt.Errorf("this is an error")
	// Test().Nil(t, err)

	expected := "this is expected"
	received := "got this"

	Test().Equals(t, expected, expected)
	Test().Equals(t, received, received)
	// Test().Equals(t, expected, received)

	// Test().Contains(t, expected, received)
	Test().Contains(t, expected, "expected")
}

func TestUUID(t *testing.T) {
	u := UUID().Get()
	valid := UUID().IsValid(u)
	Test().Equals(t, true, valid)
}

func TestString(t *testing.T) {
	str := " a-b - c-d - e-f "
	s := String().SplitAndTrim(str, "-")
	Test().Equals(t, 0, Array().Contains(s, "a", true))
	Test().Equals(t, 1, Array().Contains(s, "b", true))
	Test().Equals(t, 2, Array().Contains(s, "c", true))
	Test().Equals(t, 3, Array().Contains(s, "d", true))
	Test().Equals(t, 4, Array().Contains(s, "e", true))
	Test().Equals(t, 5, Array().Contains(s, "f", true))

	s = Array().Delete(s, 0)
	Test().Equals(t, -1, Array().Contains(s, "a", true))

	s = Array().Delete(s, 3)
	Test().Equals(t, 0, Array().Contains(s, "b", true))
	Test().Equals(t, -1, Array().Contains(s, "e", true))

	f := 131.49897
	i := Float64().GetInt64(f, 2)
	Test().Equals(t, int64(13150), i)
}

func TestContentType(t *testing.T) {
	file := "testing.go"
	contentType, err := File().GetContentType(file)
	Test().Nil(t, err)
	Test().Equals(t, "txt", contentType)

	file = "./testdata/test.zip"
	contentType, err = File().GetContentType(file)
	Test().Nil(t, err)
	Test().Equals(t, "zip", contentType)

	file = "./testdata/account.tar.gz"
	contentType, err = File().GetContentType(file)
	Test().Nil(t, err)
	Test().Equals(t, "gz", contentType)
}
