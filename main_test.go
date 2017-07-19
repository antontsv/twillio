package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStraingToVoiceMailResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	StraightToVoiceMail(recorder, nil)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("returned wrong status code: got %v expected %v", status, http.StatusOK)
	}
	if err := checkWellFormetXML(recorder.Body.String()); err != nil {
		t.Errorf("body need to be valid XML: %v", err)
	}
}

func checkWellFormetXML(s string) error {
	d := xml.NewDecoder(strings.NewReader(s))
	t, err := d.Token()
	if err != nil {
		return err
	}

	v, ok := t.(xml.ProcInst)
	if !ok || v.Target != "xml" {
		return fmt.Errorf("No XML header detected at the start")
	}

	return nil
}
