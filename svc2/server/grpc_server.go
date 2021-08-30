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


type Server struct {
	c converter.Converter
}

func (s *Server) SetConverter (c converter.Converter) {
	s.c = c
}

func (s *Server) Edit(ctx context.Context, emp *data.EmpReq) (*data.EmpResp, error) {
	log.Println("Edit called" ,emp)

	emps,err := s.c.CSVtoJson()
	if err!=nil {
		return nil,err
	}
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
	if err := os.Truncate(s.c.CSVfile, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
	for _,e := range emps {
		s.c.JSONtoCSV(*e)
	}
	return res,nil
}

func (s *Server) Get(context.Context, *data.NoArg) (*data.GetResp, error) {
	log.Println("Get func called")
	emps, err := s.c.CSVtoJson()
	if err!=nil {
		return nil,err
	}
	resp := data.GetResp{
		Emp: emps,
	}
	return &resp,nil
}

func (s *Server) Create( c context.Context, req *data.EmpReq) (*data.EmpResp, error) {
	log.Println("Create func called")
	emp := req.GetEmp()
	id := s.c.JSONtoCSV(*emp)
	res := data.EmpResp{
		Id: id,
	}
	return &res ,nil
}


func StartGrpcServer(srv Server) {
	l, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatal("Failed ", err)
	}

	cert ,err := credentials.NewServerTLSFromFile("certs/Server.crt","certs/Server.pem")
	if err!=nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.Creds(cert))
	data.RegisterCreateEmpServiceServer(s,&srv)
	if err:=s.Serve(l);err!=nil {
		log.Fatal(err)
	}
}
