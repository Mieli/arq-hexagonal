package dl_test

import (
	"testing"

	pkgdl "delegacia.com.br/infra/dl"
	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {

	port := "PORT"
	mongoPort := "MONGO_PORT"
	mongoUrl := "MONGO_URL"
	valuePort := pkgdl.GetEnv(port)
	valueMongoPort := pkgdl.GetEnv(mongoPort)
	valueMongoUrl := pkgdl.GetEnv(mongoUrl)

	assert.NotEmpty(t, valuePort, "3300")
	assert.NotEmpty(t, valueMongoPort, "27017")
	assert.NotEmpty(t, valueMongoUrl, "http://localhost:27017")
}
