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

// TODO
// make the calls on two threads and wait for both
func (s *Server) handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("X-WebScientist", "WebScientist")

	var reqRef, reqExp = duplicate(req)

	resp, err := sendFurther(reqRef, s.reference)
	if err != nil {
		log.Fatal(err)
	}
	payloadRef, err := bodyToString(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = sendFurther(reqExp, s.experiment)
	if err != nil {
		log.Fatal(err)
	}
	payloadExp, err := bodyToString(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if payloadRef != payloadExp {
		w.Header().Set("X-web-scientist-type", "reference")
		fmt.Fprintf(w, payloadRef)
	} else {
		w.Header().Set("X-web-scientist-type", "experiment")
		fmt.Fprintf(w, payloadExp)

	}

}

func sendFurther(req *http.Request, url *url.URL) (*http.Response, error) {
	req.URL = url
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func bodyToString(body io.ReadCloser) (string, error) {
	defer body.Close()
	payload, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}

	return string(payload), nil
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
