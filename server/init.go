package server

import (
	r "github.com/unrolled/render"
	"gopkg.in/olahol/melody.v1"
)

var render *r.Render
var m = melody.New()

func init() {
	render = r.New(r.Options{
		Directory: "templates",
	}) // pass options if you want

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
}
