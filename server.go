package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	r "github.com/unrolled/render"
	"gopkg.in/olahol/melody.v1"
)

var render *r.Render
var m = melody.New()

func init() {
	render = r.New() // pass options if you want

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
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

func fileDownload(w http.ResponseWriter, r *http.Request) {
	// KB
	sizeStr := chi.URLParam(r, "size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}

	buf := make([]byte, size*1024)

	for i := 0; i < len(buf); i++ {
		buf[i] = '0'
	}

	file := "data.txt"

	// set the default MIME type to send
	mime := http.DetectContentType(buf)
	fileSize := len(string(buf))

	// Generate the server headers
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment; filename="+file+"")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	http.ServeContent(w, r, file, time.Now(), bytes.NewReader(buf))
}

func fileUpload(w http.ResponseWriter, r *http.Request) {

	fileName := chi.URLParam(r, "filename")

	r.ParseMultipartForm(32 << 20)
	// file, handler, err := r.FormFile(fileName)
	file, _, err := r.FormFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	io.Copy(ioutil.Discard, file)
	w.WriteHeader(204)
}

func websocket(w http.ResponseWriter, r *http.Request) {
	m.HandleRequest(w, r)
}

func websocketIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")

}

func registerAPIs(r *chi.Mux) {
	r.HandleFunc("/ping/", pong)

	// sleep for X ms
	r.HandleFunc("/sleep/{sleep}/", sleep)

	// status for X
	r.HandleFunc("/status/{status}/", status)

	// get/delete/post/patch...
	r.HandleFunc("/echo/", echo)

	// file download and upload
	r.HandleFunc("/file/download/{size}/", fileDownload)
	r.HandleFunc("/file/upload/{filename}/", fileUpload)

	r.HandleFunc("/ws/index/", websocketIndex)
	r.HandleFunc("/ws/", websocket)
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
