// +build !appengine

package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/arunjit/gonotes/notes"
)

// Flags
var (
	hostFlag  = flag.String("host", "127.0.0.1", "Host to run the server on.")
	portFlag  = flag.Int("port", 1601, "Port to run the server on.")
	fcgiFlag  = flag.Bool("fcgi", false, "Use FCGI instead.")
	redisFlag = flag.String("redis", "127.0.0.1:6379", "Redis host:port.")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *hostFlag, *portFlag)

	db := notes.NewRedisDatabase(*redisFlag)
	svc := notes.NewNotesService(db)
	svr := notes.NewServer(svc)

	if *fcgiFlag {
		l, _ := net.Listen("tcp", addr)
		fcgi.Serve(l, svr)
	} else {
		http.Handle(notes.RPCPath, svr)
		http.ListenAndServe(addr, nil)
	}
}
