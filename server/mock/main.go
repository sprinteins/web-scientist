package mock

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Mock _
type Mock struct {
	host      string
	port      string
	delay     int
	stop      chan bool
	server    *http.Server
	transform func(text string) string
}

// Transformer _
type Transformer func(string) string

// New _
func New(host string, port string, delay int) (m *Mock) {
	return &Mock{
		host:      host,
		port:      port,
		delay:     delay,
		stop:      make(chan bool),
		transform: defaultTransform,
	}
}

// Start _
func (m *Mock) Start(
	transform Transformer,
) {

	if transform != nil {
		m.transform = transform
	}

	var address = fmt.Sprintf("%s:%s", m.host, m.port)
	var mux = http.NewServeMux()
	m.server = &http.Server{Addr: address, Handler: mux}

	mux.HandleFunc("/", m.handler)

	go func() {
		log.Fatal(m.server.ListenAndServe())
	}()

	m.waitFroStop(&m.stop, m.server)
}

// Stop _
func (m *Mock) Stop() {
	m.stop <- true
}

func (m *Mock) waitFroStop(stop *chan bool, server *http.Server) {
	<-m.stop
	m.server.Shutdown(nil)
}

func (m *Mock) handler(write http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	var payload = string(body)
	var modified = m.transform(payload)
	fmt.Fprintf(write, modified)

}

func defaultTransform(toTransform string) string {
	return toTransform
}

func checkOrigin(r *http.Request) bool {
	return true
}
