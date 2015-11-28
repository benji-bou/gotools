package tools

import (
	"compress/gzip"
	"errors"
	"io"
	"log"
	"os"
)

// later on integrate https://github.com/youtube/vitess/blob/master/go/cgzip/

var compresserror error = errors.New("Unable to generate a compressor")

//CompressFiles gzip files into result
func CompressFiles(files []string, result io.Writer) (int, error) {
	succeed := 0
	zipper := gzip.NewWriter(result)
	for _, filepath := range files {
		if file, errOpen := os.OpenFile(filepath, os.O_RDONLY, 0755); errOpen != nil {
			log.Println(errOpen)
			continue
		} else {
			if _, errCp := io.Copy(zipper, file); errCp == nil {
				succeed++
			}
			file.Close()
		}
	}
	zipper.Flush()
	return succeed, nil
}
