package server

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func RegisterGrpcServer(port string) error {
	//Initialize the lister
	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Error().Err(err).Msg("error while listening port")
	}
	//create new grpc server instance
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(50 * 1024 * 1024))
	log.Info().Msg("portal management service listening on port: 50051")
	
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Error().Err(err).Msg("error while running grpc  server")
	}
	return nil
}
