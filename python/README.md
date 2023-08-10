# Protoの生成
```
python -m grpc_tools.protoc -I ../proto --python_out=./proto --pyi_out=./proto --grpc_python_out=./proto ../proto/search.proto
```