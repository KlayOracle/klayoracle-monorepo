package storage

import (
	"context"
	"os"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/klayoracle/klayoracle-monorepo/node/config"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"google.golang.org/protobuf/encoding/protojson"
)

var Conn *pgx.Conn

var ConnCtx context.Context

func CreateDBConn() {
	dns := os.Getenv("COCKROACH_DNS_URL")
	ConnCtx = context.Background()
	conn, err := pgx.Connect(ConnCtx, dns)
	if err != nil {
		config.Loaded.Logger.Fatalw("failed to connect database", err)
	}

	Conn = conn
}

func StoreJob(provider string, adapter *protonode.Adapter) {
	jsonData, err := protojson.Marshal(adapter)
	if err != nil {
		config.Loaded.Logger.Warnw("cannot store job info", "reason", err)
	}

	id := uuid.New()

	err = crdbpgx.ExecuteTx(context.Background(), Conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ConnCtx,
			"INSERT INTO node_jobs (id, host_ip, request, period) VALUES ($1, $2, $3, $4)", id, provider, string(jsonData), time.Now()); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		config.Loaded.Logger.Fatal(err)
	}
}
