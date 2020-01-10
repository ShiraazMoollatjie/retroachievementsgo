package retroachievementsgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"testing"
)

func TestTopTenUsers(t *testing.T) {
	var res *TopTenUsersResp
	b := unmarshalGoldenFileBytes(t, "toptenusers.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ttu, err := c.GetTopTenUsers()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ttu, res) {
		t.Errorf("Unexpected top ten users. Expected:\n%v\nActual:\n%v\n", ttu, res)
	}
}

func TestConsoleIDs(t *testing.T) {
	var res *ConsoleIDsResp
	b := unmarshalGoldenFileBytes(t, "console_ids.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.GetConsoleIDs()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", ci, res)
	}
}

func TestGameList(t *testing.T) {
	var res *GameListResp
	b := unmarshalGoldenFileBytes(t, "game_list.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.GetGameList()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", ci, res)
	}
}

func TestGameInfo(t *testing.T) {
	var res *GameInfoResp
	b := unmarshalGoldenFileBytes(t, "game_info.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.GetGameInfo()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", ci, res)
	}
}

func TestGameInfoExtended(t *testing.T) {
	var res *GameInfoExtendedResp
	b := unmarshalGoldenFileBytes(t, "game_info_extended.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.GetGameInfoExtended()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", ci, res)
	}
}

func unmarshalGoldenFileBytes(t *testing.T, filename string, payload interface{}) []byte {
	p := filepath.Join("testdata", filename)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		t.Errorf("error reading testdata. Error: %+v", err)
	}

	err = json.Unmarshal(b, &payload)
	if err != nil {
		t.Errorf("error unmarshalling payload. Error: %+v", err)
	}

	return b
}
