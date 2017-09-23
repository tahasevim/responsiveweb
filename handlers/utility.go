package handlers

import (
	"io"
	"bytes"
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

func initFilemap(r *http.Request) jsonMap{
	fileMap := jsonMap{}
	r.ParseMultipartForm(256)	
	if r.MultipartForm  == nil {
		return fileMap
	}
	for k := range r.MultipartForm.File{
		var buf bytes.Buffer
		file,header,_ := r.FormFile(k)
		io.Copy(&buf,file)
		fileMap[header.Filename] = buf.String()
	}
	return fileMap
}
func initFormMap(r *http.Request) jsonMap {
	formMap := jsonMap{}
	for k,v := range r.Form{
		formMap[k] = v[0]
	}
	return formMap
}
func getAllJSONdata(r * http.Request, keys ...string) jsonMap{
	jsonData := jsonMap{}
	for _, key := range keys{
		switch key {
		case "headers":
			jsonData["headers"] = initHeadMap(r)
		case "origin":
			jsonData["origin"] = r.RemoteAddr
		case "url":
			jsonData["url"] = r.Host+r.URL.String()
		case "json":
			jsonData["json"] = ""//fix
		case "method":
			jsonData["method"] = r.Method
		case "args":
			jsonData["args"] = initQueryMap(r)
		case "user-agent":
			jsonData["user-agent"] = r.Header.Get("user-agent")
		case "uuid":
			jsonData["uuid"], _ = exec.Command("uuidgen").Output()//search for better solution
		case "form":
			r.ParseForm()
			jsonData["form"] = initFormMap(r)
		case "files":
			jsonData["files"] = initFilemap(r)
		case "data":
			jsonData["data"] = "" //fix
		}
	}
	return jsonData
}