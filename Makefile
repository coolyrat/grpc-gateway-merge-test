.PHONY: protoc
protoc:
	protoc -I=. \
	-I=third_party \
	--gogo_out=plugins=grpc,\
	Mthird_party/google/protobuf/empty.proto=github.com/gogo/protobuf/types:\
	pb \
	--grpc-gateway_out=allow_patch_feature=false,\
	Mthird_party/google/protobuf/empty.proto=github.com/gogo/protobuf/types:\
	pb \
	--swagger_out=logtostderr=true,allow_merge=true:api \
	*.proto