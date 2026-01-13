package main

import (
	"encoding/json"
	"errors"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/pkg/nhlapi"
	"net/http"
	"strconv"
	"time"
)

// NHLGameHandler returns a single scheduled game by GameID,
// and includes live scores and timing from the official api.
func (app *App) NHLGameHandler(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	// Fetch game from database.
	var g models.NHLGame
	gid, err := strconv.Atoi(common.vars["game_id"])
	if err != nil {
		app.l.Error("error nhlGameHandler", "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = g.FindByGameID(common.queries, gid)
	if err != nil {
		app.l.Error("error nhlGameHandler", "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Fetch scores, timing from the NHL api.
	gu, err := nhlapi.GetUpdate(r.Context(), client, g.GameID)
	if err != nil {
		app.l.Error("error nhlGameHandler", "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// TODO streamline this so we don't need a custom struct to return?
	g.HomeScore = int(gu.HomeScore)
	g.VisitorScore = int(gu.VisitorScore)
	g.Period = int(gu.Period)
	g.Status = gu.Status
	ngu := &models.NHLGameUpdate{}
	ngu.FromGame(g)

	err = json.NewEncoder(w).Encode(ngu)
	if err != nil {
		app.l.Error("error nhlGameHandler", "error", err)
	}
}

// NHLGamesHandler responds to GET and returns the
// schedule of NHL games for the specified date.
func (app *App) NHLGamesHandler(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	if common.queryDate == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Must pass a date"))
		return
	}

	gs := &models.NHLGames{}
	err := gs.LoadByDate(common.queries, common.queryDate)
	if err != nil {
		app.l.Error("error nhlGamesHandler", "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if len(*gs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = json.NewEncoder(w).Encode(gs)
	if err != nil {
		app.l.Error("error nhlGamesHandler", "error", gs)
	}
}

func (app *App) NHLGamesLatest(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	// TODO why is this a two-step process?
	gs, err := models.NHLGetLatestGames(r.Context(), common.queries)
	if err != nil || len(gs) == 0 {
		app.l.Error("error nhlGamesLatest", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var games models.NHLGames
	err = games.LoadByDate(common.queries, gs[0].Gametime.Format("2006-01-02"))
	if err != nil {
		app.l.Error("error nhlGamesLatest", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		app.l.Error("error nhlGamesLatest", "error", err)
	}
}

func (app *App) NHLGamesNext(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	games, err := models.NHLGetNextGameday(common.queries, common.now)
	if err != nil {
		app.l.Error("error getting next NHL games", "time", common.now.Format("2006-01-02"), "error", err)
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		app.l.Error("error nhlGamesNext", "error", err)
	}
}

// MLBGameHandler responds to GET and returns a single
// game matching the specified GameID.
func (app *App) MLBGameHandler(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	// Fetch game from database.
	gameID, err := strconv.Atoi(common.vars["game_id"])
	if err != nil {
		app.l.Error("error mlbGameHandler", "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	g, err := models.FindByGameID(r.Context(), common.queries, gameID)
	if err != nil {
		app.l.Error("error mlbGameHandler", "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Fetch scores, timing from the MLB api.
	err = g.GetUpdate(client)
	if err != nil {
		app.l.Error("error mlbGameHandler", "error", err)
		// Special error case for canceled game that should be marked deleted in DB.
		if errors.Is(err, models.ErrorGameCanceled) {
			w.WriteHeader(http.StatusGone)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// TODO we are returning the Game, not a GameUpdate anymore.
	err = json.NewEncoder(w).Encode(g)
	if err != nil {
		app.l.Error("error mlbGameHandler", "error", err)
	}
}

// MLBGamesHandler responds to GET and returns the
// schedule of MLB games for the specified date.
func (app *App) MLBGamesHandler(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	if common.queryDate == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Must pass a date"))
		return
	}

	games := &models.MLBGames{}
	err := games.LoadByDate(r.Context(), common.queries, common.queryDate)
	if err != nil {
		app.l.Error("error mlbGamesHandler", "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		app.l.Error("error mlbGamesHandler", "error", err)
	}
}

func (app *App) MLBTeamHandler(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	teamID, err := strconv.Atoi(common.vars["team_id"])
	if err != nil {
		app.l.Error("error mlbTeamHandler", "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	team := models.MLBTeam{TeamID: teamID}
	err = team.FindByTeamID(r.Context(), common.queries, teamID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Use midnight to get any teamData this day.
	teamData, err := team.GamesByTeam(r.Context(), common.queries, common.now.UTC().Truncate(24*time.Hour))
	if err != nil {
		app.l.Error("error mlbTeamHandler", "error", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	standing, err := teamData.Team.GetStandings(r.Context(), app.client)
	// this does't work in the off-season
	if len(teamData.PastGames) > 0 {
		if err != nil {
			app.l.Warn("error mlbTeamHandler", "error", err, "team name", teamData.Team.Name)
		}
	}
	teamData.Standings = standing

	err = json.NewEncoder(w).Encode(teamData)
	if err != nil {
		app.l.Error("error mlbTeamHandler", "error", err)
	}
}
