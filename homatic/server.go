package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func PairDeviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"active"}`))
}

// git init, git add server.go go.mod, git commit -m "[Nong] init project"
func main() {
	fmt.Println("hello hometic : I'm Gopher!!")
	r := mux.NewRouter()
	r.HandleFunc("/pair-device", PairDeviceHandler).Methods(http.MethodPost)
	
	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	fmt.Println("addr:", addr)
	
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}
	log.Println("starting...")
	log.Fatal(server.ListenAndServe())
}
