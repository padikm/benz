package converter

import (
	"log"
	"os"
	"svc2/svc2/data"
	"testing"
)

func TestConverter_JSONtoCSV(t *testing.T) {
	c := Converter{CSVfile:"csvfile_test.csv"}
	emp := data.Emp{
		Name: " dummy",
		Address:"dummy addr",
		Age: 26,
	}
	id ,_ := c.JSONtoCSV(emp)
	if len(id)==0 {
		t.Fatal("Failed")
	}
	os.Remove(c.CSVfile)
}

func TestConverter_CSVtoJson(t *testing.T) {
	c := Converter{CSVfile:"csvfile_test.csv"}
	_,err := c.CSVtoJson()
	if err==nil {
		t.Fatal("Failed")
	}
	emp := data.Emp{
		Name: " dummy",
		Address:"dummy addr",
		Age: 26,
	}
	id,_ := c.JSONtoCSV(emp)
	if len(id)==0 {
		t.Fatal("Failed")
	}
	_, err = c.CSVtoJson()
	if err !=nil {
		t.Fatal("Failed")
	}
	os.Remove(c.CSVfile)
}

func TestJSONtoXMLConverter(t *testing.T) {
	emp := data.Emp{
		Id : "123",
		Name: " dummy",
		Address:"dummy addr",
		Age: 26,
	}
	c := Converter{
		Xmlfile:"xmloutput.xml",
	}
	_,err := c.JSONtoXML(emp)
	if err!=nil {
		t.Fatal("Failed")
	}

	os.Remove(c.Xmlfile)
}

func TestConverter_XMLtoJSON(t *testing.T) {
	c := Converter{
		Xmlfile:"xmloutput.xml",
	}
	emp := data.Emp{
		Id : "123",
		Name: " dummy",
		Address:"dummy addr",
		Age: 26,
	}
	_,err := c.JSONtoXML(emp)
	res ,err := c.XMLtoJSON()
	if err!=nil {
		t.Fatal("Failed")
	}
	os.Remove(c.Xmlfile)
	log.Print(res)
}