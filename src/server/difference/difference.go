package difference

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Difference _
type Difference struct {
	Status    map[string]string
	Proto     map[string]string
	Header    map[string]header
	Body      map[string]string
	Identical equal
}

type equal struct {
	Status 	bool
	Proto 	bool
	Body 	bool
	Header 	bool
}

type header map[string][]string

// New _
func New() *Difference {
	return &Difference{
		make(map[string]string),
		make(map[string]string),
		make(map[string]header),
		make(map[string]string),
		equal{false, false, false, false},
	}
}

// CompareResponses _
func (JL *Difference) CompareResponses(refResp *http.Response, expResp *http.Response) ([]byte, error) {

	JL.compareStatus(refResp.Status, expResp.Status)
	JL.compareProto(refResp.Proto, expResp.Proto)
	JL.compareHeader(refResp.Header, expResp.Header)
	err := JL.compareBody(refResp.Body, expResp.Body)
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
	payload, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}

	return string(payload), nil
}

// CompareStatus _
func (JL *Difference) compareStatus(statusA string, statusB string) {

	if statusA == statusB {
		JL.Identical.Status = true
	} else {
		JL.Identical.Status = false
	}

	m := make(map[string]string)
	m["RefResponse"] = statusA
	m["ExpResponse"] = statusB
	JL.Status = m
}

// CompareProto _
func (JL *Difference) compareProto(protoA string, protoB string) {
	if protoA == protoB {
		JL.Identical.Proto = true
	} else {
		JL.Identical.Proto = false
	}

	m := make(map[string]string)
	m["RefResponse"] = protoA
	m["ExpResponse"] = protoB
	JL.Proto = m
}

// CompareHeader _
func (JL *Difference) compareHeader(headerA map[string][]string, headerB map[string][]string) {

	isIdentical := true

	for key, value := range headerA {
		for i, str := range value {
			if headerB[key] != nil {
				if str != headerB[key][i] {
					isIdentical = false
					break
				}
			} else {
				isIdentical = false
				break
			}
		}
	}

	JL.Identical.Header = isIdentical

	m := make(map[string]header)
	m["RefResponse"] = headerA
	m["ExpResponse"] = headerB
	JL.Header = m
}

// CompareBody _
func (JL *Difference) compareBody(bodyA io.ReadCloser, bodyB io.ReadCloser) error {
	bodyAStr, err := BodyToString(bodyA)
	if err != nil {
		return err
	}

	bodyBStr, err := BodyToString(bodyB)
	if err != nil {
		return err
	}

	if bodyAStr == bodyBStr {
		JL.Identical.Body = true
	} else {
		JL.Identical.Body = false
	}

	m := make(map[string]string)
	m["RefResponse"] = bodyAStr
	m["ExpResponse"] = bodyBStr
	JL.Body = m

	return nil
}
