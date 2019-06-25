package jlog

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// JLog _
type JLog struct {
	Status    map[string]string
	Proto     map[string]string
	Header    map[string]map[string][]string
	Body      map[string]string
	Identical map[string]bool
}

// New _
func New() *JLog {
	return &JLog{
		make(map[string]string),
		make(map[string]string),
		make(map[string]map[string][]string),
		make(map[string]string),
		make(map[string]bool),
	}
}

// CompareResponses _
func (JL *JLog) CompareResponses(refResp *http.Response, expResp *http.Response) ([]byte, error) {

	JL.CompareStatus(refResp.Status, expResp.Status)
	JL.CompareProto(refResp.Proto, expResp.Proto)
	JL.CompareHeader(refResp.Header, expResp.Header)
	err := JL.CompareBody(refResp.Body, expResp.Body)
	if err != nil {
		return nil, err
	}

	str, err := json.MarshalIndent(JL, "", " ")
	if err != nil {
		return nil, err
	}
	return str, nil
}

// BodyToString _
func BodyToString(body io.ReadCloser) (string, error) {
	defer body.Close()
	payload, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}

	return string(payload), nil
}

// CompareStatus _
func (JL *JLog) CompareStatus(statusA string, statusB string) {
	
	if statusA == statusB {
		JL.Identical["Status"] = true
	} else {
		JL.Identical["Status"] = false
	}
		
	m := make(map[string]string)
	m["RefResponse"] = statusA
	m["ExpResponse"] = statusB
	JL.Status = m
}

// CompareProto _
func (JL *JLog) CompareProto(protoA string, protoB string) {
	if protoA == protoB {
		JL.Identical["Proto"] = true
	} else {
		JL.Identical["Proto"] = false
	}

	m := make(map[string]string)
	m["RefResponse"] = protoA
	m["ExpResponse"] = protoB
	JL.Proto = m
}

// CompareHeader _
func (JL *JLog) CompareHeader(headerA map[string][]string, headerB map[string][]string) {

	isIdentical := true

	for key, value := range headerA {
		for i, str := range value {
			if str != headerB[key][i] {
				isIdentical = false
				break
			}
		}
	}

	JL.Identical["Header"] = isIdentical

	m := make(map[string]map[string][]string)
	m["RefResponse"] = headerA
	m["ExpResponse"] = headerB
	JL.Header = m
}

// CompareBody _
func (JL *JLog) CompareBody(bodyA io.ReadCloser, bodyB io.ReadCloser) error {
	bodyAStr, err := BodyToString(bodyA)
	if err != nil {
		return err
	}

	bodyBStr, err := BodyToString(bodyB)
	if err != nil {
		return err
	}

	if bodyAStr == bodyBStr {
		JL.Identical["Body"] = true
	} else {
		JL.Identical["Body"] = false
	}

	m := make(map[string]string)
	m["RefResponse"] = bodyAStr
	m["ExpResponse"] = bodyBStr
	JL.Body = m

	return nil
}
