# simpleshop

An example of shop web service written by Go with Gin

## How it works

```bash
# requirement postgres, redis
# configure config.yaml
# run 
go run ./cmd/api/main.go

# run test
go test

# run test with coverage
go test -timeout 9000s -a -v -coverprofile=coverage.out -coverpkg=./... ./...
# or 
make unittest
```


## References

- Documentation at: `http://localhost:8888/swagger/index.html`
- Restful API, GRPC, DDD, Gorm, Swagger, Logging, Jwt-Go, Gin-gonic, Redis