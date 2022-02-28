package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type GunzipFunctions struct{}

func Zip() GunzipFunctions {
	return GunzipFunctions{}
}

// NewGz runs gunzip and saves unzipped file to destination
func (g GunzipFunctions) Gunzip(src, dest string) error {
	handlerSrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer handlerSrc.Close()

	gzr, err := gzip.NewReader(handlerSrc)
	if err != nil {
		return err
	}
	defer gzr.Close()

	handlerDest, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer handlerDest.Close()

	_, err = io.Copy(handlerDest, gzr)
	if err != nil {
		return err
	}

	return nil
}

// NewTarGz runs gunzip, untars and saves unzipped file to destination
func (g GunzipFunctions) UntarGunzip(src, dest string) error {
	handlerSrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer handlerSrc.Close()

	gzr, err := gzip.NewReader(handlerSrc)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tarr := tar.NewReader(gzr)
	for {
		header, err := tarr.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}

		if strings.Contains(header.Name, "._") {
			// To ignore AppleDouble encoded Macintosh file
			continue
		}

		target := filepath.Join(dest, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			_, err := os.Stat(target)
			if err != nil {
				err := os.MkdirAll(target, 0755)
				if err != nil {
					return err
				}
			}
		case tar.TypeReg:
			handler, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			_, err = io.Copy(handler, tarr)
			if err != nil {
				return err
			}
			handler.Close()
		}
	}
}
