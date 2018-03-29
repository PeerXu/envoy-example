package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pb "github.com/PeerXu/envoy-example/google-grpc-client/protos"
)

var (
	addr string
	name string
	bind string
)

var (
	rootCmd = &cobra.Command{
		Use:   "greet",
		Short: "greet service with envoy proxy",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			conn, err := grpc.Dial(addr, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("[E] failed to dial to service: %v\n", err)
			}
			defer conn.Close()

			cli := pb.NewGreetServiceClient(conn)
			req := pb.GreetRequest{Value: name}
			res, err := cli.Greet(ctx, &req)
			if err != nil {
				log.Fatalf("[E] failed to greet: %v\n", err)
			}
			log.Printf("[!] %v\n", res.Value)
		},
	}

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "start a greet service",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runGRPC(); err != nil {
				log.Fatalf("[E] failed to serve: %v\n", err)
			}
		},
	}
)

type greetService struct{}

func (s *greetService) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	val := fmt.Sprintf("Hello, %v!", req.Value)
	res := pb.GreetResponse{Value: val}
	return &res, nil
}

func newGreetService() *greetService {
	log.Printf("[!] create greet service\n")
	return &greetService{}
}

func runGRPC() error {
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("[E] failed to listen: %v\n", err)
	}
	srv := grpc.NewServer()
	pb.RegisterGreetServiceServer(srv, newGreetService())
	log.Printf("[!] Listen on %v\n", bind)
	return srv.Serve(lis)
}

func main() {
	rootCmd.PersistentFlags().StringVar(&addr, "addr", "127.0.0.1:5555", "greet service address")
	rootCmd.Flags().StringVar(&name, "name", "Alice", "greet to {name}")

	serveCmd.Flags().StringVar(&bind, "bind", "127.0.0.1:5555", "greet service binding address")

	rootCmd.AddCommand(serveCmd)

	rootCmd.Execute()
}
