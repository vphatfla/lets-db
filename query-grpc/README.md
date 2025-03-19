# Protoc GRPC guides

## To generate go-specific protocol buffer code

```script 
protoc --go_out=. --go_opt=paths=source_relative \
   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    query/query.proto
``` 
