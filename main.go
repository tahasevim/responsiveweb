//responsiveweb project is inspired by Kenneth Reitz's https://httpbin.org project.
//It is implemented with built-in HTTP library.
//All handlers will be registered to the DefaultServeMux.
package main

import(
	"log"
	"net/http"
	"github.com/tahasevim/responsiveweb/handlers"
	"flag"
)
func main(){
	p := flag.String("port","8080","holds port")
	flag.Parse()
	handlerList := handlers.GetHandlers()
	for url, handlerFunc := range handlerList{
		http.HandleFunc(url,handlerFunc)
	}
	log.Println("Server started to listening at port: "+ *p)
	log.Println(http.ListenAndServe(":"+*p,nil))
}
