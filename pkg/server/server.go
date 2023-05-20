package server

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/delivery/graphql/generated"
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/delivery/graphql/resolvers"
	"github.com/shashaneRanasinghe/simpleMongoAPI/pkg/database"
	"github.com/tryfix/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func Serve() chan string {
	ctx := context.Background()

	server := http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      nil,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	dbClient := database.IntiDB()
	resolver := resolvers.NewResolver(dbClient)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.Handle("/metrics", promhttp.Handler())

	closeChannel := make(chan string)

	//This function will ensure the server to shut down gracefully
	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)
		signal.Notify(sig, syscall.SIGTERM)
		signal.Notify(sig, syscall.SIGQUIT)
		<-sig

		log.Info("service Interruption Signal Received")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		database.DisconnectDB(dbClient)

		err := server.Shutdown(ctx)
		if err != nil {
			log.Error("Server shutdown Error ", err)
		}

		log.Info("GraphQL Server Stopped")
		close(closeChannel)
	}()

	log.Info("connect to http://localhost" + server.Addr + "/ for GraphQL playground")
	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}

	return closeChannel
}
