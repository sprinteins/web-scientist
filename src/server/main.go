package server

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"github.com/sprinteins/web-scientist/server/difference"
)

// Server _
type Server struct {
	host       string
	port       string
	stop       chan os.Signal
	reference  *url.URL
	experiment *url.URL
	server     http.Server
}

// New _
func New(host string, port string) (s *Server) {
	return &Server{
		host: host,
		port: port,
		stop: make(chan os.Signal, 1),
	}
}

// Start _
func (s *Server) Start() {

	var address = fmt.Sprintf("%s:%s", s.host, s.port)
	var mux = http.NewServeMux()
	s.server = http.Server{Addr: address, Handler: mux}
	mux.HandleFunc("/", s.handle)

	go func() {
		s.waitForStop(&s.stop, &s.server)
	}()

	log.Fatal(s.server.ListenAndServe())
}

// Address _
func (s *Server) Address() string {
	return fmt.Sprintf("http://%s:%s", s.host, s.port)
}

// Stop _
func (s *Server) Stop() {
	s.stop <- os.Interrupt
}

// SetReference _
func (s *Server) SetReference(target string) {
	s.reference, _ = url.Parse(target)
}

// SetExperiment _
func (s *Server) SetExperiment(target string) {
	s.experiment, _ = url.Parse(target)
}

// TODO:
// make the calls parallel
func (s *Server) handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("X-WebScientist", "WebScientist")

	var reqRef, reqExp = duplicate(req)

	respA, err := sendFurther(reqRef, s.reference)
	if err != nil {
		log.Fatal(err)
	}

	respB, err := sendFurther(reqExp, s.experiment)
	if err != nil {
		log.Fatal(err)
	}

	diff := difference.New()
	out, _ := diff.CompareResponses(respA, respB)
	ioutil.WriteFile("log.json", out, 0755)

}

func sendFurther(req *http.Request, url *url.URL) (*http.Response, error) {
	req.URL = url
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func duplicate(request *http.Request) (request1 *http.Request, request2 *http.Request) {
	b1 := new(bytes.Buffer)
	b2 := new(bytes.Buffer)
	w := io.MultiWriter(b1, b2)
	io.Copy(w, request.Body)
	defer request.Body.Close()
	request1 = &http.Request{
		Method:        request.Method,
		URL:           request.URL,
		Proto:         request.Proto,
		ProtoMajor:    request.ProtoMajor,
		ProtoMinor:    request.ProtoMinor,
		Header:        request.Header,
		Body:          ioutil.NopCloser(b1),
		Host:          request.Host,
		ContentLength: request.ContentLength,
		Close:         true,
	}
	request2 = &http.Request{
		Method:        request.Method,
		URL:           request.URL,
		Proto:         request.Proto,
		ProtoMajor:    request.ProtoMajor,
		ProtoMinor:    request.ProtoMinor,
		Header:        request.Header,
		Body:          ioutil.NopCloser(b2),
		Host:          request.Host,
		ContentLength: request.ContentLength,
		Close:         true,
	}
	return
}

func (s *Server) waitForStop(stop *chan os.Signal, server *http.Server) {
	<-s.stop
	err := s.server.Shutdown(nil)
	if err != nil {
		log.Fatal(s.server.Addr)
		log.Fatal(err)
	}
}
