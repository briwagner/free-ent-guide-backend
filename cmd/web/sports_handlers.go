package main

import (
	"encoding/json"
	"errors"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/pkg/nhlapi"
	"log"
	"net/http"
	"strconv"
	"time"
)

// NHLGameHandler returns a single scheduled game by GameID,
// and includes live scores and timing from the official api.
func NHLGameHandler(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	// Fetch game from database.
	var g models.NHLGame
	gid, err := strconv.Atoi(common.vars["game_id"])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = g.FindByGameID(common.queries, gid)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Fetch scores, timing from the NHL api.
	gu, err := nhlapi.GetUpdate(g.GameID)
	if err != nil {
		log.Print(err)
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
		log.Print(err)
	}
}

// NHLGamesHandler responds to GET and returns the
// schedule of NHL games for the specified date.
func NHLGamesHandler(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	if common.queryDate == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Must pass a date"))
		return
	}

	gs := &models.NHLGames{}
	err := gs.LoadByDate(common.queries, common.queryDate)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if len(*gs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = json.NewEncoder(w).Encode(gs)
	if err != nil {
		log.Print(gs)
	}
}

func NHLGamesLatest(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	// TODO why is this a two-step process?
	gs, err := models.NHLGetLatestGames(common.queries)
	if err != nil || len(gs) == 0 {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var games models.NHLGames
	err = games.LoadByDate(common.queries, gs[0].Gametime.Format("2006-01-02"))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		log.Print(err)
	}
}

func NHLGamesNext(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	games, err := models.NHLGetNextGameday(common.queries, common.now)
	if err != nil {
		log.Printf("error getting next NHL games for %s: %s", common.now.Format("2006-01-02"), err)
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		log.Print(err)
	}
}

// MLBGameHandler responds to GET and returns a single
// game matching the specified GameID.
func MLBGameHandler(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	// Fetch game from database.
	gameID, err := strconv.Atoi(common.vars["game_id"])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	g, err := models.FindByGameID(common.queries, gameID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Fetch scores, timing from the MLB api.
	err = g.GetUpdate()
	if err != nil {
		log.Print(err)
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
		log.Print(err)
	}
}

// MLBGamesHandler responds to GET and returns the
// schedule of MLB games for the specified date.
func MLBGamesHandler(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	if common.queryDate == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Must pass a date"))
		return
	}

	games := &models.MLBGames{}
	err := games.LoadByDate(common.queries, common.queryDate)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		log.Print(err)
	}
}

func MLBTeamHandler(w http.ResponseWriter, r *http.Request) {
	common := prepareResponse(w, r)

	teamID, err := strconv.Atoi(common.vars["team_id"])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	team := models.MLBTeam{TeamID: teamID}
	// Use midnight to get any games this day.
	games, err := team.GamesByTeam(r.Context(), common.queries, common.now.UTC().Truncate(24*time.Hour))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		log.Print(err)
	}
}
