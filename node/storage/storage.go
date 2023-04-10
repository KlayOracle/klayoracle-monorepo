package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

var Conn *pgx.Conn

var ConnCtx context.Context

func CreateDBConn() {
	dns := os.Getenv("COCKROACH_DNS_URL")
	ConnCtx = context.Background()
	conn, err := pgx.Connect(ConnCtx, dns)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	Conn = conn
}
