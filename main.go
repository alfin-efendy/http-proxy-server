package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
)

func main() {
	verbose := flag.Bool("verbose", false, "should every proxy request be logged to stdout")
	port := flag.String("port", "3128", "proxy listen address")
	username := flag.String("user", "admin", "username for proxy authentication")
	password := flag.String("pass", "password", "password for proxy authentication")

	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	auth.ProxyBasic(proxy, "master", func(user, pass string) bool {
		return user == *username && pass == *password
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), proxy))
}
