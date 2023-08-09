#!/bin/bash
index_name="products"

settings=$(cat settings.json | jq -c .settings)
mappings=$(cat mappings.json | jq -c .mappings)

# Send PUT request
curl -X PUT "http://localhost:9200/${index_name}" -H 'Content-Type: application/json' -d "{\"settings\": $settings, \"mappings\": $mappings}"


# curl -X DELETE "http://localhost:9200/products"