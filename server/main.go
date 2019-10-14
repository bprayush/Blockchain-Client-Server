package main

import (
	context "context"
	"log"
	"net"

	"github.com/bprayush/blockchain-server-client/proto"
	"github.com/bprayush/blockchain-server-client/server/blockchain"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Port 8080 busy or unavailable.: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{
		BlockChain: blockchain.NewBlockchain(),
	})
	srv.Serve(listener)
}

// Server anonymous
type Server struct {
	BlockChain *blockchain.Blockchain
}

// AddBlock @parms: context.Context, *proto.AddBlockRequest @return: *proto.AddBlockResponse, error
func (s *Server) AddBlock(ctx context.Context, req *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.BlockChain.AddBlock(req.Data)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

// GetBlockchain @parms: context.Context, *proto.GetBlockchainRequest @return: *proto.GetBlockchainResponse, error
func (s *Server) GetBlockchain(ctx context.Context, req *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	rsp := new(proto.GetBlockchainResponse)
	for _, b := range s.BlockChain.Blocks {
		rsp.Blocks = append(rsp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash:          b.Hash,
			Data:          b.Data,
		})
	}

	return rsp, nil
}
