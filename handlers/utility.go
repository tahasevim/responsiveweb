package handlers

import (
	"io/ioutil"
	//"log"
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
	if r.ProtoMajor == 1{
		head["Connection"] = "close"
	}
	head["Host"] = r.URL.String()
	return head
}

func deleteJSONval(val []byte, keys ...string) []byte{
	var x map[string]interface{}
	json.Unmarshal(val,&x)
	h,_:= (x["headers"].(map[string]interface{}))
	delete(h,"Host")
	delete(h,"User-Agent")
	delete(h,"Accept-Encoding")	
	x["headers"] = interface{}(h)
	for _,key := range keys{
		delete(x,key)
	}
	result, _ :=json.MarshalIndent(x,"","  ")
	result = append(result,byte('\n'))
	return result
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
	r.ParseForm()	
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
			jsonData["origin"] = r.Header.Get("X-Forwarded-For")
		case "url":
			jsonData["url"] = r.Host+r.URL.String()
		case "json":
			jsonData["json"] = "" //if data can be encoded json,then encode it as json.
		case "method":
			jsonData["method"] = r.Method
		case "args":
			jsonData["args"] = initQueryMap(r)
		case "user-agent":
			jsonData["user-agent"] = r.Header.Get("user-agent")
		case "uuid":
			jsonData["uuid"], _ = exec.Command("uuidgen").Output()//search for better solution
		case "form":
			jsonData["form"] = initFormMap(r)
		case "files":
			jsonData["files"] = initFilemap(r)
		case "data":
			body,_ := ioutil.ReadAll(r.Body)
			jsonData["data"] = body //if data can be encoded json,then encode it as json.
		case "brotli":
			jsonData["brotli"] = true
		case "deflated":
			jsonData["deflated"] = true
		case "gzipped":
			jsonData["gzipped"] = true
		}
	}
	return jsonData
}