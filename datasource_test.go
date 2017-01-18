package common

import (
	"os"
	"testing"
)

func TestDatasource(t *testing.T) {

	var config *Configuration
	var err error

	p, _ := os.Getwd()

	t.Logf("CWD Is %s", p)

	config, err = GetConfig()
	if err != nil {
		t.Errorf("Error loading configuration: %s", err.Error())
		return
	}

	t.Run("SeedData() doesn't fail", func(t *testing.T) {
		// On seed une base de données de test
		err = SeedData(config.DatabaseDriver, config.ConnectionString, config.SeedDataPath)

		if err != nil {
			t.Errorf("Unexpected exception: %s", err.Error())
			return
		}
	})

	d := NewDatasource(config.DatabaseDriver, config.ConnectionString)

	t.Run("GetCurrentSeason() doesn't fail", func(t *testing.T) {
		_, err := d.GetCurrentSeason()

		if err != nil {
			t.Errorf("Unhandled exception: %s", err.Error())
		}
	})

	if config.DatabaseDriver == "sqlite3" {
		// Teardown de la BD de test
		os.Remove(config.ConnectionString)
	}
}
