.: generate

generate:
	@protoc-b2bchain --version
	@echo "greeting proto generation"
	@protoc-b2bchain -I=. \
    	--cc-gateway-b2bchain_out=logtostderr=true:. \
    	--grpc-gateway-b2bchain_out=logtostderr=true:. \
    	--swagger-b2bchain_out=logtostderr=true:. \
    	./schema.proto


#    	protoc-b2bchain -I=. \
        #	-I=../../../../vendor \
        #	-I=../../../../third_party/googleapis \
        #	-I=../../../../../  \
        #	--govalidators-b2bchain_out=. \
        #	--cc-gateway-b2bchain_out=logtostderr=true:. \
        #	--go-b2bchain_out=plugins=grpc:. \
        #	--grpc-gateway-b2bchain_out=logtostderr=true:. \
        #    --swagger-b2bchain_out=logtostderr=true:. \
        #	./*.proto