package tools

import (
	"bytes"
	"encoding/json"
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
