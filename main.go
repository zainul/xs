package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zainul/xs/internal/gateways/cache/redis"

	"github.com/gorilla/mux"
	HTTPHandler "github.com/zainul/xs/internal/delivery/http"
	"github.com/zainul/xs/internal/gateways/database/postgres"
	"github.com/zainul/xs/internal/pkg/initial"
	userRepository "github.com/zainul/xs/internal/repository/user"
	"github.com/zainul/xs/internal/usecase/user"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait,
		"graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	db := initial.GetDB()
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	////////////////////////////////////////////////////////////////////////////
	////////////////////////////// APPLICATION /////////////////////////////////
	////////////////////////////////////////////////////////////////////////////

	userDBStore := postgres.NewUserStore(db)
	userCached := redis.NewCached("someconfig")
	userCounter := redis.NewCounter("someconfig")

	userRepo := userRepository.NewUserRepository(userDBStore, userCached, userCounter)
	useCaseUser := user.NewUserUseCase(userRepo)
	HTTPHandler.NewUserHandler(r, useCaseUser)

	////////////////////////////////////////////////////////////////////////////

	http.Handle("/", r)

	// config the http server
	srv := &http.Server{
		Addr: "0.0.0.0:7000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		// WriteTimeout: time.Second * 1,
		ReadTimeout: time.Second * 15, // -> it will be in config file
		// IdleTimeout: time.Second * 1,
		Handler: r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Server up :7000")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
