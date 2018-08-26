package server_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/trusz/web-scientist/server"
	"github.com/trusz/web-scientist/server/mock"
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

}

func teardown() {
	scientist.Stop()
}

func TestDiffernceIsDetected(t *testing.T) {

	var toUpper = mock.New("localhost", "9999", 0)
	var toLower = mock.New("localhost", "9998", 10)

	// defer toUpper.Stop()
	go toUpper.Start(nil)
	// defer toLower.Stop()
	go toLower.Start(toLowerCase)

	scientist.SetTargetOne("http://localhost:9999")
	scientist.SetTargetTwo("http://localhost:9998")
	var url = fmt.Sprintf("%s://%s:%s", PROTOCOL, HOST, PORT)

	var message = "TeSt"
	var payload = []byte(message)

	var resp, err = http.Post(url, "text/plain", bytes.NewBuffer(payload))
	if err != nil {
		t.Fail()
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
		log.Fatal(err)
	}

	var respPayload = string(body)
	if message != respPayload {
		t.Fail()
		log.Fatal(fmt.Sprintf("_%s_ != _%s_", message, respPayload))
	}
}

func toUpperCase(text string) string {
	return strings.ToUpper(text)
}

func toLowerCase(text string) string {
	return strings.ToLower(text)
}
