
gen:
	@protoc \
		--proto_path=chat "chat/chat.proto" \
		--go_out=chat/gen --go_opt=paths=source_relative \
  	--go-grpc_out=chat/gen --go-grpc_opt=paths=source_relative