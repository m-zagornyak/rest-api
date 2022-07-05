package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/m-zagornyak/rest-api.git/internal/config"
	"github.com/m-zagornyak/rest-api.git/pkg/utils"
	_ "github.com/spf13/viper"
	"log"
	"time"
)

type ClientPostgres interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, arguments ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, arguments ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func Client(ctx context.Context, sc config.StorageConfig) (pool *pgxpool.Pool, err error) {
	//	var pool *pgxpool.Pool
	//	var err error

	dsn := fmt.Sprintf("posgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	err = utils.DoWithTries(func() error {

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, sc.MaxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgres")
	}

	return pool, nil
}
