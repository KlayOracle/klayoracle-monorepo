module github.com/klayoracle/klayoracle-monorepo/node

go 1.18

require (
	github.com/joho/godotenv v1.5.1
	github.com/klayoracle/klayoracle-monorepo/data-provider v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.1
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)

replace github.com/klayoracle/klayoracle-monorepo/data-provider => ../data-provider
