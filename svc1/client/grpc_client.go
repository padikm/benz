package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"svc1/svc2/data"
)

var c *grpc.ClientConn

func IntiClient(client GrpcClient) {
	c = client.NewGrpcClient()
}
type GrpcClientImpl struct {

}
type GrpcClient interface {
	NewGrpcClient() *grpc.ClientConn
}

func GrpcClientCreateEmp(emp data.Emp) (*data.EmpResp,error){
	clientEmp := data.NewCreateEmpServiceClient(c)
	log.Println("Create grpc client ", c)
	req := data.EmpReq{
		Emp: &data.Emp{
			Name:    emp.Name,
			Age:     emp.Age,
			Address: emp.Address,
		},
	}
	log.Println("grpc create req ", req)
	res, err := clientEmp.Create(context.TODO(), &req)
	if err != nil {
		return nil,err
	}
	log.Println("Create resp ", res)
	return res,nil
}

func GrpcClientGetEmps() (*data.GetResp,error){
	clientEmp := data.NewCreateEmpServiceClient(c)
	log.Println("Created grpc client ", c)
	res, err := clientEmp.Get(context.TODO(),&data.NoArg{})
	if err != nil {
		return nil,err
	}
	log.Println(res)
	return res,nil
}

func GrpcClientEditEmps(emp data.Emp) (*data.EmpResp,error){
	clientEmp := data.NewCreateEmpServiceClient(c)
	//log.Println("Created grpc client ", c)
	req := data.EmpReq{
		Emp: &data.Emp{
			Id:emp.Id,
			Name:    emp.Name,
			Age:     emp.Age,
			Address: emp.Address,
		},
	}
	res, err := clientEmp.Edit(context.TODO(),&req)
	if err != nil {
		return nil,err
	}
	log.Println(res)
	return res,nil
}

func (GrpcClientImpl) NewGrpcClient() *grpc.ClientConn{
	creds,err := credentials.NewClientTLSFromFile("certs/ca.crt","")
	if err != nil {
		log.Fatal("ERRRR " , err)
	}
	opts := grpc.WithTransportCredentials(creds)
	c , err := grpc.Dial("localhost:50051",opts)
	if err != nil {
		log.Fatal(err)
	}
	return c
}