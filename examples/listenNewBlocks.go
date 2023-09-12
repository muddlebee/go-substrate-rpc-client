package main

import (
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
)

func main() {
	// This example shows how to instantiate a Substrate API and use it to connect to a node and retrieve balance
	// updates
	//
	// NOTE: The example runs until you stop it with CTRL+C

	api, err := gsrpc.NewSubstrateAPI("wss://westend.api.onfinality.io/public-ws")
	if err != nil {
		fmt.Printf("Error occurred while connecting: %s\n", err)
		return
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

		if count == 10 {
			sub.Unsubscribe()
			break
		}
	}
}
