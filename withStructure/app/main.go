package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"week4/withStructure/handler"
)

func main() {
	// tanımlama yaptık
	r := mux.NewRouter()
	// handler oluşturuyoruz:
	schoolHandler := handler.NewSchoolHandler("necati bey")
	r.HandleFunc("/okul", schoolHandler.GetSchoolInfo).Methods(http.MethodGet)
	r.HandleFunc("/ogrenciler", schoolHandler.GetAllStudents).Methods(http.MethodGet)
	log.Println("Okul Sunucusu Ayağa kalktı...")
	log.Fatal(http.ListenAndServe(":9090", r))
}
