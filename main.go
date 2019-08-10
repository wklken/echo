package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/wklken/echo/server"
)

func main() {
	// listen port required
	args := os.Args
	if len(args) != 2 {
		fmt.Println("the listen port required")
		os.Exit(1)
	}
	port, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("the listen port should be integer")
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	server.RegisterAPIs(r)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("listen on %s\n", addr)
	http.ListenAndServe(addr, r)
}
