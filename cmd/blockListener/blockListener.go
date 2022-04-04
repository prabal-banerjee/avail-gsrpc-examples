package main

import (
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
)

func main() {
	// This example shows how to subscribe to new blocks.
	//
	// It displays the block number every time a new block is seen by the node you are connected to.
	//
	// NOTE: The example runs until 10 blocks are received or until you stop it with CTRL+C

	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}

	sub, err := api.RPC.Chain.SubscribeNewHeads()
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	count := 0

	for {
		head := <-sub.Chan()
		fmt.Printf("Chain is at block: #%v\n", head.Number)
		count++

		hash, err := api.RPC.Chain.GetBlockHash(uint64(head.Number))
		if err != nil {
			panic(err)
		}
		fmt.Printf("Chain is at block: #%v with hash %v\n", head.Number, hash.Hex())

		ret, err := api.RPC.Chain.GetBlock(hash)
		if err != nil {
			panic(err)
		}
		fmt.Println("Get block value: ", fmt.Sprint(ret))

		if count == 10 {
			sub.Unsubscribe()
			break
		}
	}
}