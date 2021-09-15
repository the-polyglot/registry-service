package db

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
	"github.com/rs/zerolog/log"
)

// func init() {
// 	// Use singular table names for psql
// 	orm.SetTableNameInflector(func(s string) string {
// 		return s
// 	})
// }

func InitPg() (con *pg.DB) {

	con = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		PoolSize: 50,
	})
	if con != nil {
		log.Error().Msg("cannot connect to database")
	}
	return
}