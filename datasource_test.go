package common

import (
	"os"
	"testing"

	"github.com/kelseyhightower/envconfig"
)

type testCase struct {
	TestID   uint
	IsNil    bool
	ExpectID uint
}

type configuration struct {
	DatabaseDriver   string
	ConnectionString string
	SeedDataPath     string
}

func getConfig() (*configuration, error) {
	var c configuration

	err := envconfig.Process("TSAP", &c)

	if err != nil {
		return nil, err
	}

	return &c, err
}

func TestDatasource(t *testing.T) {

	config, err := getConfig()

	if err != nil {
		t.Errorf("Error loading configuration: %s", err.Error())
	}

	// On seed une base de données de test
	err = SeedData(config.DatabaseDriver, config.ConnectionString, config.SeedDataPath)

	if err != nil {
		t.Errorf("Unexpected exception: %s", err.Error())
	}

	d := NewDatasource(config.DatabaseDriver, config.ConnectionString)

	t.Run("GetCurrentSeason doesn't fail", func(t *testing.T) {
		_, err := d.GetCurrentSeason()

		if err != nil {
			t.Errorf("Unexpected exception: %s", err.Error())
		}
	})

	t.Run("GetCurrentSeason returns correct season", func(t *testing.T) {
		s, _ := d.GetCurrentSeason()

		if s.Annees != "2015-2016" {
			t.Errorf("Expected %s, got %s", "2015-2016", s.Annees)
		}
	})

	teamCases := []testCase{
		{TestID: 1, IsNil: false, ExpectID: 1},
		{TestID: 99999, IsNil: true, ExpectID: 1},
	}

	for _, c := range teamCases {
		t.Run("GetTeam() doesn't fail", func(t *testing.T) {
			_, err := d.GetTeam(c.TestID)

			if !c.IsNil && err != nil {
				t.Errorf("Unexpected exception: %s", err.Error())
			}
		})

		t.Run("GetTeam() returns correct team", func(t *testing.T) {
			team, _ := d.GetTeam(c.TestID)

			if !c.IsNil && (team.ID != c.ExpectID) {
				t.Errorf("Expected %d, got %d", c.ExpectID, team.ID)
			}
		})

		t.Run("GetTeam() returns nil when team not found", func(t *testing.T) {
			team, err := d.GetTeam(c.TestID)

			if c.IsNil && ((team != nil) || err == nil) {
				t.Errorf("Expected team to be Nil, got ID %d instead", team.ID)
			}
		})
	}

	playerCases := []testCase{
		{TestID: 101, IsNil: false, ExpectID: 101},
		{TestID: 99999, IsNil: true, ExpectID: 1},
	}

	for _, c := range playerCases {
		t.Run("GetPlayer() doesn't fail", func(t *testing.T) {
			_, err := d.GetPlayer(c.TestID)

			if !c.IsNil && err != nil {
				t.Errorf("Unexpected exception: %s", err.Error())
			}
		})

		t.Run("GetPlayer() returns correct player", func(t *testing.T) {
			player, _ := d.GetPlayer(c.TestID)

			if !c.IsNil && (player.ID != c.ExpectID) {
				t.Errorf("Expected %d, got %d", c.ExpectID, player.ID)
			}
		})

		t.Run("GetPlayer returns nil when player not found", func(t *testing.T) {
			player, err := d.GetPlayer(c.TestID)

			if c.IsNil && ((player != nil) || err == nil) {
				t.Errorf("Expected player to be Nil, got ID %d instead", player.ID)
			}
		})
	}

	matchCases := []testCase{
		{TestID: 3, IsNil: false, ExpectID: 3},
		{TestID: 99999, IsNil: true, ExpectID: 1},
	}

	for _, c := range matchCases {

		t.Run("GetLatestMatch() doesn't fail", func(t *testing.T) {
			_, err := d.GetLatestMatch(c.TestID)

			if !c.IsNil && err != nil {
				t.Errorf("Unexpected exception: %s", err.Error())
			}
		})

		t.Run("GetLatestMatch() returns correct match", func(t *testing.T) {
			match, _ := d.GetLatestMatch(c.TestID)

			if !c.IsNil && (match.EquipeMaisonID != int(c.ExpectID)) && (match.EquipeAdverseID != int(c.ExpectID)) {
				t.Errorf("Expected a team with this ID %d", c.ExpectID)
			}
		})

		t.Run("GetLatestMatch returns nil when team not found", func(t *testing.T) {
			match, err := d.GetLatestMatch(c.TestID)

			if c.IsNil && ((match != nil) || err == nil) {
				t.Errorf("Expected match to be Nil, got ID %d instead", match.ID)
			}
		})
	}

	// Teardown de la BD de test
	if config.DatabaseDriver == "sqlite3" {
		os.Remove(config.ConnectionString)
	}
}
