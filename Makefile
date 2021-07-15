.: generate

generate:
	protoc-b2bchain -I=. \
	   -I=/Users/ural/Projects/s7techlab/monorepo/third_party/googleapis \
	   -I=./vendor \
	   --go-b2bchain_out=plugins=grpc:. \
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