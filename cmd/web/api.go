package main

import (
	"context"
	"expvar"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/bri_otel"
	"free-ent-guide-backend/pkg/cred"
	"log"
	"os"
	"os/signal"

	"net/http"
	"time"

	_ "github.com/shaj13/libcache/fifo"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
)

var (
	c        cred.Cred
	counters *expvar.Map
	Queries  *modelstore.Queries
	client   *http.Client
)

const (
	instrumentationName    = "github.com/briwagner/free-entertainment-guide-backend"
	instrumentationVersion = "0.1.0"
	appName                = "free-entertainment-guide"
)

func main() {
	// Set-up application config.
	c.GetCreds("creds", ".")

	// Set-up database.
	Queries = modelstore.New(models.Setup(c))
	client = &http.Client{Timeout: time.Second * 5} // TODO why is this not on the app?

	// Set-up authentication.
	var app App
	app.setupGoGuardian()
	app.setupClient(5)

	// Setup OTel for logging, tracing (TODO metrics?)
	ctx := context.Background()
	lp, tp := bri_otel.SetupOtel(ctx, instrumentationName, instrumentationVersion, appName)
	if lp == nil || tp == nil {
		panic("failed to setup open telemetry")
	}
	defer func() {
		// TODO combine these shutdown funcs
		if err := lp.Shutdown(ctx); err != nil {
			log.Printf("shutdown logger error %v", err)
		}
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("shutdown tracer error %v", err)
		}
	}()
	global.SetLoggerProvider(lp)
	app.l = otelslog.NewLogger(appName)
	app.t = otel.Tracer(appName)

	// Create router
	mux := NewRouter(app)
	counters = expvar.NewMap("counters")
	counters.Set("cache_hit", new(expvar.Int))
	counters.Set("cache_miss", new(expvar.Int))

	fmt.Printf("ENT API is live. Listening on port %v ...\n", c.GetPort())
	app.l.Info("app starting up")
	srv := &http.Server{
		Addr:         "0.0.0.0" + c.GetPort(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// do something before shutdown.
	// e.g. send slack message. Need to move cli_slack to pkg/

	srv.Shutdown(ctx)
	app.l.Info("shutting down")
	os.Exit(0)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if c.Env != "prod" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		return
	}

	(*w).Header().Set("Access-Control-Allow-Origin", c.CorsOrigin)
	(*w).Header().Set("Vary", "Origin")
}

func cacheStatus(w *http.ResponseWriter, status bool) {
	if status {
		(*w).Header().Set("Cache", "HIT")
		counters.Add("cache_hit", 1)
	} else {
		(*w).Header().Set("Cache", "MISS")
		counters.Add("cache_miss", 1)
	}
}
