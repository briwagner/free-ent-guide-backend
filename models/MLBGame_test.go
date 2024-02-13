package models

import (
	"encoding/json"
	"free-ent-guide-backend/pkg/cred"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	c := cred.Cred{DB: "gorm:password@tcp(127.0.0.1:3306)/demo_gorm_test?charset=utf8mb4&parseTime=True&loc=Local"}
	AppStore, _ = Setup(c)
}

func TestMarshal_MLBGame(t *testing.T) {
	liveGame, err := os.ReadFile("testdata/mlbgame.json")
	require.NoError(t, err)

	gu := MLBGameUpdate{}
	err = json.Unmarshal(liveGame, &gu)
	if err != nil {
		panic(err)
	}

	if gu.ID != 715721 {
		t.Errorf("Game ID. Got %d, want %d", gu.ID, 715721)
	}
	if gu.Status != "In Progress" {
		t.Errorf("Game status. Got %s, want %s", gu.Status, "In Progress")
	}
	if gu.Inning != 5 {
		t.Errorf("Game inning. Got %d, want %d", gu.Inning, 9)
	}
	if gu.HomeScore != 0 {
		t.Errorf("Home score. Got %d, want %d", gu.HomeScore, 0)
	}
	if gu.VisitorScore != 5 {
		t.Errorf("Visitor score. Got %d, want %d", gu.VisitorScore, 5)
	}

	oldGame, err := os.ReadFile("testdata/mlbgame_old.json")
	require.NoError(t, err)

	gu2 := MLBGameUpdate{}
	err = json.Unmarshal(oldGame, &gu2)
	if err != nil {
		panic(err)
	}

	if gu2.ID != 715720 {
		t.Errorf("Game ID. Got %d, want %d", gu2.ID, 715720)
	}
	if gu2.Status != "Scheduled" {
		t.Errorf("Game status. Got %s, want %s", gu2.Status, "Scheduled")
	}
	if gu2.Inning != 0 {
		t.Errorf("Game inning. Got %d, want %d", gu2.Inning, 0)
	}
	if gu2.HomeScore != 0 {
		t.Errorf("Home score. Got %d, want %d", gu2.HomeScore, 0)
	}
	if gu2.VisitorScore != 0 {
		t.Errorf("Visitor score. Got %d, want %d", gu2.VisitorScore, 0)
	}
}
