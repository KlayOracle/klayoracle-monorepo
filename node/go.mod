module github.com/klayoracle/klayoracle-monorepo/node

go 1.18

require (
	github.com/PaesslerAG/jsonpath v0.1.1
	github.com/cockroachdb/cockroach-go/v2 v2.3.3
	github.com/google/uuid v1.3.0
	github.com/jackc/pgx/v4 v4.18.1
	github.com/joho/godotenv v1.5.1
	github.com/klayoracle/klayoracle-monorepo/data-provider v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.1
	go.uber.org/zap v1.24.0
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/PaesslerAG/gval v1.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.2 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)

replace github.com/klayoracle/klayoracle-monorepo/data-provider => ../data-provider
