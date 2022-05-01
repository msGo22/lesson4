package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "ooo Hoş Geldin. \n")
		return
	}
	writer.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(writer, "Hacker mısın?")
}

func headers(writer http.ResponseWriter, request *http.Request) {
	if val := request.Header.Get("alameddin"); val != "" {
		fmt.Println("Alameddin sunucuya giriş yaptı", val)
		fmt.Fprintf(writer, "%s geldi haberin olsun \n", val)
	}
	for name, headers := range request.Header {
		for _, h := range headers {
			fmt.Fprintf(writer, "%v: %v\n", name, h)
		}
	}
}

func methods(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%s\n", request.Method)
	switch request.Method {
	case http.MethodGet:
		fmt.Println("Get ile gelindi")
		fmt.Fprintf(writer, "%s", request.URL.Query()["adanali"])
	case http.MethodPost:
		type requestBody struct {
			Name          string `json:"name"`
			StudentNumber string `json:"student_number,omitempty"`
		}
		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Post ile gelindi")
		var reqBody requestBody
		json.Unmarshal(b, &reqBody)
		fmt.Fprintf(writer, "Name: %s\n", reqBody.Name)
		fmt.Fprintf(writer, "Student No: %s\n", reqBody.StudentNumber)
		return
	case http.MethodPut:
		fmt.Println("Put ile gelindi")
	case http.MethodDelete:
		fmt.Println("Delete ile gelindi")
	}
}

func main() {
	http.HandleFunc("/merhaba", hello)
	http.HandleFunc("/header-detay", headers)
	http.HandleFunc("/method-bilgi", methods)
	fmt.Println("Started....")
	http.ListenAndServe(":9900", nil)
}

func celik(w http.ResponseWriter, r *http.Request) {

}
