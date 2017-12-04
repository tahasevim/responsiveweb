package handlers

import(
	"net/http"
	"github.com/tahasevim/responsiveweb/templates"
	"log"
	"strconv"
	"strings"
	"encoding/json"
	"time"
)
//Registers handlers to the default server mux.
func RegisterHandlers(){
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/ip",ipHandler)
	http.HandleFunc("/headers",headersHandler)
	http.HandleFunc("/get",getHandler)
	http.HandleFunc("/user-agent",useragentHandler)
	http.HandleFunc("/uuid",uuidHandler)
	http.HandleFunc("/post",postHandler)
	http.HandleFunc("/delete",deleteHandler)
	http.HandleFunc("/put",putHandler)	
	http.HandleFunc("/anything",anythingHandler)
	http.HandleFunc("/anything/",anythingHandler)
	http.HandleFunc("/encoding/utf8",utf8Handler)
	http.HandleFunc("/gzip",gzipHandler)
	http.HandleFunc("/deflate",deflateHandler)	
	http.HandleFunc("/brotli",brotliHandler)
	http.HandleFunc("/status/",statusHandler)
	http.HandleFunc("/response-headers",responseHeaderHandler)
	http.HandleFunc("/redirect/",redirectMultiHandler)
	http.HandleFunc("/redirect-to",redirectToHandler)
	//http.HandleFunc("/relative-redirect",relativeRedHandler)
	//http.HandleFunc("/absolute-redirect",absoluteRedHandler)
	http.HandleFunc("/cookies",cookieHandler)
	http.HandleFunc("/cookies/",cookieSetDelhandler)
	http.HandleFunc("/basic-auth/",basicAuthHandler)	
	http.HandleFunc("/hidden-basic-auth",hiddenBasicAuthHandler)
	http.HandleFunc("/stream/",streamHandler)
	http.HandleFunc("/delay/",delayHandler)

}

func ipHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"origin")	
	w.Write(makeJSONresponse(jsonData))
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	templates.IndexTemplate.ExecuteTemplate(w, "index", nil)
}

func headersHandler(w http.ResponseWriter,r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"headers")	
	w.Write(makeJSONresponse(jsonData))
}

func getHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"args","headers","origin","url")
	w.Write(makeJSONresponse(jsonData))
}

func useragentHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"user-agent")	
	w.Write(makeJSONresponse(jsonData))
}

func uuidHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"uuid")	
	w.Write(makeJSONresponse(jsonData))
}

func postHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

func deleteHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "DELETE" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

func putHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "PUT" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

func patchHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "PATCH" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

func anythingHandler(w http.ResponseWriter, r *http.Request){
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))
}

func utf8Handler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	templates.Utf8Template.ExecuteTemplate(w,"utf8",nil)
}

func gzipHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"gzipped","headers","method","origin")
	w.Write(makeJSONresponse(jsonData))
	
}

func brotliHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"brotli","headers","method","origin")
	w.Write(makeJSONresponse(jsonData))	
}

func deflateHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"deflated","headers","method","origin")
	w.Write(makeJSONresponse(jsonData))	
}

func statusHandler(w http.ResponseWriter, r *http.Request){
	stat,_ := strconv.ParseInt(r.URL.Path[len("/status/"):],10,64)
	if int(stat)==0{
		w.WriteHeader(418)
		http.Redirect(w,r,"/status/418",418)
	}
	http.Redirect(w,r,"/status/"+string(stat),int(stat))
}

func responseHeaderHandler(w http.ResponseWriter, r *http.Request){
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
	w.Write(makeJSONresponse(jsonData))	
}

func redirectMultiHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	ntime, err := strconv.ParseInt(r.URL.Path[len("/redirect/"):],10,64)
	log.Println(ntime)
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

func redirectToHandler(w http.ResponseWriter, r *http.Request){
	var stat int
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	url := r.URL.Query().Get("url")
	statstr, _ := strconv.ParseInt(r.URL.Query().Get("status_code"),10,64)
	/*if err != nil {
		w.Write([]byte("Invalid status code"))
		return
	}*/
	if statstr == 0{
		stat = 302
	}else{
		stat = int(statstr)
	}
	http.Redirect(w,r,url,stat)
}
func cookieHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	jsonData := jsonMap{}
	cookieMap := jsonMap{}
	for _,cookie := range r.Cookies(){
		cookieMap[cookie.Name] = cookie.Value
	}
	jsonData["Cookies"] = cookieMap
	w.Write(makeJSONresponse(jsonData))
}
func cookieSetDelhandler(w http.ResponseWriter, r *http.Request){
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
func basicAuthHandler(w http.ResponseWriter, r *http.Request){
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
	w.Header().Set("WWW-Authenticate", `Basic realm="localhost:8080"`)	
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
func hiddenBasicAuthHandler(w http.ResponseWriter, r *http.Request){
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

func streamHandler(w http.ResponseWriter, r *http.Request){
	var n int
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
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

func delayHandler(w http.ResponseWriter, r *http.Request){
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
	jsonData = getAllJSONdata(r,"args","data","files","forms","headers","origin","url")
	w.Write(makeJSONresponse(jsonData))
	
}