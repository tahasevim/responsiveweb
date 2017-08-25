package main

import(
	"log"
	"net/http"
	"github.com/tahasevim/go-httpserver/handlers"
	"flag"
)
func main(){
	p := flag.String("port","8080","holds port")
	flag.Parse()
	handlers.RegisterHandlers()
	log.Println("Server started to listening at port: "+ *p)
	log.Println(http.ListenAndServe(":"+*p,nil))
}
