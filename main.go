package main

import(
	"net/http"
	"fmt"
	"log"
)

type Engine struct{}

func (e *Engine)ServeHTTP(w http.ResponseWriter,r *http.Request){
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w,"URL.Path=%s\n",r.URL.Path)
	case "/hello":
		for k,v := range r.Header{
			fmt.Fprintf(w,"Header[%q]=%q\n",k,v)
		}
	default:
		fmt.Fprintf(w,"404 NOT FOUND: %s \n",r.URL)
	}
	
}

func main()  {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe("localhost:8080",engine))
}

func handler(w http.ResponseWriter, r *http.Request){
	
}