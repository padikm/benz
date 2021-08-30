package converter

import (
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
	id := c.JSONtoCSV(emp)
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
	id := c.JSONtoCSV(emp)
	if len(id)==0 {
		t.Fatal("Failed")
	}
	_, err = c.CSVtoJson()
	if err !=nil {
		t.Fatal("Failed")
	}
	os.Remove(c.CSVfile)
}
