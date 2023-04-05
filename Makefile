all: ms-users

ms-users:
	protoc \
		-I ${API_COMMON_PROTOS} \
		-I ${APIS} \
		--plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=${GOPATH}/bin/protoc-gen-go-grpc \
		--plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro \
		--micro_out=. \
        --micro_opt=module=${GENPROTO_GO_MODULE} \
		--go_out=. \
		--go_opt=module=${GENPROTO_GO_MODULE} \
		--go-grpc_out=. \
		--go-grpc_opt=module=${GENPROTO_GO_MODULE} \
		--grpc-gateway_opt=module=${GENPROTO_GO_MODULE} \
        --grpc-gateway_out=. \
		${APIS}/crea/ms-users/v1/*.proto
