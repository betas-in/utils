package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestZip(t *testing.T) {
	err := Zip().Gunzip("one", "two")
	Test().Equals(t, "open one: no such file or directory", err.Error())

	err = Zip().Gunzip("gunzip.go", "gunzippedFile")
	Test().Equals(t, "gzip: invalid header", err.Error())

	err = Zip().Gunzip(filepath.Join(".", "testdata", "testFile.txt.gz"), "gunzippedFile")
	Test().Nil(t, err)

	defer os.Remove("gunzippedFile")

	err = Zip().UntarGunzip("one", "two")
	Test().Equals(t, "open one: no such file or directory", err.Error())

	err = Zip().UntarGunzip("gunzip.go", "unTarredGunzippedFile")
	Test().Equals(t, "gzip: invalid header", err.Error())
	defer os.Remove("unTarredGunzippedFile")

	err = Zip().UntarGunzip(filepath.Join(".", "testdata", "testFolder.tar.gz"), "three")
	Test().Nil(t, err)
	defer os.RemoveAll("three")

	err = Zip().UntarGunzip(filepath.Join(".", "testdata", "account.tar.gz"), "four")
	Test().Nil(t, err)
	defer os.RemoveAll("four")
}
