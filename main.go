package main 

import (
	"fmt"
	"net/http"
)
 
func main(){
  mux := http.NewServerMux()
  mux.HandleFunc("/",handleRoot)
}

func handleRoot(){

}