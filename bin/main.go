package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"gitlab.com/hunter1271/greeting"
)

func main() {
	cc, err := greeting.NewCC()
	if err != nil {
		fmt.Printf("Error creating Greeting chaincode: %s", err)
	}

	if err := shim.Start(cc); err != nil {
		fmt.Printf("Error starting Greeting chaincode: %s", err)
	}
}
