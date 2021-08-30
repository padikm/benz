package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"svc1/client"
	"svc1/mock"
	"testing"
)

func MockServer() *httptest.Server {
	ts := httptest.NewServer(NewMux())
	mc := mock.MockGrpcClient{}
	client.IntiClient(mc)
	return ts
}

func TestGetEmp(t *testing.T) {
	ts := MockServer()
	defer ts.Close()
	log.Println(ts.URL)
	res, err := http.Get(ts.URL + "/employees")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(string(b))
}

func TestEditEmp(t *testing.T) {
	ts := MockServer()
	log.Println(ts.URL)
	defer ts.Close()
	client := &http.Client{}
	putreq := ` {
            "id": "67a32827-ac63-4664-9cc3-5bee94cf39c0",
            "name": "don",
            "age": 28,
            "address": "bangalore"
      }`
	reader := strings.NewReader(putreq)
	req, err := http.NewRequest(http.MethodPut, ts.URL+"/employees/123", reader)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(string(b))
	putreq = ` {
            "id": "67a32827-ac63-4664-9cc3-5bee94cf39c0",
            "name": "don",
            "age": 28,
            "address": "bangalore.,,
      }`
	reader = strings.NewReader(putreq)
	req, err = http.NewRequest(http.MethodPut, ts.URL+"/employees/123", reader)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	resp, err = client.Do(req)
	if resp.StatusCode!=http.StatusInternalServerError {
		t.Fatal("Failed")
	}

}

func TestCreateEmp(t *testing.T) {
	ts := MockServer()
	log.Println(ts.URL)
	defer ts.Close()
	client := &http.Client{}
	putreq := `{
            "name": "don",
            "age": 28,
            "address": "bangalore"
      }`
	reader := strings.NewReader(putreq)
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/employees", reader)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(string(b))

	putreq = `{
            "name": "don",
            "age": 28,
            "address": "bangalo.,
      }`

	reader = strings.NewReader(putreq)
	req, err = http.NewRequest(http.MethodPost, ts.URL+"/employees", reader)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	resp, err = client.Do(req)
	if resp.StatusCode!=http.StatusInternalServerError {
		t.Fatal("Failed")
	}
}

