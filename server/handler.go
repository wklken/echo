package server

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// Pong will return 200 with body: pong
func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// Sleep will sleep for x ms before response
func Sleep(w http.ResponseWriter, r *http.Request) {
	sleepTime := chi.URLParam(r, "sleep")
	s, err := strconv.Atoi(sleepTime)

	if err == nil && s > 0 {
		time.Sleep(time.Duration(s) * time.Millisecond)
	} else {
		runtime.Gosched()
	}

	w.Write([]byte("ok"))
}

// Status will return the specific response status
func Status(w http.ResponseWriter, r *http.Request) {
	status := chi.URLParam(r, "status")
	statusCode, err := strconv.Atoi(status)

	if err != nil {
		http.Error(w, "wrong status", 500)
		return
	}

	http.Error(w, http.StatusText(statusCode), statusCode)
}

// Echo will return 200 withe json body contains all the data of the request
func Echo(w http.ResponseWriter, r *http.Request) {
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

// FileDownload will down the specific size(in KB) file
func FileDownload(w http.ResponseWriter, r *http.Request) {
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

// FileUpload will receive the file upload request
func FileUpload(w http.ResponseWriter, r *http.Request) {

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

// WebsocketIndex will render the websocket index page
func WebsocketIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// Websocket will handle the ws requests
func Websocket(w http.ResponseWriter, r *http.Request) {
	m.HandleRequest(w, r)
}

// ============================
// httpbin handlers begin

const (
	RobotText = `User-agent: *
	Disallow: /deny
	`

	DenyText = `
	.-''''''-.
	.' _      _ '.
   /   O      O   \\
  :                :
  |                |
  :       __       :
   \  .-"` + "`" + "  " + "`" + `"-.  /
	'.          .'
	  '-......-'
 YOU SHOULDN'T BE HERE `
)

// ! DONE
func robots(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(RobotText))
	w.Header().Set("Content-Type", "text/plain")
}
func deny(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(DenyText))
	w.Header().Set("Content-Type", "text/plain")
}

func ip(w http.ResponseWriter, r *http.Request) {

	i := r.Header.Get("X-Forwarded-For")
	if i == "" {
		i = r.RemoteAddr
	}

	data := map[string]interface{}{
		"origin": i,
	}

	render.JSON(w, 200, data)
}

func uuidResponse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")

	data := map[string]interface{}{
		"uuid": uuid.Must(uuid.NewV4()).String(),
	}

	render.JSON(w, 200, data)
}

func headers(w http.ResponseWriter, r *http.Request) {
	data := GetDict(r, nil, "headers")
	render.JSON(w, 200, data)
}

func userAgent(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"user-agent": r.Header.Get("User-Agent"),
	}
	render.JSON(w, 200, data)
}

func get(w http.ResponseWriter, r *http.Request) {
	items := []string{
		"url", "args", "headers", "origin",
	}
	data := GetDict(r, nil, items...)
	render.JSON(w, 200, data)
}

func post(w http.ResponseWriter, r *http.Request) {
	items := []string{
		"url", "args", "form", "data", "origin", "headers", "files", "json",
	}
	data := GetDict(r, nil, items...)
	render.JSON(w, 200, data)
}

func put(w http.ResponseWriter, r *http.Request) {
	items := []string{
		"url", "args", "form", "data", "origin", "headers", "files", "json",
	}
	data := GetDict(r, nil, items...)
	render.JSON(w, 200, data)
}

func patch(w http.ResponseWriter, r *http.Request) {
	items := []string{
		"url", "args", "form", "data", "origin", "headers", "files", "json",
	}
	data := GetDict(r, nil, items...)
	render.JSON(w, 200, data)
}

func delete(w http.ResponseWriter, r *http.Request) {
	items := []string{
		"url", "args", "form", "data", "origin", "headers", "files", "json",
	}
	data := GetDict(r, nil, items...)
	render.JSON(w, 200, data)
}

func anything(w http.ResponseWriter, r *http.Request) {
	items := []string{
		"url", "args", "form", "data", "origin", "headers", "files", "json", "method",
	}
	data := GetDict(r, nil, items...)
	render.JSON(w, 200, data)
}

// ! ====================================

func gzip(w http.ResponseWriter, r *http.Request) {

	items := []string{
		"origin", "headers",
	}
	extras := map[string]string{
		"method": r.Method,
	}
	data := GetDict(r, extras, items...)

	render.JSON(w, 200, data)
}

func deflate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func brotli(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func redirectNTimes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func redirectTo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func relativeRedirectNTimes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func absoluteRedirectNTimes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func streamNMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func statusCode(w http.ResponseWriter, r *http.Request) {
	status := chi.URLParam(r, "status")
	statusCode, err := strconv.Atoi(status)

	if err != nil {
		http.Error(w, "wrong status", 500)
		return
	}

	http.Error(w, http.StatusText(statusCode), statusCode)
}

func responseHeaders(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func cookies(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{}

	for _, cookie := range r.Cookies() {
		data[cookie.Name] = cookie.Value
	}
	render.JSON(w, 200, data)
}

func formsPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func setCookies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func deleteCookies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func basicAuth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func hiddenBasicAuth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func bearerAuth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func digestAuthMd5(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func digestAuthNostale(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func digestAuth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func delayResponse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func drip(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func decodeBase64(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func cache(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func etag(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func cacheControl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func encoding(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func randomBytes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func streamBytes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func rangeRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func linkPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func links(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func image(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}
func imagePng(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func imageJpeg(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func imageWebp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func imageSvg(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.Header().Set("Content-Type", "text/plain")
}

func xml(w http.ResponseWriter, r *http.Request) {

	render.XML(w, 200, "")
	// w.Write([]byte(""))
	// w.Header().Set("Content-Type", "text/plain")
}

func json(w http.ResponseWriter, r *http.Request) {
	slideShow := map[string]interface{}{
		"title":  "Sample Slide Show",
		"date":   "date of publication",
		"author": "Yours Truly",
		"slides": []interface{}{
			map[string]string{
				"type":  "all",
				"title": "Wake up to WonderWidgets!",
			},
			map[string]interface{}{
				"type":  "all",
				"title": "Overview",
				"items": []string{
					"Why <em>WonderWidgets</em> are great",
					"Who <em>buys</em> WonderWidgets",
				},
			},
		},
	}
	render.JSON(w, 200, slideShow)
}
