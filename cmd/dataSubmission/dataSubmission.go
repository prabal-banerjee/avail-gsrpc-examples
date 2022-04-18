package main

import (
	"avail-gsrpc-examples/internal/extrinsics"
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
)

func main() {
	// This sample shows how to create a transaction to make a Avail data submission
	// Using appID=0

	// Instantiate the API
	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}

	// Set data and appID according to need
	data := "this_is_a_random_blob_of_data"
	_, err = extrinsics.SubmitData(api, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data submitted by Alice: %v against appID %v\n", data, 0)
}
