package main

import (
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"runtime/debug"
)

func main() {
	// This example shows how to instantiate a Substrate API and use it to connect to a node and retrieve balance
	// updates
	//
	// NOTE: The example runs until you stop it with CTRL+C
	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		fmt.Println("Error occurred:", err)
		debug.PrintStack()
		return
	}

	res, err := api.RPC.Author.PendingExtrinsics()
	fmt.Printf("Pending extrinsics: %v\n", res)
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}
	fmt.Printf("meta: %v\n", meta)

}
