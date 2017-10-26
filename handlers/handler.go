package handlers

import(
	"net/http"
	"github.com/tahasevim/responsiveweb/templates"
	//"log"
	"strconv"
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
	stat,err := strconv.ParseInt(r.URL.Path[len("/status/"):],10,64)
	if err != nil {
		return
	}
	w.WriteHeader(int(stat))
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
	if int(ntime)<=0{
		w.Write([]byte("Invalid n"))
		return
	}
	if err != nil {
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
	statstr, _:= strconv.ParseInt(r.URL.Query().Get("status_code"),10,64)
	if statstr == 0{
		stat = 302
	}else{
		stat = int(statstr)
	}
	http.Redirect(w,r,url,stat)
}
