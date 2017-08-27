package handlers

import (
	"encoding/json"
	"net/http"
	"os/exec"
)
type jsonMap map[string]interface{}

func makeJSONresponse(v interface {}) []byte {
	jsonVal, _ :=json.MarshalIndent(v,"","  ")
	jsonVal = append(jsonVal,byte('\n'))
	return jsonVal
}

func initHeadMap(r * http.Request) jsonMap{
	head := jsonMap{}
	for k,v := range r.Header{
		head[k] = v[0]
	}
	return head
}

func initQueryMap(r * http.Request) jsonMap{
	getData := jsonMap{}
	for k,v := range r.URL.Query(){
		getData[k] = v[0]
	}
	return getData
}

func getAllJSONdata(r * http.Request, keys ...string) jsonMap{
	jsonData := jsonMap{}
	for _ ,key := range keys{
		switch key {
		case "headers":
			jsonData["headers"] = initHeadMap(r)
		case "origin":
			jsonData["origin"] = r.RemoteAddr
		case "url":
			jsonData["url"] = r.Host+r.URL.String()
		case "json":
			jsonData["json"] = ""
		case "method":
			jsonData["method"] = r.Method
		case "args":
			jsonData["args"] = initQueryMap(r)
		case "user-agent":
			jsonData["user-agent"] = r.Header.Get("user-agent")
		case "uuid":
			jsonData["uuid"], _ = exec.Command("uuidgen").Output()
		}
	}
	return jsonData
}