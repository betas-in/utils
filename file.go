package utils

import (
	"net/http"
	"os"
	"strings"
)

type FileFunctions struct{}

func File() FileFunctions {
	return FileFunctions{}
}

// GetFileContentType ..
func (ff FileFunctions) GetContentType(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	contentTypeString := ff.GetContentTypeString([]string{contentType})
	return contentTypeString, nil
}

// GetContentTypeString ...
func (ff FileFunctions) GetContentTypeString(list []string) string {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	// https://www.iana.org/assignments/media-types/media-types.xhtml
	types := map[string]string{
		"application/zip":              "zip",
		"application/x-zip-compressed": "zip",
		"application/x-rar-compressed": "rar",
		"application/gzip":             "gz",
		"application/x-gzip":           "gz",
		"application/json":             "json",
		"application/pdf":              "pdf",
		"text/html; charset=UTF-8":     "html",
		"text/plain; charset=utf-8":    "txt",
	}
	for key, val := range types {
		match := Array().Contains(list, key, false)
		if match == -1 {
			continue
		}
		return val
	}
	return strings.Join(list, ";;")
}
