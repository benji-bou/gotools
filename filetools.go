package tools

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"os"
)

// later on integrate https://github.com/youtube/vitess/blob/master/go/cgzip/

func init() {
	// see https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT
	// section 4.4.5 for the available compression methods...which are not implemented by Golang....
	// only algorithm 0 and number 8(deflate) available
	//
}

var compresserror error = errors.New("Unable to generate a compressor")

//CompressFiles gzip files into result
func CompressFiles(files []string, result io.Writer) (int, error) {
	succeed := 0
	archive := zip.NewWriter(result)
	defer archive.Close()

	for _, filepath := range files {
		if file, errOpen := os.OpenFile(filepath, os.O_RDONLY, 0755); errOpen != nil {
			log.Println(errOpen)
			continue
		} else {
			defer file.Close()
			info, errStat := file.Stat()
			if errStat != nil {
				continue
			}
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				continue
			}

			header.Method = zip.Deflate

			zipper, err := archive.CreateHeader(header)
			if err != nil {
				continue
			}

			if _, errCp := io.Copy(zipper, file); errCp == nil {
				succeed++
			} else {
				log.Println(errCp)
			}
		}
	}
	return succeed, nil
}
