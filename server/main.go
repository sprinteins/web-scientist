package server

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Server _
type Server struct {
	host      string
	port      string
	stop      chan bool
	targetOne *url.URL
	targetTwo *url.URL
	proxyOne  *httputil.ReverseProxy
	proxyTwo  *httputil.ReverseProxy
	server    http.Server
}

// New _
func New(host string, port string) (s *Server) {
	return &Server{
		host: host,
		port: port,
		stop: make(chan bool),
	}
}

// Start _
func (s *Server) Start() {

	var address = fmt.Sprintf("%s:%s", s.host, s.port)
	var mux = http.NewServeMux()
	s.server = http.Server{Addr: address, Handler: mux}

	mux.HandleFunc("/", s.handle)

	go func() {
		log.Fatal(s.server.ListenAndServe())
	}()

	s.waitForStop(&s.stop, &s.server)
}

// Stop _
func (s *Server) Stop() {
	s.stop <- true
}

// SetTargetOne _
func (s *Server) SetTargetOne(target string) {
	s.targetOne, _ = url.Parse(target)
	s.proxyOne = httputil.NewSingleHostReverseProxy(s.targetOne)
}

// SetTargetTwo _
func (s *Server) SetTargetTwo(target string) {
	s.targetTwo, _ = url.Parse(target)
	s.proxyTwo = httputil.NewSingleHostReverseProxy(s.targetTwo)
}

func (s *Server) handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("X-WebScientist", "WebScientist")

	var reqOne, reqTwo = duplicate(req)

	resp, err := sendFurther(reqOne, s.targetOne)
	if err != nil {
		log.Fatal(err)
	}
	payloadOne, err := bodyToString(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = sendFurther(reqTwo, s.targetTwo)
	if err != nil {
		log.Fatal(err)
	}
	payloadTwo, err := bodyToString(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if payloadOne != payloadTwo {
		fmt.Fprintf(w, payloadOne)
	} else {
		fmt.Fprintf(w, payloadTwo)

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

func (s *Server) waitForStop(stop *chan bool, server *http.Server) {
	<-s.stop
	s.server.Shutdown(nil)
}
