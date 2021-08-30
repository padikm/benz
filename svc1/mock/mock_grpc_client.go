package mock

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"svc1/svc2/data"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
type server struct{}
type MockGrpcClient struct {

}

func (*server) Edit(ctx context.Context, emp *data.EmpReq) (*data.EmpResp, error) {
	log.Println("Edit called" ,emp)
	return &data.EmpResp{},nil
}

func (*server) Get(context.Context, *data.NoArg) (*data.GetResp, error) {
	log.Println("Get func called")

	return &data.GetResp{},nil
}

func (*server) Create( c context.Context, req *data.EmpReq) (*data.EmpResp, error) {

	return &data.EmpResp{} ,nil
}


func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	data.RegisterCreateEmpServiceServer(s,&server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func(MockGrpcClient) NewGrpcClient() *grpc.ClientConn{
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}
	return conn
}
