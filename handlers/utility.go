package handlers

import (
	"strconv"
	"log"
	"io/ioutil"
	//"log"
	"io"
	"bytes"
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"

)

type jsonMap map[string]interface{}
func makeJSONresponse(v interface {}) []byte {
	jsonVal, err :=json.MarshalIndent(v,"","  ")
	if err != nil {
		return nil
	}
	jsonVal = append(jsonVal,byte('\n'))
	return jsonVal
}

func initHeadMap(r * http.Request,body []byte) jsonMap{
	head := jsonMap{}
	for k,v := range r.Header{
		head[k] = v[0]
	}
	head["Host"] = r.URL.String()
	if r.Method == "POST" || r.Method == "DELETE" || r.Method == "PUT"{
		head["Content-Length"] = strconv.Itoa(len(body))
	}
	if r.ProtoMajor == 1{
		head["Connection"] = "close"
	}
	return head
}

func deleteJSONval(val []byte, keys ...string) []byte{
	x := make(map[string]interface{})
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
func setCooki(w http.ResponseWriter, r * http.Request) jsonMap{
	cookieMap :=jsonMap{}
	jsonData := jsonMap{}
	for k,v := range r.URL.Query(){
		cook := &http.Cookie{Name:k,Value:v[0]}
		http.SetCookie(w,cook)
		cookieMap[k] = v[0]
	}
	for _,cookie := range r.Cookies(){
		cookieMap[cookie.Name] = cookie.Value
	}
	//cookieMap["k1"] = "v1"
	//cookieMap["k2"] = "v2"
	jsonData["cookies"] = cookieMap
	return jsonData
}
func delCooki(w http.ResponseWriter, r * http.Request){
	for k := range r.URL.Query(){
		for _,cookie := range r.Cookies(){
			if cookie.Name == k{
				cook := &http.Cookie{Name:k,MaxAge:-1,}
				http.SetCookie(w,cook)
			}
		}
	}

}
func getMapForJSON(str string) map[string]interface{}{
	var vals []string
	resp := make(map[string]interface{})
	vals = nil
	for _,v:= range str{
		switch v{
		case ':':
			vals = strings.Split(str,string(v))
			break
		case '=':
			vals = strings.Split(str,string(v))
			break
		}
	}
	if vals == nil {
		return nil
	}
	for i:=0;i<len(vals)-1;i+=2{
		log.Println(vals[i])
		log.Println(vals[i+1])
		vals[i] = strings.TrimLeft(vals[i],"{")
		vals[i+1] = strings.TrimRight(vals[i+1],"}")
		resp[vals[i]] = vals[i+1]
	}
	//log.Println(vals)
	//log.Println(resp)
	return resp
}
func getAllJSONdata(r *http.Request, keys ...string) jsonMap{
	jsonData := jsonMap{}
	var body []byte
	if r.Body == nil{
		body = []byte("")
	}else{
		body,_ = ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	
	//body :="repair"
	for _, key := range keys{
		switch key {
		case "headers":
			jsonData["headers"] = initHeadMap(r,body)
		case "origin":
			jsonData["origin"] = r.Header.Get("X-Forwarded-For")
		case "url":
			jsonData["url"] = r.Host+r.URL.String()
		case "json":
			js := getMapForJSON(string(body))
			jsonData["json"] = strings.TrimRight(string(makeJSONresponse(js)),"\n")
		case "method":
			jsonData["method"] = r.Method
		case "args":
			jsonData["args"] = initQueryMap(r)
		case "user-agent":
			jsonData["user-agent"] = r.Header.Get("user-agent")
		case "uuid":
			jsonData["uuid"], _ = exec.Command("uuidgen").Output()
		case "form":
			jsonData["form"] = initFormMap(r)
		case "files":
			jsonData["files"] = initFilemap(r)
		case "data":
			jsonData["data"] = string(body) //if data can be encoded json,then encode it as json.
		case "brotli":
			jsonData["brotli"] = true//search
		case "deflated":
			jsonData["deflated"] = true//search
		case "gzipped":
			jsonData["gzipped"] = true//search
		case "authenticated":
			jsonData["authenticated"] = true
		case "user":
			if len(r.URL.String())==12{
				jsonData["user"] = "user"
			} else{
				stripUser := r.URL.String()[len("/basic-auth/"):]
				jsonData["user"] = stripUser[:strings.Index(stripUser,"/")]
			}
		
		}
	}
	return jsonData
}
func check(username,password string,r *http.Request) bool{
	user,passwd,ok := r.BasicAuth()
	return user == username && password == passwd && ok
}