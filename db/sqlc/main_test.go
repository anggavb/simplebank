package db

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/jackc/pgx/v5"
)

const (
	dbHost     = "localhost"
	dbPort     = 5050
	dbName     = "grpc_bank"
	dbUser     = "grpc_bank"
	dbPassword = "secret"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	log.Println(buildDBURL())
	conn, err := pgx.Connect(context.Background(), buildDBURL())
	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}

func buildDBURL() string {
	return "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + strconv.Itoa(dbPort) + "/" + dbName + "?sslmode=disable"
}
