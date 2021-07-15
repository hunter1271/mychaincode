package main

import (
	"context"
	"fmt"
	"github.com/s7techlab/cckit/gateway/service"
	sdkconfig "github.com/s7techlab/hlf-sdk-go/api/config"
	"github.com/s7techlab/hlf-sdk-go/client"
	"github.com/s7techlab/hlf-sdk-go/identity"
	"gitlab.com/hunter1271/greeting"
	"log"
	"time"
)

type (
	Identity struct {
		mspID, certPath, keyPath string
	}
)

var (
	config = Identity{
		mspID:    "Org1MSP",
		certPath: "/Users/ural/Projects/fabric-tour/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/tls/client.crt",
		keyPath:  "/Users/ural/Projects/fabric-tour/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/tls/client.key",
	}

	sdkConfig = sdkconfig.Config{
		Crypto: sdkconfig.CryptoConfig{
			Type: "ecdsa",
			Options: map[string]interface{}{
				"curve":              "P256",
				"hash":               "SHA2-256",
				"signatureAlgorithm": "SHA256",
			},
		},
		Orderers: []sdkconfig.ConnectionConfig{
			{
				Host: "127.0.0.1:7050",
				Timeout: sdkconfig.Duration{
					Duration: time.Second * 5,
				},
			},
		},
		Discovery: sdkconfig.DiscoveryConfig{
			Type: "local",
		},
		MSP: []sdkconfig.MSPConfig{
			{
				Name: "Org1MSP",
				Endorsers: []sdkconfig.ConnectionConfig{
					{
						Host: "localhost:7051",
					},
				},
			},
			{
				Name: "Org2MSP",
				Endorsers: []sdkconfig.ConnectionConfig{
					{
						Host: "localhost:9051",
					},
				},
			},
		},
		Pool: sdkconfig.PoolConfig{
			DeliverTimeout: sdkconfig.Duration{
				Duration: 0,
			},
		},
	}
)

func main() {
	id, err := identity.NewMSPIdentity(config.mspID, config.certPath, config.keyPath)
	if err != nil {
		log.Fatal(err)
	}

	sdk, err := client.NewCore(config.mspID, id, client.WithConfigRaw(sdkConfig))
	if err != nil {
		log.Fatal(err)
	}

	ccService := service.New(sdk)
	helloService := greeting.NewHelloServiceGateway(ccService, "mychannel", "greeting")

	res, err := helloService.SayHello(context.Background(), &greeting.HelloRequest{Greeting: "Alloha"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Reply)
}
