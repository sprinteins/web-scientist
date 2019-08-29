package server_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
	"net"
	"github.com/sprinteins/web-scientist/server"
	. "github.com/sprinteins/web-scientist/server/test_helpers"
)

const PROTOCOL = "http"
const HOST = "localhost"
const PORT = "2345"

var scientist *server.Server

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func setup() {
	scientist = server.New(HOST, PORT)
	go scientist.Start()

	active := false
	tryConnect := 0
	timeout := time.Duration(1 * time.Second)
	for active {
		_, err := net.DialTimeout("tcp","localhost:2345", timeout)
		if err == nil {
			active = true
		}
		if tryConnect > 10 {
			panic("Web-Scientist cannot be reached.")
		}
		tryConnect++
	}
}

func teardown() {
	scientist.Stop()
}

func Test_By_Failed_Experiment_Reference_Sent(t *testing.T) {

	active := true
	tryConnect := 0
	timeout := time.Duration(1 * time.Second)
	for !active {
		_, err := net.DialTimeout("tcp","localhost:9998", timeout)
		if err != nil {
			active = false
		}
		if tryConnect > 10 {
			panic("Experimental port already in use.")
		}
		tryConnect++
	}
	
	tryConnect = 0
	for !active {
		_, err := net.DialTimeout("tcp","localhost:9999", timeout)
		if err != nil {
			active = false
		}
		if tryConnect > 10 {
			panic("Reference port already in use")
		}
		tryConnect++
	}

	var reference, experiment = CreateNonEqualMocks()

	scientist.SetReference(reference.Address())
	scientist.SetExperiment(experiment.Address())

	var message = "TeSt"
	var payload = []byte(message)

	var resp, err = http.Post(scientist.Address(), "text/plain", bytes.NewBuffer(payload))
	Ok(t, err)

	var header = resp.Header.Get("X-WebScientist")
	Equals(t, "WebScientist", header)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	Ok(t, err)

	var respPayload = string(body)
	Equals(t, respPayload, message)

	reference.Stop()
	experiment.Stop()

}

func Test_By_Successfull_Experiment_Experiment_Sent(t *testing.T) {
	
	active := true
	tryConnect := 0
	timeout := time.Duration(1 * time.Second)
	for !active {
		_, err := net.DialTimeout("tcp","localhost:9998", timeout)
		if err != nil {
			active = false
		}
		if tryConnect > 10 {
			panic("Experimental port already in use.")
		}
		tryConnect++
	}
	
	tryConnect = 0
	for !active {
		_, err := net.DialTimeout("tcp","localhost:9999", timeout)
		if err != nil {
			active = false
		}
		if tryConnect > 10 {
			panic("Reference port already in use")
		}
		tryConnect++
	}
	
	var reference, experiment = CreateEqualMocks()

	scientist.SetReference(reference.Address())
	scientist.SetExperiment(experiment.Address())

	var message = "TeSt"
	var payload = []byte(message)

	var resp, err = http.Post(scientist.Address(), "text/plain", bytes.NewBuffer(payload))
	Ok(t, err)

	var header = resp.Header.Get("X-WebScientist")
	Equals(t, "WebScientist", header)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	Ok(t, err)

	var respPayload = string(body)
	Equals(t, respPayload, message)

	reference.Stop()
	experiment.Stop()
}
