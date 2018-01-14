// +build !solution

package main

import (
	"flag"
	"log"
	"math"
	"net/http"
	"strconv"
)

var httpAddr = flag.String("http", ":8080", "Listen address")

func main() {
	flag.Parse()
	server := NewServer()
	log.Fatal(http.ListenAndServe(*httpAddr, server))
}

// Server implements the web server specification found at
// https://github.com/uis-dat520/labs/blob/master/lab2/README.md#web-server
type Server struct {
	// TODO(student): Add needed fields
	count int
	value int
}

// NewServer returns a new Server with all required internal state initialized.
// NOTE: It should NOT start to listen on an HTTP endpoint.
func NewServer() *Server {
	s := &Server{}
	// TODO(student): Implement
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	str := r.FormValue("value")
	fizzbuzzpath := "/fizzbuzz"
	s.count++
	// fmt.Println(fizzbuzzpath)
	if path == "/" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!\n"))
	} else if path == "/counter" {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("counter: " + strconv.Itoa(s.count) + "\n"))

	} else if path == "/lab2" {
		w.WriteHeader(http.StatusMovedPermanently)
		l := "<a href=\"http://www.github.com/dat520-2017/labs/tree/master/lab2\">Moved Permanently</a>.\n\n"
		w.Write([]byte(l))
	} else if path == fizzbuzzpath {
		i, err := strconv.Atoi(str)
		checkError(err)
		w.WriteHeader(http.StatusOK)
		f := float64(i)
		if str == "" {
			w.Write([]byte("no value provided\n"))
		} else if err != nil {
			w.Write([]byte("not an integer\n"))
		}
		if err == nil {
			if math.Mod(f, 3) == 0 && math.Mod(f, 5) == 0 {
				w.Write([]byte("fizzbuzz\n"))
			} else if math.Mod(f, 3) == 0 {
				w.Write([]byte("fizz\n"))
			} else if math.Mod(f, 5) == 0 {
				w.Write([]byte("buzz\n"))
			} else {
				w.Write([]byte(strconv.Itoa(i) + "\n"))
			}
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found\n"))
	}

}

func checkError(err error) {
	if err != nil {

	}
}
