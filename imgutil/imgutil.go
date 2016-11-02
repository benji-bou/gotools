package imgutil

import (
	"bytes"
	"errors"
	"github.com/vincent-petithory/dataurl"
	"gotools/fileutil"
	"gotools/httputil"
	"image"
	"image/png"
	"log"
	"os"
)

var (
	ErrSubImgNotImplemented = errors.New("SubImager is not implemented in this image.Image")
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func Extract(i image.Image, r image.Rectangle) (image.Image, error) {
	if subImager, ok := i.(SubImager); ok {
		return subImager.SubImage(r), nil
	} else {
		return i, ErrSubImgNotImplemented
	}
}

func GetImageSrc(src string) (image.Image, string, error) {
	var img image.Image
	var format string
	var err error
	log.Print("GetImageSrc -> ", src)
	urlData, err := dataurl.DecodeString(src)
	if err != nil {
		img, format, err = httputil.GetImage(src)
		if err != nil {
			log.Println("GetImageSrc - httputil.GetImage(src) FAILED", err, src)
			return img, format, err
		}
		err = nil
	} else {
		img, format, err = image.Decode(bytes.NewReader(urlData.Data))
		if err != nil {
			log.Println("GetImageSrc - Decode urlData.Data  FAILED", err, src)
		}
	}
	return img, format, err
}

func DrawPngFile(title string, img image.Image) error {
	return fileutil.ActionOnFile(title, func(file *os.File) error {
		return png.Encode(file, img)
	})

}
