protoc  --go_out=./  \
        --go-grpc_out=./ \
        user.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./user.proto