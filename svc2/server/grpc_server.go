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
	data.UnimplementedCreateEmpServiceServer
}


func (s *Server) SetConverter (c converter.Converter) {
	s.c = c
}

func (s *Server) Edit(ctx context.Context, emp *data.EmpReq) (*data.EmpResp, error) {
	log.Println("Edit called" ,emp)
	var res= &data.EmpResp{}
	if emp.FileType=="CSV" {
		emps, err := s.c.CSVtoJson()
		if err != nil {
			return nil, err
		}
		for i, e := range emps {
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
			_,err = s.c.JSONtoCSV(*e)
			log.Println(err)
		}
	} else {
		emps, err := s.c.XMLtoJSON()
		if err != nil {
			return nil, err
		}
		for i, e := range emps {
			if e.Id == emp.Emp.Id {
				emps[i].Id = emp.Emp.Id
				emps[i].Name = emp.Emp.Name
				emps[i].Address = emp.Emp.Address
				emps[i].Age = emp.Emp.Age
				res.Id = emp.Emp.Id
				break;
			}
		}
		if err := os.Truncate(s.c.Xmlfile, 0); err != nil {
			log.Printf("Failed to truncate: %v", err)
		}
		for _,e := range emps {
			empData := data.Emp{
				Id: e.Id,
				Name: e.Name,
				Age: e.Age,
				Address: e.Address,
			}
			_,err = s.c.JSONtoXML(empData)
			log.Println(err)
		}
	}

	return res,nil
}

func (s *Server) Get(context.Context, *data.NoArg) (*data.GetResp, error) {
	log.Println("Get func called")
	emps, err := s.c.CSVtoJson()
	empsXml,err := s.c.XMLtoJSON()
	for _,e := range empsXml {
		empData := data.Emp{
			Id: e.Id,
			Name: e.Name,
			Age: e.Age,
			Address: e.Address,
		}
		emps = append(emps,&empData)
		log.Println(err)
	}
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
	var id string
	var err error
	if req.FileType=="CSV" {
		id,err = s.c.JSONtoCSV(*emp)
	} else {
		id,err = s.c.JSONtoXML(*emp)
	}
	if err!=nil {
		log.Println(err)
		return nil,err
	}
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
