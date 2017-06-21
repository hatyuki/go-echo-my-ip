package askmyip

import (
	"fmt"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/", requestHandler)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, GetRemoteAddr(r))
}

func GetRemoteAddr(r *http.Request) string {
	var ip string

	if ip = getOriginalIP(r); ip == "" {
		ip = r.RemoteAddr
	}

	return stripPort(ip)
}

func getOriginalIP(r *http.Request) string {
	fmt.Printf("%#v\n", r.Header.Get("X-Forwarded-For"))
	ips := strings.Split(r.Header.Get("X-Forwarded-For"), ",")
	return strings.TrimSpace(ips[0])
}

func stripPort(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return hostport
	}
	if i := strings.IndexByte(hostport, ']'); i != -1 {
		return strings.TrimPrefix(hostport[:i], "[")
	}
	return hostport[:colon]
}
