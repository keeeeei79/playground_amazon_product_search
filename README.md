# Playground for amazon product search dataset

## Prepare
- `python/notebooks/makeDataset.ipynb`でデータセットを作成し、`data/`以下にデータセットを格納する

## Indexing
### init elasticsearch
```
cd elasticsearch
docker compose up
./boot.sh
```

### start indexing
```
go run ./cmd i
```

## Search
### start server
```
go run ./cmd s
```
### grpcurl
```
grpcurl -plaintext -d '{"keyword": "sports"}' localhost:50051 search.SearchService.Search
```

## TODO
- Indexing
- Search
- Vector Search
- Reranking
- Query Expansion
- Boosting
- Query Auto Complete

## Protoの生成
```
# Go
protoc --go_out=. --go_opt=module=github.com/keeeeei79/playground_amazon_product_search --go-grpc_out=. --go-grpc_opt=module=github.com/keeeeei79/playground_amazon_product_search proto/search.proto
# Python
protoc -I . --python_out=. proto/search.proto
```
