package handlers

import (
	"encoding/json"
	"net/http"
)
type JsonMap map[string]interface{}

func makeJSONresponse(v interface {}) []byte {
	jsonVal, _ :=json.MarshalIndent(v,"","  ")
	jsonVal = append(jsonVal,byte('\n'))
	return jsonVal
}

func initHeadMap(r * http.Request) JsonMap{
	head := JsonMap{}
	for k,v := range r.Header{
		head[k] = v[0]
	}
	return head
}