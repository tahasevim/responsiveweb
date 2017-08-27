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
}

func ipHandler(w http.ResponseWriter, r *http.Request){
	jsonData := getAllJSONdata(r,"origin")	
	w.Write(makeJSONresponse(jsonData))
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	templates.IndexTemplate.ExecuteTemplate(w, "index", nil)
}

func headersHandler(w http.ResponseWriter,r *http.Request){
	jsonData := getAllJSONdata(r,"headers")	
	w.Write(makeJSONresponse(jsonData))
}

func getHandler(w http.ResponseWriter, r *http.Request){
	jsonData := getAllJSONdata(r,"args","headers","origin","url")
	w.Write(makeJSONresponse(jsonData))
}

func useragentHandler(w http.ResponseWriter, r *http.Request){
	jsonData := getAllJSONdata(r,"user-agent")	
	w.Write(makeJSONresponse(jsonData))
}

func uuidHandler(w http.ResponseWriter, r *http.Request){
	jsonData := getAllJSONdata(r,"uuid")	
	w.Write(makeJSONresponse(jsonData))
}

