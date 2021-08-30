package client

import (
	"svc1/mock"
	"svc1/svc2/data"
	"testing"
)


func TestGrpcClientCreateEmps(t *testing.T) {
	client := mock.MockGrpcClient{}
	IntiClient(client)
	defer c.Close()
	empResp,err := GrpcClientCreateEmp(data.Emp{})
	if  err!=nil && empResp.Id != "" {
		t.Error("Failed")
	}

}

func TestGrpcClientEditEmps(t *testing.T) {
	client := mock.MockGrpcClient{}
	IntiClient(client)
	defer c.Close()
	empResp,err := GrpcClientEditEmps(data.Emp{})
	if  err!=nil || empResp.Id != "" {
		t.Error("Failed")
	}

}

func TestGrpcClientGetEmps(t *testing.T) {
	client := mock.MockGrpcClient{}
	IntiClient(client)
	defer c.Close()
	empResp,err  := GrpcClientGetEmps()
	if  err!=nil ||  len(empResp.Emp) !=0 {
		t.Error("Failed")
	}
}