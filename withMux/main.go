package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// tanımlama yaptık
	r := mux.NewRouter()
	r.HandleFunc("/wiki/spaces/{space_id}/pages/{page_id}/{page_title_slug}", neCalisacak).Methods("GET")
	/// www.alameddin.com/...
	subRouter := r.PathPrefix("/adana").Subrouter()
	/// www.alameddin.com/adana/....
	//
	subRouter.HandleFunc("/ikinci", ikinciAdanali)
	subRouter.HandleFunc("/ucuncu", ikinciAdanali)
	subRouter.HandleFunc("/dorduncu", ikinciAdanali)
	subRouter.HandleFunc("/besince", ikinciAdanali)
	log.Println("Started")
	log.Fatal(http.ListenAndServe(":9091", r))

}

func neCalisacak(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "ilk fonksiyon çalıştı")
}

func ikinciAdanali(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "ikinci adanali fonksiyonu çalıştı")
}
