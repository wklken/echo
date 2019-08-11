package server

import (
	"net/http/pprof"

	"github.com/go-chi/chi"
)

func RegisterAPIs(r *chi.Mux) {
	r.HandleFunc("/ping/", Pong)

	// sleep for X ms
	r.HandleFunc("/sleep/{sleep}/", Sleep)

	// status for X
	r.HandleFunc("/status/{status}/", Status)

	// get/delete/post/patch...
	r.HandleFunc("/echo/", Echo)

	// file download and upload
	r.HandleFunc("/file/download/{size}/", FileDownload)
	r.HandleFunc("/file/upload/{filename}/", FileUpload)

	// websocket
	r.HandleFunc("/ws/index/", WebsocketIndex)
	r.HandleFunc("/ws/", Websocket)

	// pprof
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/profile/", pprof.Profile)
	r.HandleFunc("/debug/pprof/cmdline/", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/symbol/", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace/", pprof.Trace)
	r.Handle("/debug/pprof/block/", pprof.Handler("block"))
	r.Handle("/debug/pprof/goroutine/", pprof.Handler("goroutine"))
	r.Handle("/debug/pprof/heap/", pprof.Handler("heap"))
	r.Handle("/debug/pprof/threadcreate/", pprof.Handler("threadcreate"))

	// ! DONE ================================================
	// TODO: finish all the handler one by one

	// httpbin api
	// TODO
	r.HandleFunc("/legacy", pong)
	// TODO
	r.HandleFunc("/html", pong)

	// TODO: app.route default method is?

	r.HandleFunc("/robots.txt", robots)
	r.HandleFunc("/deny", deny)
	r.HandleFunc("/ip", ip)
	r.HandleFunc("/uuid", uuidResponse)
	r.HandleFunc("/headers", headers)
	r.HandleFunc("/user-agent", userAgent)
	r.Get("/get", get)
	r.Post("/post", post)
	r.Put("/put", put)
	r.Patch("/patch", patch)
	r.Delete("/delete", delete)
	r.HandleFunc("/anything", anything)
	// TODO: here is path:anything
	r.HandleFunc("/anything/{anything}", anything)

	// TODO
	r.HandleFunc("/gzip", gzip)
	r.HandleFunc("/deflate", deflate)
	r.HandleFunc("/brotli", brotli)
	r.HandleFunc("/redirect/<int:n>", redirectNTimes)
	r.HandleFunc("/redirect-to", redirectTo)
	r.HandleFunc("/relative-redirect/<int:n>", relativeRedirectNTimes)
	r.HandleFunc("/absolute-redirect/<int:n>", absoluteRedirectNTimes)
	r.HandleFunc("/stream/<int:n>", streamNMessage)

	// TODO: to support weight
	r.HandleFunc("/status/<codes>", statusCode)
	r.HandleFunc("/response-headers", responseHeaders)

	r.HandleFunc("/cookies", cookies)
	// ! DONE ================================================

	r.HandleFunc("/forms/post", formsPost)
	r.HandleFunc("/cookies/set/<name>/<value>", setCookie)
	r.HandleFunc("/cookies/set", setCookies)
	r.HandleFunc("/cookies/delete", deleteCookies)
	r.HandleFunc("/basic-auth/<user>/<passwd>", basicAuth)
	r.HandleFunc("/hidden-basic-auth/<user>/<passwd>", hiddenBasicAuth)
	r.HandleFunc("/bearer", bearerAuth)
	r.HandleFunc("/digest-auth/<qop>/<user>/<passwd>", digestAuthMd5)
	r.HandleFunc("/digest-auth/<qop>/<user>/<passwd>/<algorithm>", digestAuthNostale)
	r.HandleFunc("/digest-auth/<qop>/<user>/<passwd>/<algorithm>/<stale_after>", digestAuth)
	r.HandleFunc("/delay/<delay>", delayResponse)
	r.HandleFunc("/drip", drip)
	r.HandleFunc("/base64/<value>", decodeBase64)
	r.HandleFunc("/cache", cache)
	r.HandleFunc("/etag/<etag>", etag)
	r.HandleFunc("/cache/<int:value>", cacheControl)
	r.HandleFunc("/encoding/utf8", encoding)
	r.HandleFunc("/bytes/<int:n>", randomBytes)
	r.HandleFunc("/stream-bytes/<int:n>", streamBytes)
	r.HandleFunc("/range/<int:numbytes>", rangeRequest)
	r.HandleFunc("/links/<int:n>/<int:offset>", linkPage)
	r.HandleFunc("/links/<int:nou>", links)
	r.HandleFunc("/image", image)
	r.HandleFunc("/image/png", imagePng)
	r.HandleFunc("/image/jpeg", imageJpeg)
	r.HandleFunc("/image/webp", imageWebp)
	r.HandleFunc("/image/svg", imageSvg)
	r.HandleFunc("/xml", pong)
	r.HandleFunc("/json", json)

}
