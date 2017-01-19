package common_test

import (
	"os"
	"testing"

	"github.com/tsap-laval/tsap-common/common"
)

type testCase struct {
	TestID   uint
	IsNil    bool
	ExpectID uint
}

func TestDatasource(t *testing.T) {

	config, err := common.GetConfig()

	if err != nil {
		t.Errorf("Error loading configuration: %s", err.Error())
	}

	// On seed une base de données de test
	err = common.SeedData(config.DatabaseDriver, config.ConnectionString, config.SeedDataPath)

	if err != nil {
		t.Errorf("Unexpected exception: %s", err.Error())
	}

	d := common.NewDatasource(config.DatabaseDriver, config.ConnectionString)

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

	// Teardown
	os.Remove("test.db")
}
