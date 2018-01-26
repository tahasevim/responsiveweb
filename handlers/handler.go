//Package handlers contains all endpoints's handlers.
//All endpoints based on HTTP methods and entire handlers operate coming requests according to it.
package handlers

import(
	"net/http"
	"github.com/tahasevim/responsiveweb/templates"
	"log"
	"strconv"
	"strings"
	"encoding/json"
	"time"
	"fmt"
	"math/rand"
)


//GetHandlers adds handlers to the a map and returns it.
func GetHandlers()map[string]func(http.ResponseWriter,*http.Request){
	handlerList := make(map[string]func(http.ResponseWriter,*http.Request))
	handlerList["/"] = IndexHandler
	handlerList["/ip"] = IpHandler
	handlerList["/headers"] = HeadersHandler
	handlerList["/get"] = GetHandler
	handlerList["/user-agent"] = UseragentHandler
	handlerList["/uuid"] = UuidHandler
	handlerList["/post"] = PostHandler
	handlerList["/delete"] = DeleteHandler
	handlerList["/put"] = PutHandler
	handlerList["/anything"] = AnythingHandler
	handlerList["/anything/"] = AnythingHandler
	handlerList["/encoding/utf8"] = Utf8Handler
	handlerList["/gzip"] = GzipHandler
	handlerList["/deflate"] = DeflateHandler
	handlerList["/brotli"] = BrotliHandler
	handlerList["/status/"] = StatusHandler
	handlerList["/response-headers"] = ResponseHeaderHandler
	handlerList["/redirect/"] = RedirectMultiHandler
	handlerList["/redirect-to"] = RedirectToHandler
	//handlerList["/relative-redirect"] = relativeRedHandler
	//handlerList["/absolute-redirect"] = absoluteRedHandler
	handlerList["/cookies"] = CookieHandler
	handlerList["/cookies/"] = CookieSetDelHandler
	handlerList["/basic-auth"] = BasicAuthHandler
	handlerList["/hidden-basic-auth"] = HiddenBasicAuthHandler
	handlerList["/stream/"] = StreamHandler
	handlerList["/delay/"] = DelayHandler
	handlerList["/html"] = HtmlHandler
	handlerList["/robots.txt"] = RobotsTextHandler
	handlerList["/deny"] = DenyHandler
	handlerList["/cache"] = CacheHandler
	handlerList["/cache/"] = CacheControlHandler
	handlerList["/bytes/"] = BytesHandler
	handlerList["/links/"] = LinkHandler
	handlerList["/image"] = ImageHandler
	handlerList["/image/png"] = PngHandler
	handlerList["/image/jpeg"] = JpegHandler
	handlerList["/image/webp"] = WebpHandler
	handlerList["/image/svg"] = SvgHandler
	handlerList["/forms/post"] = FormsHandler
	handlerList["/xml"] = XmlHandler
	return handlerList
}

//IpHandler handles a GET request and sends a response in JSON format that contains IP address of client which made request.
func IpHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"origin")	
	w.Write(makeJSONresponse(jsonData))
}

//IndexHandler handles a GET request and sends a HTML page that contains links of endpoints.
func IndexHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	templates.IndexTemplate.ExecuteTemplate(w, "index", nil)
}
//HeadersHandler handles a GET request and sends a response in JSON format that contains header of the coming request.
func HeadersHandler(w http.ResponseWriter,r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"headers")	
	w.Write(makeJSONresponse(jsonData))
}
//GetHandler handles a GET request and sends a response in JSON format that contains args,IP,headers,url of the coming request.
func GetHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"args","headers","origin","url")
	w.Write(makeJSONresponse(jsonData))
}

//UseragentHandler handles a GET request and sends a response in JSON format that contains user-agent of the coming request.
func UseragentHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"user-agent")	
	w.Write(makeJSONresponse(jsonData))
}

//UuidHandler handles a GET request and sends a response in JSON format that contains uuid (Universally unique identifier).
//uuid is obtained by operating system's "uuidgen" tool.
func UuidHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"uuid")	
	w.Write(makeJSONresponse(jsonData))
}

//PostHandler handles a POST request and sends a response in JSON format that contains args,data,files,form,headers,IP,url of the coming request.
func PostHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

//DeleteHandler handles a DELETE request and sends a response in JSON format that contains args,data,files,form,headers,IP,url of the coming request.
func DeleteHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "DELETE" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

//PutHandler handles a PUT request and sends a response in JSON format that contains args,data,files,form,headers,IP,url of the coming request.
func PutHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "PUT" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

//patchHandler handles a PATCH request and sends a response in JSON format that contains args,data,files,form,headers,IP,url of the coming request.
func patchHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "PATCH" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

//AnythingHandler can handle any type of request and sends a response in JSON format that contains args,data,files,form,headers,IP,url,method of the coming request.
func AnythingHandler(w http.ResponseWriter, r *http.Request){
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url","method")
	w.Write(makeJSONresponse(jsonData))
}

//Utf8Handler handles a GET request and sends a UTF8 encoded template that contains a lot of different UTF8 encoded characters.
func Utf8Handler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	templates.Utf8Template.ExecuteTemplate(w,"utf8",nil)
}

//GzipHandler handles a GET request and sends a response in JSON format that contains gzipped,headers,method,IP of the coming request.
func GzipHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"gzipped","headers","method","origin")
	w.Write(makeJSONresponse(jsonData))
	
}

//BrotliHandler handles a GET request and sends a response in JSON format that contains gzipped,headers,method,IP of the coming request.
func BrotliHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"brotli","headers","method","origin")
	w.Write(makeJSONresponse(jsonData))	
}

//DeflateHandler handles a GET request and sends a response in JSON format that contains gzipped,headers,method,IP of the coming request.
func DeflateHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"deflated","headers","method","origin")
	w.Write(makeJSONresponse(jsonData))	
}

//StatusHandler can handle any type of request and redirects the coming request to the that status's url.
func StatusHandler(w http.ResponseWriter, r *http.Request){
	stat,_ := strconv.ParseInt(r.URL.Path[len("/status/"):],10,64)
	if int(stat)==0{
		w.WriteHeader(418)
		http.Redirect(w,r,"/status/418",418)
	}
	http.Redirect(w,r,"/status/"+string(stat),int(stat))
}

//ResponseHeaderHandler handles a GET or POST request and sends a response in JSON format.
//It prepares a JSON response from url of the coming request.
func ResponseHeaderHandler(w http.ResponseWriter, r *http.Request){
	if !(r.Method == "GET" || r.Method == "POST"){
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := jsonMap{}
	for key,value := range r.URL.Query(){
		if len(value) == 1{
			jsonData[key] = value[0]
		}else {
			jsonData[key] = value
		}
	}
	jsonData["Content-Type"] = "application/json"
	w.Write(makeJSONresponse(jsonData))	
}

//RedirectMultiHandler handles a GET request and redirects the coming request n times.
func RedirectMultiHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	ntime, err := strconv.ParseInt(r.URL.Path[len("/redirect/"):],10,64)
	if int(ntime)<0{
		w.Write([]byte("Invalid n"))
		return
	}
	if int(ntime)== 0{
		http.Redirect(w,r,"/get",302)
	}
	if err != nil {
		w.Write([]byte("Invalid n"))
		return
	}
	for i:=0;i<int(ntime);i++{
		http.Redirect(w,r,"/get",302)	
	}
}

//RedirectToHandler handles a GET request and redirects the coming request to the given url parameter.
func RedirectToHandler(w http.ResponseWriter, r *http.Request){
	var stat int
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	url := r.URL.Query().Get("url")
	statstr, _ := strconv.ParseInt(r.URL.Query().Get("status_code"),10,64)
	if statstr == 0{
		stat = 302
	}else{
		stat = int(statstr)
	}
	http.Redirect(w,r,url,stat)
}

//CookieHandler handles a GET request and sends a response in JSON format that contains cookies.
func CookieHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	jsonData := jsonMap{}
	cookieMap := jsonMap{}
	for _,cookie := range r.Cookies(){
		cookieMap[cookie.Name] = cookie.Value
	}
	jsonData["cookies"] = cookieMap
	w.Write(makeJSONresponse(jsonData))
}

//CookieSetDelHandler handles a GET request and sends a response in JSON format that contains cookies.
//It sets or deletes cookies according to given url (/set or /delete).
func CookieSetDelHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	if r.URL.Path == "/cookies/set"{
		jsonData := setCooki(w,r)
		w.Write(makeJSONresponse(jsonData))
		return	
	}
	if r.URL.Path == "/cookies/delete"{
		delCooki(w,r)
		http.Redirect(w,r,"/cookies",302)
		return
	}
	http.Redirect(w,r,"/cookies",302)
}

//BasicAuthHandler handles a GET request and sends a response in JSON format.
//It recieves password and username from client and checks that it is valid or not.
func BasicAuthHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	jsonData := jsonMap{}
	jsonData = getAllJSONdata(r,"authenticated","user")
	if len(r.URL.String())==12{
		w.Write(makeJSONresponse(jsonData))
		return
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="localhost:8080"`)//localhost
	if strings.Count(r.URL.String()[len("/basic-auth/"):],"/") != 1 {
		http.Error(w,"Not Found",http.StatusNotFound)
		return
	}
	userAndpasswd := r.URL.String()[len("/basic-auth/"):]
	user := userAndpasswd[:strings.Index(userAndpasswd,"/")]
	pass := userAndpasswd[strings.Index(userAndpasswd,"/")+1:]
	if !check(user,pass,r){
		http.Error(w,"Unauthorised Attempt",http.StatusUnauthorized)
		return
	}

	w.Write(makeJSONresponse(jsonData))
	log.Println("User logged in:",user)
	
}
//HiddenBasicAuthHandler handles a GET request and sends a response in JSON format.
//It recieves password and username from client and checks that it is valid or not.
func HiddenBasicAuthHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="localhost:8080"`)	
	if len(r.URL.String())==12 || strings.Count(r.URL.String()[len("/basic-auth/"):],"/") != 1 {
		http.Error(w,"Not Found",http.StatusNotFound)
		return
	}
	userAndpasswd := r.URL.String()[len("/basic-auth/"):]
	user := userAndpasswd[:strings.Index(userAndpasswd,"/")]
	pass := userAndpasswd[strings.Index(userAndpasswd,"/")+1:]
	if !check(user,pass,r){
		http.Error(w,"Unauthorised Attempt",http.StatusUnauthorized)
		return
	} 
	http.Error(w,"Not Found",http.StatusNotFound)		
}

//StreamHandler handles a GET request and sends a response in JSON format that contains url,args,headers,IP of the coming request.
//It sends response n times.
func StreamHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	var n int
	nparam, err := strconv.ParseInt(r.URL.String()[len("/stream/"):],10,64)
	switch{
	case err != nil:
		n = 20
	case int(nparam)>100:
		n = 100
	case int(nparam)<=100:
		n = int(nparam)
	}
	jsonData := jsonMap{}
	jsonData = getAllJSONdata(r,"url","args","headers","origin")
	for i:=0;i<n;i++{
		jsonResp,_:= json.Marshal(jsonData)
		w.Write(jsonResp)
		w.Write([]byte("\n"))
	}
}

//DelayHandler handles a GET request and sends a response in JSON format that contains args,data,files,form,headers,IP,url of the coming request.
//It sends response with a delayed time according to given n.
func DelayHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	var n int
	nparam, err := strconv.ParseInt(r.URL.String()[len("/delay/"):],10,64)
	switch{
	case err != nil:
		n = 3
	case int(nparam)>10:
		n = 10
	case int(nparam)<=10:
		n = int(nparam)
	}
	time.Sleep(time.Second * time.Duration(n))
	jsonData := jsonMap{}
	jsonData = getAllJSONdata(r,"args","data","files","form","headers","origin","url")
	w.Write(makeJSONresponse(jsonData))
}

//HtmlHandler handles a GET request and sends a sample HTML template.
func HtmlHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	templates.SampleTemplate.ExecuteTemplate(w,"sample",nil)
}

//RobotsTextHandler handles a GET request and sends a message that contains some robots.txt rules.
func RobotsTextHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	w.Write([]byte("User-agent: *\nDisallow: /deny"))

}

//DenyHandler handles a GET request and sends a message which recites that denied by robots.txt rules.
func DenyHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	w.Write([]byte(` 
	  .-''''''-.
        .' _      _ '.
       /   O      O   \\
      :                :
      |                |
      :       __       :
        '.         	 .'
          '-......-'
     YOU SHOULDN'T BE HERE`))
}

//ImageHandler handles a GET request and redirects it to "https://httpbin.org/image". 
func ImageHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	http.Redirect(w,r,"https://httpbin.org/image",http.StatusOK)
}

//PngHandler handles a GET request and redirects it to "https://httpbin.org/image/png". 
func PngHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	http.Redirect(w,r,"https://httpbin.org/image/png",http.StatusOK)
}

//JpegHandler handles a GET request and redirects it to "https://httpbin.org/image/jpeg". 
func JpegHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	http.Redirect(w,r,"https://httpbin.org/image/jpeg",http.StatusOK)
}

//WebpHandler handles a GET request and redirects it to "https://httpbin.org/image/webp". 
func WebpHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	http.Redirect(w,r,"https://httpbin.org/image/webp",http.StatusOK)
}

//SvgHandler handles a GET request and redirects it to "https://httpbin.org/image/svg". 
func SvgHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	http.Redirect(w,r,"https://httpbin.org/image/svg",http.StatusOK)
}

//FormsHandler handles a GET request and a sends sample form template.
func FormsHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	templates.FormsTemplate.ExecuteTemplate(w,"forms",nil)
}

//XmlHandler handles a GET request and sends sample XML template.
func XmlHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	w.Header().Set("Content-Type","application/xml")
	templates.XmlTemplate.ExecuteTemplate(w,"xml",nil)
}

//LinkHandler handles a GET request and sends n numbers link.
func LinkHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	var n int
	strArr := strings.Split(r.URL.String(),"/")
	nparam,err := strconv.ParseInt(strArr[2],10,64)
	switch{
	case err != nil:
		n = 10
	case int(nparam)>200:
		n = 200
	case int(nparam)<=200:
		n = int(nparam)
	}
	var html []string
	html = append(html,"<html><head><title>Links</title></head><body>")
	for i:=0;i<n;i++{
		html = append(html,fmt.Sprintf(` <a href=/links/%d/%d> %d </a> `,n,i,i))
	}
	html = append(html,"</body></html>")
	resp := strings.Join(html,"")
	w.Write([]byte(resp))
}

//CacheHandler handles a GET request and sends a response in JSON format that contains url,args,header,IP of the coming request.
//If "If-Modified-Since" or "If-None-Match" header is provided,it returns 304 status code.
func CacheHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	jsonData := getAllJSONdata(r,"url","args","headers","origin")
	//http date and uuid check
	if r.Header.Get("If-Modified-Since" ) == "" && r.Header.Get("If-None-Match") == ""{
		//uuid :=getAllJSONdata(r,"uuid")["uuid"]
		w.Header().Set("Last-Modified","")
		//w.Header().Set("ETag",strconv.Itoa([]byte(uuid.([]uint8))))
		w.Write(makeJSONresponse(jsonData))
	}else{
		w.WriteHeader(304)
	}
}

//CacheControlHandler handles a GET request and sends a response in JSON format that contains url,args,headers,IP of the coming request.
//It sets a Cache-Control header for n seconds.
func CacheControlHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	//url check
	strArr := strings.Split(r.URL.String(),"/")
	nparam,err := strconv.ParseInt(strArr[2],10,64)
	if err != nil {
		nparam = 60
	}
	jsonData := getAllJSONdata(r,"url","args","headers","origin")
	w.Header().Set("Cache-Control",fmt.Sprintf("public, max-age=%d",nparam))
	w.Write(makeJSONresponse(jsonData))
}

//BytesHandler handles a GET request and sends a response that contains bytes which are generated n times randomly.
func BytesHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	var n int
	var byteArr []byte
	urlStrArr := strings.Split(r.URL.String(),"/")
	nparam, err := strconv.ParseInt(urlStrArr[2],10,64)
	switch{
	case err != nil:
		n = 1024
	case int(nparam)>100*1024:
		n = 100*1024
	case int(nparam)<=100*1024:
		n = int(nparam)
	}
	for i:=0;i<n;i++{
		byteArr = append(byteArr,byte(rand.Int()))
	}
	w.Header().Set("Content-Type","application/octet-stream")
	w.Write(byteArr)
}