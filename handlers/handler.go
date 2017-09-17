package handlers

import(
	"net/http"
	"github.com/tahasevim/responsiveweb/templates"
)

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
