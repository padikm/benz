package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"svc1/client"
	"svc1/svc2/data"
	"time"
)

func NewMux() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/employees", GetEmp).Methods("GET")
	r.HandleFunc("/employees", CreateEmp).Methods("POST")
	r.HandleFunc("/employees/{id}", EditEmp).Methods("PUT")
	return r
}

func StartHttpServer() {
	srv := &http.Server{
		Handler:      NewMux(),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting server")
	log.Fatal(srv.ListenAndServe())
}

func CreateEmp(response http.ResponseWriter, request *http.Request) {
	b,err := io.ReadAll(request.Body)
	//if err!=nil {
	//	response.WriteHeader(http.StatusBadRequest)
	//	response.Write([]byte(err.Error()))
	//}
	emp := data.Emp{}
	err = json.Unmarshal(b, &emp)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	res,err := client.GrpcClientCreateEmp(emp)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	b, err = json.Marshal(&res)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(b)
	log.Println(emp)
}

func GetEmp(response http.ResponseWriter, request *http.Request) {
	res,err := client.GrpcClientGetEmps()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	b, err := json.Marshal(&res)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(b)
	log.Println(res)
}

func EditEmp(response http.ResponseWriter, request *http.Request) {
	b,err := io.ReadAll(request.Body)
	//if err!=nil {
	//	response.WriteHeader(http.StatusBadRequest)
	//	response.Write([]byte(err.Error()))
	//}
	emp := data.Emp{}
	err = json.Unmarshal(b, &emp)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	res,err := client.GrpcClientEditEmps(emp)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	b, err = json.Marshal(&res)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(b)
	log.Println(emp)
}
