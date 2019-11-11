package util

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	stdlog "log"
)

// RunHTTPServer takes the argument http server and runs it in a goroutine.
// Listens to system signals (eg. SIGKILL) to gracefully shutdown the server.
// This method never returns; it calls os.Exit() after the graceful server
// shutdown.
func RunHTTPServer(server *http.Server) {
	go func() {
		// Start serving
		err := server.ListenAndServe()
		if err != nil {
			stdlog.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	// Listen to SIGINT, SIGKILL, SIGQUIT, SIGTERM and gracefully shut down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()

	server.Shutdown(ctx)
	os.Exit(0)
}
