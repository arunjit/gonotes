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
	hostFlag = flag.String("host", "127.0.0.1", "Host to run the server on.")
	portFlag = flag.Int("port", 1601, "Port to run the server on.")
	fcgiFlag = flag.Bool("fcgi", false, "Use FCGI instead.")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *hostFlag, *portFlag)
	server := notes.NewServer(new(notes.RedisNotesService))

	if *fcgiFlag {
		l, _ := net.Listen("tcp", addr)
		fcgi.Serve(l, server)
	} else {
		http.Handle("/rpc", server)
		http.ListenAndServe(addr, nil)
	}
}
