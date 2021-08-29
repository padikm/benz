package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
	"svc2/converter"
	"svc2/svc2/data"
)

func init() {
	_, err := os.OpenFile("outputfile.csv",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}

type server struct {

}

func (*server) Edit(ctx context.Context, emp *data.EmpReq) (*data.EmpResp, error) {
	log.Println("Edit called" ,emp)
	emps := converter.CSVtoJson()
	var res = &data.EmpResp{}
	for i,e:= range emps {
		if e.Id == emp.Emp.Id {
			emps[i].Id = emp.Emp.Id
			emps[i].Name = emp.Emp.Name
			emps[i].Address = emp.Emp.Address
			emps[i].Age = emp.Emp.Age
			res.Id = emp.Emp.Id
			break;
		}
	}
	if err := os.Truncate("outputfile.csv", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
	for _,e := range emps {
		converter.JSONtoCSV(*e)
	}
	return res,nil
}

func (*server) Get(context.Context, *data.NoArg) (*data.GetResp, error) {
	log.Println("Get func called")
	resp := data.GetResp{
		Emp: converter.CSVtoJson(),
	}
	return &resp,nil
}

func (*server) Create( c context.Context, req *data.EmpReq) (*data.EmpResp, error) {
	log.Println("Create func called")
	emp := req.GetEmp()
	id := converter.JSONtoCSV(*emp)
	res := data.EmpResp{
		Id: id,
	}
	return &res ,nil
}


func StartGrpcServer() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatal("Failed ", err)
	}

	cert ,err := credentials.NewServerTLSFromFile("certs/server.crt","certs/server.pem")
	if err!=nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(grpc.Creds(cert))
	data.RegisterCreateEmpServiceServer(s,&server{})
	if err:=s.Serve(l);err!=nil {
		log.Fatal(err)
	}
}
