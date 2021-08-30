package server

import (
	"context"
	"log"
	"os"
	"svc2/converter"
	"svc2/svc2/data"
	"testing"
)

func TestServer_Create(t *testing.T) {
	srv := Server{}
	srv.SetConverter(converter.Converter{
		CSVfile: "csvoutput_test.csv",
	})

	_ = os.Remove("csvoutput_test.csv")

	_ , err := srv.Get(context.TODO(),&data.NoArg{})
	if err==nil {
		t.Fatal("Expected err")
	}

	f,_  := os.OpenFile("csvoutput_test.csv",os.O_APPEND|os.O_CREATE|os.O_RDWR, 777)
	resp , err := srv.Get(context.TODO(),&data.NoArg{})

	if len(resp.Emp)!=0 {
		t.Fatal("failed err")
	}
	f.Close()
	err = os.Remove("csvoutput_test.csv")
	if err != nil {
		t.Error(err)
	}

}

func TestServer_Edit(t *testing.T) {
	srv := Server{}
	srv.SetConverter(converter.Converter{
		CSVfile: "csvoutput_test.csv",
	})

	_ = os.Remove("csvoutput_test.csv")
	emps := data.EmpReq{}
	emps.Emp = &data.Emp{
		Id: "1234",
		Name: "dummy",
		Age: 25,
		Address:"NA",
	}

	_ , err := srv.Edit(context.TODO(),&emps)
	if err==nil {
		t.Fatal("Expected err")
	}
	srv.Create(context.TODO(),&emps)
	resp , err := srv.Edit(context.TODO(),&emps)
	log.Print(resp.Id)
	if resp.Id!= "1234" {
		t.Fatal("failed err")
	}
	err = os.Remove("csvoutput_test.csv")
	if err != nil {
		t.Error(err)
	}

}
