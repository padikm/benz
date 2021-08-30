package main

import (
	"svc2/converter"
	"svc2/server"
)

func main()  {
	srv := server.Server{}
	srv.SetConverter(converter.Converter{
		CSVfile: "csvoutput.csv",
	})
	server.StartGrpcServer(srv)
}
