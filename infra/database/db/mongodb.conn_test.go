package db_test

import (
	"context"
	"os"
	"testing"
	"time"

	pkgdb "delegacia.com.br/infra/database/db"
	pkgdl "delegacia.com.br/infra/dl"
	"github.com/stretchr/testify/assert"
)

func TestConnectMongDB(t *testing.T) {
	os.Setenv("APP_MODE", "debug")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	config := pkgdb.Config{
		Url:    pkgdl.GetEnv("MONGO_URL"),
		DBName: pkgdl.GetEnv("MONGO_DATABASE"),
		Ctx:    ctx,
	}

	db := pkgdb.GetMongoConnetion(config)
	defer db.Disconnect(ctx)

	assert.NotNil(t, db)

}
