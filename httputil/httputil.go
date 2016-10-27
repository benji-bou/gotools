package httputil

import (
	"bytes"
	"encoding/json"
	"image"
	"io/ioutil"
	"log"
	"net/http"
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
	// raw, err := GetBytes(url)
	// if err != nil {
	// 	return nil, "", err
	// }
	// return image.Decode(bytes.NewReader(raw))
}
