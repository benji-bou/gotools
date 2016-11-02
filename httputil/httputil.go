package httputil

import (
	"bytes"
	"encoding/json"
	"github.com/goware/urlx"
	"image"
	"io/ioutil"
	"log"
	"net/http"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

//JSONRequestDecode decode
func JSONRequestDecode(r *http.Request) interface{} {
	var v interface{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String() // Does a
	log.Println(s)
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Println(err)
		return nil
	}
	return v
}

func GetBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	return raw, err

}

func GetImage(url string) (image.Image, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	return image.Decode(resp.Body)
}

func UriSolver(urlPath, urlBase string) (string, error) {
	objUrlPath, err := urlx.Parse(urlPath)
	if err != nil {
		return "", err
	}
	objUrlBase, err := urlx.Parse(urlBase)
	if err != nil {
		return "", err
	}
	resUrl := objUrlBase.ResolveReference(objUrlPath)
	return resUrl.String(), nil
}

func UriCleaner(urlRaw string) (string, error) {
	objUrl, err := urlx.Parse(urlRaw)
	if err != nil {
		return urlRaw, err
	}
	if objUrl.Scheme == "" && objUrl.IsAbs() == true && objUrl.Opaque == "" {
		objUrl.Scheme = "http"
	}
	return objUrl.String(), nil
}
