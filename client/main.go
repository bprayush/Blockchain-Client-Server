package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/bprayush/blockchain-server-client/proto"

	"google.golang.org/grpc"
)

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "list the blocks")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial to the server: %v", err)
	}

	client := proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock(client)
	}

	if *listFlag {
		listBlock(client)
	}
}

func addBlock(client proto.BlockchainClient) {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})

	if err != nil {
		log.Fatalf("Unable to add block: %v\n", err)
	}

	log.Printf("New block hash: %v\n", block.Hash)
}

func listBlock(client proto.BlockchainClient) {
	bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})

	if err != nil {
		log.Fatalf("Unable to get block: %v\n", err)
	}

	// log.Printf("Blockchains: %v\n", bc.Blocks)

	for _, block := range bc.Blocks {
		log.Printf("Hash: %v Prev Hash: %v Data: %v\n", block.Hash, block.PrevBlockHash, block.Data)
	}
}
