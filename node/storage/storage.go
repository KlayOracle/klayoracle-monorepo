package storage

import (
	"context"
	"fmt"
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

func StoreJob(adapter *protonode.Adapter) error {
	jsonData, err := protojson.Marshal(adapter)
	if err != nil {
		config.Loaded.Logger.Warnw("cannot store job info", "reason", err)
	}

	id := uuid.New()

	err = crdbpgx.ExecuteTx(context.Background(), Conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ConnCtx,
			"INSERT INTO node_jobs (id, adapter_id, request, period) VALUES ($1, $2, $3, $4)", id, adapter.AdapterId, string(jsonData), time.Now()); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		config.Loaded.Logger.Warn(err)
	}

	return nil
}

func StoreResponses(adapter *protonode.Adapter, responses []int64, roundAnswer int64) error {
	id := uuid.New()

	err := crdbpgx.ExecuteTx(context.Background(), Conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ConnCtx,
			"INSERT INTO node_responses (id, responses, round_answer, adapter_id, period) VALUES ($1, $2, $3, $4, $5)", id, responses, roundAnswer, adapter.AdapterId, time.Now()); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		config.Loaded.Logger.Warnw("error storing aggregated responses", "feed", adapter.Name, "err", err)
	}

	return nil
}

func FetchResponsesFromT0(adapter *protonode.Adapter, t0 time.Time) (pgx.Rows, error) {

	query := fmt.Sprintf("SELECT responses from node_responses WHERE adapter_id='%s' AND CAST(period as TIMESTAMP) >= '%s'", adapter.AdapterId, t0.Format("2006-01-02 15:04:05.840"))

	rows, err := Conn.Query(ConnCtx, query)
	if err != nil {
		config.Loaded.Logger.Warnw("error fetching TWAP", "feed", adapter.Name, "err", err)
	}

	return rows, nil
}

func FetchRoundSize(adapter *protonode.Adapter) (pgx.Rows, error) {

	query := fmt.Sprintf("SELECT COUNT(*) as count from node_responses WHERE adapter_id='%s'", adapter.AdapterId)

	rows, err := Conn.Query(ConnCtx, query)
	if err != nil {
		config.Loaded.Logger.Warnw("error fetching round size", "feed", adapter.Name, "err", err)
	}

	return rows, nil
}
