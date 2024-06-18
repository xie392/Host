package conf_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/xie392/restful-api/conf"
	"os"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/config.toml")
	if should.NoError(err) {
		should.Equal("host", conf.C().App.Name)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	err := os.Setenv("MYSQL_DATABASE", "config")
	if err != nil {
		return
	}

	err = conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("config", conf.C().MySQL.Database)
	}
}

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/config.toml")
	if should.NoError(err) {
		conf.C().MySQL.GetDB()
	}
}
