package nhlapi

import (
	"encoding/json"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/cred"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	Queries *modelstore.Queries
)

func init() {
	c := cred.Cred{DB: "ent_user:ent_password@tcp(127.0.0.1:3306)/ent_guide_test?charset=utf8mb4&parseTime=True&loc=Local"}
	_, st := models.Setup(c)

	Queries = modelstore.New(st)
}

func TestUnmarshal_Schedule(t *testing.T) {
	data, err := os.ReadFile("testdata/schedule.json")
	require.NoError(t, err)

	var payload GameWeekPayload
	err = json.Unmarshal(data, &payload)
	assert.NoError(t, err)

	assert.Equal(t, 7, len(payload.Days))
	assert.Equal(t, 11, len(payload.Days[0].Games))

	g1 := payload.Days[0].Games[0]
	assert.Equal(t, 2, g1.Away.ID, "away team ID")
	assert.Equal(t, 6, g1.Home.ID, "home team ID")
	assert.Equal(t, 2023020193, g1.ID, "game ID")
	assert.Equal(t, "2023-11-10 00:00:00 +0000 UTC", g1.Gametime.String(), "gametime in UTC")
	assert.Equal(t, "Regular Season", g1.GameType.Type, "maps gameType")
	assert.Equal(t, "FUT", g1.GameState, "maps gamestate")
}

func TestUnmarshal_Teams(t *testing.T) {
	data, err := os.ReadFile("testdata/nhl_teams.json")
	require.NoError(t, err)

	var payload GetTeamsPayload
	err = json.Unmarshal(data, &payload)
	require.NoError(t, err)

	assert.Len(t, payload.Data, 59, "NHL api has 59 teams")
	assert.Equal(t, "Atlanta Thrashers", payload.Data[0].Name, "team name is mapped")
	assert.Equal(t, "ATL", payload.Data[0].Tricode, "team tricode is mapped")
	assert.Equal(t, 35, payload.Data[0].FranchiseID, "team franchiseID is mapped")
	assert.Equal(t, 11, payload.Data[0].ID, "team ID is mapped")
}
