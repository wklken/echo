package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	r "github.com/unrolled/render"
)

var render *r.Render

func init() {
	render = r.New() // pass options if you want
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func sleep(w http.ResponseWriter, r *http.Request) {
	sleepTime := chi.URLParam(r, "sleep")
	s, err := strconv.Atoi(sleepTime)

	if err == nil && s > 0 {
		time.Sleep(time.Duration(s) * time.Millisecond)
	} else {
		runtime.Gosched()
	}

	w.Write([]byte("ok"))
}

func status(w http.ResponseWriter, r *http.Request) {
	status := chi.URLParam(r, "status")
	statusCode, err := strconv.Atoi(status)

	if err != nil {
		http.Error(w, "wrong status", 500)
		return
	}

	http.Error(w, http.StatusText(statusCode), statusCode)
}

func echo(w http.ResponseWriter, r *http.Request) {
	// read body first, will the parse form will drain the body
	body := ""
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		body = string(buf)
	}

	// get the form
	r.ParseForm()
	form := r.PostForm

	data := map[string]interface{}{
		"url":     r.URL.Path,
		"method":  r.Method,
		"headers": r.Header,
		"query":   r.URL.Query(),
		"form":    form,
		"body":    body,
	}

	render.JSON(w, 200, data)
}

func registerAPIs(r *chi.Mux) {
	r.HandleFunc("/ping/", pong)

	// sleep for X ms
	r.HandleFunc("/sleep/{sleep}/", sleep)

	// status for X
	r.HandleFunc("/status/{status}/", status)

	// get/delete/post/patch...
	r.HandleFunc("/echo/", echo)
}

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
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	registerAPIs(r)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("listen on %s\n", addr)
	http.ListenAndServe(addr, r)
}
