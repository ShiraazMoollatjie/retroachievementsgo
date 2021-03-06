package retroachievementsgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"
	"time"
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
	testCases := []struct {
		desc      string
		consoleID string
	}{
		{
			desc: "No console id",
		},
		{
			desc:      "console id = 2",
			consoleID: "2",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *GameListResp
			b := unmarshalGoldenFileBytes(t, "game_list.json", &res)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				cID := r.URL.Query().Get("console")
				if cID != tC.consoleID {
					t.Errorf("Unexpected console ids. Expected:\n%v\nActual:\n%v\n", tC.consoleID, cID)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			c := NewClient(withBaseURL(ts.URL))
			ci, err := c.GetGameList(tC.consoleID)
			if err != nil {
				t.Errorf("error retrieving users. Error: %+v", err)
			}

			if !reflect.DeepEqual(ci, res) {
				t.Errorf("Unexpected game lists. Expected:\n%v\nActual:\n%v\n", res, ci)
			}
		})
	}
}

func TestGameInfo(t *testing.T) {
	testCases := []struct {
		desc   string
		gameID string
	}{
		{
			desc: "No gameID",
		},
		{
			desc:   "gameID = 2",
			gameID: "2",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *GameInfoResp
			b := unmarshalGoldenFileBytes(t, "game_info.json", &res)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gID := r.URL.Query().Get("game")
				if gID != tC.gameID {
					t.Errorf("Unexpected gameIDs. Expected:\n%v\nActual:\n%v\n", tC.gameID, gID)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			c := NewClient(withBaseURL(ts.URL))
			ci, err := c.GetGameInfo(tC.gameID)
			if err != nil {
				t.Errorf("error retrieving users. Error: %+v", err)
			}

			if !reflect.DeepEqual(ci, res) {
				t.Errorf("Unexpected game infos. Expected:\n%v\nActual:\n%v\n", res, ci)
			}
		})
	}

}

func TestGameInfoExtended(t *testing.T) {
	testCases := []struct {
		desc   string
		gameID string
	}{
		{
			desc: "No gameID",
		},
		{
			desc:   "gameID = 2",
			gameID: "2",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *GameInfoExtendedResp
			b := unmarshalGoldenFileBytes(t, "game_info_extended.json", &res)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gID := r.URL.Query().Get("game")
				if gID != tC.gameID {
					t.Errorf("Unexpected gameIDs. Expected:\n%v\nActual:\n%v\n", tC.gameID, gID)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			c := NewClient(withBaseURL(ts.URL))
			ci, err := c.GetGameInfoExtended(tC.gameID)
			if err != nil {
				t.Errorf("error retrieving users. Error: %+v", err)
			}

			if !reflect.DeepEqual(ci, res) {
				t.Errorf("Unexpected game info extended. Expected:\n%v\nActual:\n%v\n", res, ci)
			}
		})
	}
}

func TestGameProgress(t *testing.T) {
	testCases := []struct {
		desc   string
		gameID string
	}{
		{
			desc: "No gameID",
		},
		{
			desc:   "gameID = 2",
			gameID: "2",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *GameProgressResp
			b := unmarshalGoldenFileBytes(t, "game_progress.json", &res)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gID := r.URL.Query().Get("game")
				if gID != tC.gameID {
					t.Errorf("Unexpected gameIDs. Expected:\n%v\nActual:\n%v\n", tC.gameID, gID)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			c := NewClient(withBaseURL(ts.URL))
			ci, err := c.GetGameProgress(tC.gameID)
			if err != nil {
				t.Errorf("error retrieving users. Error: %+v", err)
			}

			if !reflect.DeepEqual(ci, res) {
				t.Errorf("Unexpected game progress. Expected:\n%v\nActual:\n%v\n", res, ci)
			}
		})
	}
}

func TestUserRecent(t *testing.T) {
	testCases := []struct {
		desc string
		user string
	}{
		{
			desc: "No user",
		},
		{
			desc: "user=shiraaz",
			user: "shiraaz",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *UserRecentResp
			b := unmarshalGoldenFileBytes(t, "user_recent.json", &res)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				uID := r.URL.Query().Get("member")
				if uID != tC.user {
					t.Errorf("Unexpected user. Expected:\n%v\nActual:\n%v\n", tC.user, uID)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			c := NewClient(withBaseURL(ts.URL))
			ci, err := c.GetUserRecent(tC.user)
			if err != nil {
				t.Errorf("error retrieving users. Error: %+v", err)
			}

			if !reflect.DeepEqual(ci, res) {
				t.Errorf("Unexpected user recents. Expected:\n%v\nActual:\n%v\n", res, ci)
			}
		})
	}
}

func TestUserRank(t *testing.T) {
	testCases := []struct {
		desc   string
		user   string
		gameID string
	}{
		{
			desc: "No gameID, user",
		},
		{
			desc:   "gameID=2, user=shiraaz",
			gameID: "2",
			user:   "shiraaz",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *UserRankResp
			b := unmarshalGoldenFileBytes(t, "user_rank.json", &res)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				uID := r.URL.Query().Get("member")
				if uID != tC.user {
					t.Errorf("Unexpected user. Expected:\n%v\nActual:\n%v\n", tC.user, uID)
				}
				gID := r.URL.Query().Get("game")
				if gID != tC.gameID {
					t.Errorf("Unexpected gameIDs. Expected:\n%v\nActual:\n%v\n", tC.user, uID)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			c := NewClient(withBaseURL(ts.URL))
			ci, err := c.GetUserRank(tC.user, tC.gameID)
			if err != nil {
				t.Errorf("error retrieving users. Error: %+v", err)
			}

			if !reflect.DeepEqual(ci, res) {
				t.Errorf("Unexpected user ranks. Expected:\n%v\nActual:\n%v\n", res, ci)
			}
		})
	}
}

func TestUserProgress(t *testing.T) {
	testCases := []struct {
		desc   string
		gameID string
		user   string
	}{
		{
			desc: "No gameID, user",
		},
		{
			desc:   "gameID=2, user=shiraaz",
			gameID: "2",
			user:   "shiraaz",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *UserProgressResp
			b := unmarshalGoldenFileBytes(t, "user_progress.json", &res)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				uID := r.URL.Query().Get("member")
				if uID != tC.user {
					t.Errorf("Unexpected user. Expected:\n%v\nActual:\n%v\n", tC.user, uID)
				}
				gID := r.URL.Query().Get("game")
				if gID != tC.gameID {
					t.Errorf("Unexpected gameIDs. Expected:\n%v\nActual:\n%v\n", tC.user, uID)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			c := NewClient(withBaseURL(ts.URL))
			ci, err := c.GetUserProgress(tC.user, tC.gameID)
			if err != nil {
				t.Errorf("error retrieving users. Error: %+v", err)
			}

			if !reflect.DeepEqual(ci, res) {
				t.Errorf("Unexpected user progress. Expected:\n%v\nActual:\n%v\n", res, ci)
			}
		})
	}
}

func TestUserSummary(t *testing.T) {
	testCases := []struct {
		desc string
		user string
	}{
		{
			desc: "No user",
		},
		{
			desc: "user=shiraaz",
			user: "shiraaz",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *UserSummaryResp
			b := unmarshalGoldenFileBytes(t, "user_summary.json", &res)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				uID := r.URL.Query().Get("member")
				if uID != tC.user {
					t.Errorf("Unexpected user. Expected:\n%v\nActual:\n%v\n", tC.user, uID)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			c := NewClient(withBaseURL(ts.URL))
			ci, err := c.GetUserSummary(tC.user)
			if err != nil {
				t.Errorf("error retrieving users. Error: %+v", err)
			}

			if !reflect.DeepEqual(ci, res) {
				t.Errorf("Unexpected user summary. Expected:\n%v\nActual:\n%v\n", res, ci)
			}
		})
	}
}

func TestListUserAchievementsByDate(t *testing.T) {
	testCases := []struct {
		desc        string
		user        string
		start       time.Time
		expectStart bool
	}{
		{
			desc: "No user, no start",
		},
		{
			desc:        "user=shiraaz",
			user:        "shiraaz",
			start:       time.Date(2010, 10, 2, 6, 0, 0, 0, time.UTC),
			expectStart: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *UserByDateResp
			b := unmarshalGoldenFileBytes(t, "user_by_date.json", &res)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				uID := r.URL.Query().Get("member")
				if uID != tC.user {
					t.Errorf("Unexpected user. Expected:\n%v\nActual:\n%v\n", tC.user, uID)
				}
				start := r.URL.Query().Get("date")
				exp := strconv.FormatInt(tC.start.Unix(), 10)
				if start != exp && tC.expectStart {
					t.Errorf("Unexpected start date. Expected:\n%v\nActual:\n%v\n", exp, start)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			c := NewClient(withBaseURL(ts.URL))
			ci, err := c.ListUserAchievementsByDate(tC.user, tC.start)
			if err != nil {
				t.Errorf("error retrieving users. Error: %+v", err)
			}

			if !reflect.DeepEqual(ci, res) {
				t.Errorf("Unexpected user achievement lists. "+
					"Expected:\n%v\nActual:\n%v\n", res, ci)
			}
		})
	}
}

func TestListUserAchievementsByDateRange(t *testing.T) {
	testCases := []struct {
		desc        string
		user        string
		start       time.Time
		expectStart bool
		end         time.Time
		expectEnd   bool
	}{
		{
			desc: "No user",
		},
		{
			desc:        "user=shiraaz",
			user:        "shiraaz",
			start:       time.Date(2010, 10, 2, 6, 0, 0, 0, time.UTC),
			expectStart: true,
			end:         time.Date(2015, 10, 2, 6, 0, 0, 0, time.UTC),
			expectEnd:   true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *UserByDateResp
			b := unmarshalGoldenFileBytes(t, "user_by_date.json", &res)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				uID := r.URL.Query().Get("member")
				if uID != tC.user {
					t.Errorf("Unexpected user. Expected:\n%v\nActual:\n%v\n", tC.user, uID)
				}
				start := r.URL.Query().Get("startdate")
				exp := strconv.FormatInt(tC.start.Unix(), 10)
				if start != exp && tC.expectStart {
					t.Errorf("Unexpected start date. Expected:\n%v\nActual:\n%v\n", exp, start)
				}

				end := r.URL.Query().Get("enddate")
				exp = strconv.FormatInt(tC.end.Unix(), 10)
				if end != exp && tC.expectStart {
					t.Errorf("Unexpected start date. Expected:\n%v\nActual:\n%v\n", exp, end)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			c := NewClient(withBaseURL(ts.URL))
			ci, err := c.ListUserAchievementsByDateRange(tC.user, tC.start, tC.end)
			if err != nil {
				t.Errorf("error retrieving users. Error: %+v", err)
			}

			if !reflect.DeepEqual(ci, res) {
				t.Errorf("Unexpected user achievements. Expected:\n%v\nActual:\n%v\n", res, ci)
			}
		})
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
