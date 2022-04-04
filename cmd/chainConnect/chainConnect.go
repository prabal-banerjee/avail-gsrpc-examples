package main

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/config"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
)

func main() {
	// The following example shows how to instantiate a Substrate API and use it to connect to a node

	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}

	chain, err := api.RPC.System.Chain()
	if err != nil {
		panic(err)
	}
	nodeName, err := api.RPC.System.Name()
	if err != nil {
		panic(err)
	}
	nodeVersion, err := api.RPC.System.Version()
	if err != nil {
		panic(err)
	}

	fmt.Printf("You are connected to chain %v using %v v%v\n", chain, nodeName, nodeVersion)

}
