package main

import (
	"context"
	"time"

	"delegacia.com.br/cmd/app"
	pkgdb "delegacia.com.br/infra/database/db"
	pkgdl "delegacia.com.br/infra/dl"
)

func main() {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	config := pkgdb.Config{
		Url:    pkgdl.GetEnv("MONGO_URL"),
		DBName: pkgdl.GetEnv("MONGO_DATABASE"),
		Ctx:    ctx,
	}
	db := pkgdb.GetMongoConnetion(config)
	defer db.Disconnect(ctx)

	app.Start(db)
}
