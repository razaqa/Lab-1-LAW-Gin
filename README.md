# Lab-1-LAW-Gin

## How to run

1. ```go install```
2. ```go run main.go```


## How to access calculate

```bash
curl --location --request POST 'http://localhost:8080/calculate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "a": 7,
    "b": 8
}'
```