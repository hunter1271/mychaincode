// Code generated by protoc-gen-cc-gateway. DO NOT EDIT.
// source: schema.proto

/*
Package greeting contains
  *   chaincode interface definition
  *   chaincode gateway definition
  *   chaincode service to cckit router registration func
*/
package greeting

import (
	context "context"

	cckit_gateway "github.com/s7techlab/cckit/gateway"
	cckit_ccservice "github.com/s7techlab/cckit/gateway/service"
	cckit_router "github.com/s7techlab/cckit/router"
	cckit_param "github.com/s7techlab/cckit/router/param"
	cckit_defparam "github.com/s7techlab/cckit/router/param/defparam"
)

// HelloServiceChaincode  method names
const (
	HelloServiceChaincode_SayHello = "SayHello"
)

// HelloServiceChaincodeResolver interface for service resolver
type HelloServiceChaincodeResolver interface {
	HelloServiceChaincode(ctx cckit_router.Context) (HelloServiceChaincode, error)
}

// HelloServiceChaincode chaincode methods interface
type HelloServiceChaincode interface {
	SayHello(cckit_router.Context, *HelloRequest) (*HelloResponse, error)
}

// RegisterHelloServiceChaincode registers service methods as chaincode router handlers
func RegisterHelloServiceChaincode(r *cckit_router.Group, cc HelloServiceChaincode) error {

	r.Query(HelloServiceChaincode_SayHello,
		func(ctx cckit_router.Context) (interface{}, error) {
			if v, ok := ctx.Param().(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return nil, cckit_param.PayloadValidationError(err)
				}
			}
			return cc.SayHello(ctx, ctx.Param().(*HelloRequest))
		},
		cckit_defparam.Proto(&HelloRequest{}))

	return nil
}

// NewHelloServiceGateway creates gateway to access chaincode method via chaincode service
func NewHelloServiceGateway(ccService cckit_ccservice.Chaincode, channel, chaincode string, opts ...cckit_gateway.Opt) *HelloServiceGateway {
	return &HelloServiceGateway{Gateway: cckit_gateway.NewChaincode(ccService, channel, chaincode, opts...)}
}

// gateway implementation
// gateway can be used as kind of SDK, GRPC or REST server ( via grpc-gateway or clay )
type HelloServiceGateway struct {
	Gateway cckit_gateway.Chaincode
}

// ServiceDef returns service definition
func (c *HelloServiceGateway) ServiceDef() cckit_gateway.ServiceDef {
	return cckit_gateway.ServiceDef{
		Desc:                        &_HelloService_serviceDesc,
		Service:                     c,
		HandlerFromEndpointRegister: RegisterHelloServiceHandlerFromEndpoint,
	}
}

// ApiDef deprecated, use ServiceDef
func (c *HelloServiceGateway) ApiDef() cckit_gateway.ServiceDef {
	return c.ServiceDef()
}

// Events returns events subscription
func (c *HelloServiceGateway) Events(ctx context.Context) (cckit_gateway.ChaincodeEventSub, error) {
	return c.Gateway.Events(ctx)
}

func (c *HelloServiceGateway) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Gateway.Query(ctx, HelloServiceChaincode_SayHello, []interface{}{in}, &HelloResponse{}); err != nil {
		return nil, err
	} else {
		return res.(*HelloResponse), nil
	}
}
