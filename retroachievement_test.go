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
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
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
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
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
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
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
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
	}
}

func TestUserRecent(t *testing.T) {
	var res *UserRecentResp
	b := unmarshalGoldenFileBytes(t, "user_recent.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.GetUserRecent()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
	}
}

func TestUserRank(t *testing.T) {
	var res *UserRankResp
	b := unmarshalGoldenFileBytes(t, "user_rank.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.GetUserRank()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
	}
}

func TestGameProgress(t *testing.T) {
	var res *GameProgressResp
	b := unmarshalGoldenFileBytes(t, "game_progress.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.GetGameProgress()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
	}
}

func TestUserProgress(t *testing.T) {
	var res *UserProgressResp
	b := unmarshalGoldenFileBytes(t, "user_progress.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.GetUserProgress()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
	}
}

func TestUserSummary(t *testing.T) {
	var res *UserSummaryResp
	b := unmarshalGoldenFileBytes(t, "user_summary.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.GetUserSummary()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
	}
}

func TestListUserAchievementsByDate(t *testing.T) {
	var res *UserByDateResp
	b := unmarshalGoldenFileBytes(t, "user_by_date.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.ListUserAchievementsByDate()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
	}
}

func TestListUserAchievementsByDateRange(t *testing.T) {
	var res *UserByDateResp
	b := unmarshalGoldenFileBytes(t, "user_by_date.json", &res)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	c := NewClient(withBaseURL(ts.URL))
	ci, err := c.ListUserAchievementsByDateRange()
	if err != nil {
		t.Errorf("error retrieving users. Error: %+v", err)
	}

	if !reflect.DeepEqual(ci, res) {
		t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", res, ci)
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
