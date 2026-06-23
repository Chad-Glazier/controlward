package httpserver

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/Chad-Glazier/controlward/httpserver/handler"
	"github.com/Chad-Glazier/controlward/httpserver/middleware"
	"github.com/Chad-Glazier/controlward/riot"
	"github.com/Chad-Glazier/controlward/store/memstore"
)

const Domain = "localhost"

var Port string

func init() {
	Port = os.Getenv("CONTROLWARD_PORT")
	if Port == "" {
		panic("missing CONTROLWARD_PORT environment variable")
	}
	if !ValidPort(Port) {
		panic("environment variable CONTROLWARD_PORT is not a valid port number")
	}
}

// Starts the HTTP server.
func Start() {

	//
	// Set the configuration.
	//

	cache := riot.Cache{
		Matches: memstore.NewMatchStore(2048),
	}
	riotClient := riot.NewClientWithCache(cache)
	riotClient.Logger = slog.Default()

	conf := &handler.Config{
		Riot:   riot.NewClientWithCache(cache),
		Logger: slog.Default(),
	}

	//
	// Register the handlers.
	//

	mux := http.NewServeMux()

	mux.HandleFunc("GET /openapi.yaml", handler.OpenAPISpec)
	mux.HandleFunc("GET /", handler.DocsPage)
	mux.HandleFunc("GET /health", handler.Health(conf))

	mux.HandleFunc("GET /history/{gameName}/{tagLine}", handler.GetHistory(conf))

	//
	// Register global middleware.
	//

	h := middleware.Logger(mux)
	h = middleware.Cors(h)

	//
	// Configure the server.
	//

	server := &http.Server{
		Addr:    ":" + Port,
		Handler: h,
	}

	slog.Info("starting server at http://" + Domain + ":" + Port)

	err := server.ListenAndServe()
	if err != nil {
		slog.Error("server failed to start")
	}
}

// Returns true if and only if the given string represents a valid port number.
func ValidPort(port string) bool {
	portNum, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return false
	}
	if portNum < 1 || portNum > 65535 {
		return false
	}
	return true
}
