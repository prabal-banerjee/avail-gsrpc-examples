package main

import (
	"avail-gsrpc-examples/internal/extrinsics"
	"fmt"
	"log"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func main() {
	// This example shows how to subscribe to new blocks, as well as sending the extrinsic data.
	//
	// It displays the block number every time a new block is seen by the node you are connected to.
	//
	// The submitted exintrnsic is dumped to stdout if found in the new block
	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	log.Println("gsrpc connected to Substrate API...")

	sub, err := api.RPC.Chain.SubscribeNewHeads()
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
	log.Println("Subscribed to new headers...")

	count := 0

	submittedHash, err := extrinsics.SubmitData(api, "random test data")
	if err != nil {
		panic(err)
	}

	for {
		head := <-sub.Chan()
		fmt.Printf("Chain is at block: #%v\n", head.Number)
		count++

		blockHash, err := api.RPC.Chain.GetBlockHash(uint64(head.Number))
		if err != nil {
			panic(err)
		}
		fmt.Printf("Chain is at block: #%v with hash %v\n", head.Number, blockHash.Hex())

		ret, err := api.RPC.Chain.GetBlock(blockHash)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, extrinsic := range ret.Block.Extrinsics {
			extHash, err := types.GetHash(extrinsic)
			if err != nil {
				panic(err)
			}
			if extHash == submittedHash {
				fmt.Printf("SUCCESS!! Extrinsic data: %#v\n\n", extrinsic)
			}
		}

		if count == 10 {
			sub.Unsubscribe()
			break
		}
	}
}
