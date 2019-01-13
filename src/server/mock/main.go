package mock

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Mock _
type Mock struct {
	host      string
	port      string
	delay     int
	stop      chan os.Signal
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
		stop:      make(chan os.Signal, 1),
		transform: defaultTransform,
	}
}

// Address _
func (m *Mock) Address() string {
	return fmt.Sprintf("http://%s:%s", m.host, m.port)
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
		// log.Fatal(m.server.ListenAndServe())
		if err := m.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERR] server exited with: %s", err)
		}
	}()

	signal.Notify(m.stop, os.Interrupt, syscall.SIGTERM)
	m.waitFroStop(&m.stop, m.server)
}

// Stop _
func (m *Mock) Stop() {
	m.stop <- os.Interrupt
}

func (m *Mock) waitFroStop(stop *chan os.Signal, server *http.Server) {
	<-m.stop
	err := m.server.Shutdown(nil)
	if err != nil {
		log.Fatal(m.server.Addr)
		log.Fatal(err)
	}
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
