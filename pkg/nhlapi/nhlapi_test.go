package nhlapi

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestMarshal_NHLGameUpdate(t *testing.T) {
	nhlgame, err := os.ReadFile("testdata/nhlgame_update.json")
	require.NoError(t, err)

	gu := NHLGameUpdate{}
	err = json.Unmarshal(nhlgame, &gu)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 2023020672, int(gu.ID))
	assert.Equal(t, "Final", gu.Status)
	assert.Equal(t, 3, int(gu.Period))
	assert.Equal(t, 3, int(gu.HomeScore))
	assert.Equal(t, 0, int(gu.VisitorScore))
}

func Test_GetTeam(t *testing.T) {
	data := `{"players":[],"teams":[{"name":{"default":"Utah Hockey Club","fr":"Club de hockey de l'Utah"},"tricode":"UTA","teamId":59}],"seasonStates":[]}`
	tpl := GetTeamPayload{}
	err := json.Unmarshal([]byte(data), &tpl)
	require.NoError(t, err)

	nhlteam := tpl.ToTeam()
	assert.Equal(t, "Utah Hockey Club", nhlteam.Name)
	assert.Equal(t, "UTA", nhlteam.Tricode)
	assert.Equal(t, 59, nhlteam.ID)

	t.Skip("http request")
	team, err := GetTeam(t.Context(), &http.Client{}, "59")
	require.NoError(t, err)
	assert.Equal(t, "Utah Hockey Club", team.Name)
	assert.Equal(t, "UTA", team.Tricode)
	assert.Equal(t, 59, team.ID)
}
