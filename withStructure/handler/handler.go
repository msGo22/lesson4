package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"week4/withStructure/domains"
)

type SchoolHandler struct {
	name     string
	students []domains.Student
}

func NewSchoolHandler(name string) *SchoolHandler {
	return &SchoolHandler{
		name: name,
		students: []domains.Student{
			{"Ali"},
			{"Cengiz"},
			{"Kemal"},
			{"Adem"},
		},
	}
}

func (h *SchoolHandler) GetSchoolInfo(writer http.ResponseWriter, request *http.Request) {
	// opsiyonel yol (json'a çevirip yollar)
	// json.NewEncoder(writer).Encode(h.name + " okulunun bilgilerine ulaştınız")
	fmt.Fprintf(writer, "%s okulunun bilgilerine ulaştınız", h.name)
}

func (h *SchoolHandler) GetAllStudents(writer http.ResponseWriter, request *http.Request) {
	// opsiyonel yol (json'a çevirip yollar)
	json.NewEncoder(writer).Encode(h.students)
	//fmt.Fprintf(writer, "%s", h.students)
}
