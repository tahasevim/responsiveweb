package handlers

import(
	"net/http"
	"encoding/json"
	"github.com/tahasevim/go-httpserver/templates"
)
type JsonMap map[string]interface{}

func RegisterHandlers(){
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/ip",ipHandler)
	http.HandleFunc("/headers",headersHandler)
	http.HandleFunc("/get",getHandler)
}

func ipHandler(w http.ResponseWriter, r *http.Request){
	jsonData := JsonMap{}
	jsonData["origin"] = r.RemoteAddr
	w.Write(makeJSONresponse(jsonData))
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	templates.IndexTemplate.ExecuteTemplate(w, "index", nil)
}

func headersHandler(w http.ResponseWriter,r *http.Request){
	jsonData := JsonMap{}
	jsonData["headers"] = initHeadMap(r)
	w.Write(makeJSONresponse(jsonData))
}

func getHandler(w http.ResponseWriter, r *http.Request){
	jsonData := JsonMap{}
	getData := JsonMap{}
	for k,v := range r.URL.Query(){
		getData[k] = v[0]
	}
	jsonData["args"] = getData
	jsonData["headers"] = initHeadMap(r)
	jsonData["origin"] = r.RemoteAddr
	jsonData["url"] = r.Host+r.URL.String()
	w.Write(makeJSONresponse(jsonData))
}

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