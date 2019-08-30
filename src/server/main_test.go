package server_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"fmt"
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
	for !active {
		conn, err := net.Dial("tcp","localhost:2345")
		if err == nil {
			active = true
		}
		if tryConnect > 20 {
			panic("Web-Scientist cannot be reached.")
		}
		tryConnect++
		time.Sleep( 1 * time.Second )
		
		conn.Close()
	}
}

func teardown() {
	scientist.Stop()
}

func Test_By_Failed_Experiment_Reference_Sent(t *testing.T) {
	
	var reference, experiment = CreateNonEqualMocks()
	
	scientist.SetReference(reference.Address())
	scientist.SetExperiment(experiment.Address())
	
	active := false
	tryConnect := 0
	for !active {
		conn, err := net.Dial("tcp","localhost:9996")
		fmt.Print("Testing experimental connection in failed test.\n\n")
		if err == nil {
			active = true
		}
		if tryConnect > 20 {
			panic("Reference cannot be reached")
		}
		tryConnect++
		time.Sleep( 1 * time.Second )
		conn.Close()
	}
	

	fmt.Print("Testing reference connection in failed test.")
	active = false
	tryConnect = 0
	for !active {
		conn, err := net.Dial("tcp","localhost:9997")
		if err == nil {
			active = true
		}
		if tryConnect > 20 {
			panic("Experimental cannot be reached")
		}
		tryConnect++
		time.Sleep( 1 * time.Second )
		conn.Close()
	}

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
	
	var reference, experiment = CreateEqualMocks()
	
	scientist.SetReference(reference.Address())
	scientist.SetExperiment(experiment.Address())
	
	active := false
	tryConnect := 0
	for !active {
		conn, err := net.Dial("tcp","localhost:9998")

		fmt.Print("Testing experimental connection in succ test.\n\n")
		if err == nil {
			active = true
		}
		if tryConnect > 20 {
			panic("Reference cannot be reached")
		}
		tryConnect++
		time.Sleep( 1 * time.Second )
		conn.Close()
	}


	fmt.Print("Testing reference connection in succ test.")
	active = false
	tryConnect = 0
	for !active {
		conn, err := net.Dial("tcp","localhost:9999")
		if err == nil {
			active = true
		}
		if tryConnect > 20 {
			panic("Experimental cannot be reached")
		}
		tryConnect++
		time.Sleep( 1 * time.Second )
		conn.Close()
	}

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
