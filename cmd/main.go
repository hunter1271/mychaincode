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
		Orderers: []sdkconfig.ConnectionConfig{
			{
				Host: "0.0.0.0:7053",
				Timeout: sdkconfig.Duration{
					Duration: time.Second * 5,
				},
			},
		},
		Discovery: sdkconfig.DiscoveryConfig{
			Type: "local",
			Options: map[string]interface{}{
				"channels": map[string]interface{}{
					"name": "mychannel",
					"chaincodes": map[string]interface{}{
						"name":   "greeting",
						"type":   "golang",
						"policy": "AND('Org1MSP.member','Org2MSP.member')",
					},
				},
			},
		},
		MSP: []sdkconfig.MSPConfig{
			{
				Name: "Org1MSP",
				Endorsers: []sdkconfig.ConnectionConfig{
					{
						Host: "0.0.0.0:7051",
					},
				},
			},
			{
				Name: "Org2MSP",
				Endorsers: []sdkconfig.ConnectionConfig{
					{
						Host: "0.0.0.0:9051",
					},
				},
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
	fmt.Println("final")
}
