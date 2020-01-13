package retroachievementsgo

import (
	"encoding/json"
	"strconv"
)

// RetroInt is a retroachievement type that is a string for a number > 0 or an
// int when it is zero
type RetroInt int

func (ri *RetroInt) UnmarshalJSON(data []byte) error {
	i, err := strconv.Atoi(string(data))
	if err != nil {
		// might be a string
		var str string
		err := json.Unmarshal(data, &str)
		if err != nil {
			return err
		}

		i, err = strconv.Atoi(str)
		if err != nil {
			return err
		}
	}

	*ri = RetroInt(i)
	return nil
}

// TopTenUsersResp is the top ten users. Generated using https://mholt.github.io/json-to-go/.
type TopTenUsersResp struct {
	Top10 struct {
		Place1 struct {
			User      interface{} `json:"user"`
			Score     interface{} `json:"score"`
			Trueratio interface{} `json:"trueratio"`
		} `json:"place_1"`
		Place2 struct {
			User      interface{} `json:"user"`
			Score     interface{} `json:"score"`
			Trueratio interface{} `json:"trueratio"`
		} `json:"place_2"`
		Place3 struct {
			User      interface{} `json:"user"`
			Score     interface{} `json:"score"`
			Trueratio interface{} `json:"trueratio"`
		} `json:"place_3"`
		Place4 struct {
			User      interface{} `json:"user"`
			Score     interface{} `json:"score"`
			Trueratio interface{} `json:"trueratio"`
		} `json:"place_4"`
		Place5 struct {
			User      interface{} `json:"user"`
			Score     interface{} `json:"score"`
			Trueratio interface{} `json:"trueratio"`
		} `json:"place_5"`
		Place6 struct {
			User      interface{} `json:"user"`
			Score     interface{} `json:"score"`
			Trueratio interface{} `json:"trueratio"`
		} `json:"place_6"`
		Place7 struct {
			User      interface{} `json:"user"`
			Score     interface{} `json:"score"`
			Trueratio interface{} `json:"trueratio"`
		} `json:"place_7"`
		Place8 struct {
			User      interface{} `json:"user"`
			Score     interface{} `json:"score"`
			Trueratio interface{} `json:"trueratio"`
		} `json:"place_8"`
		Place9 struct {
			User      interface{} `json:"user"`
			Score     interface{} `json:"score"`
			Trueratio interface{} `json:"trueratio"`
		} `json:"place_9"`
		Place10 struct {
			User      interface{} `json:"user"`
			Score     interface{} `json:"score"`
			Trueratio interface{} `json:"trueratio"`
		} `json:"place_10"`
	} `json:"top10"`
}

type ConsoleIDsResp struct {
	Console [][]struct {
		ID   string `json:"ID"`
		Name string `json:"Name"`
	} `json:"console"`
}

type GameListResp struct {
	Game [][]struct {
		Title       string `json:"Title"`
		ID          string `json:"ID"`
		ConsoleID   string `json:"ConsoleID"`
		ImageIcon   string `json:"ImageIcon"`
		ConsoleName string `json:"ConsoleName"`
	} `json:"game"`
}

type GameInfoResp struct {
	Title        string      `json:"Title"`
	ForumTopicID string      `json:"ForumTopicID"`
	ConsoleID    string      `json:"ConsoleID"`
	ConsoleName  string      `json:"ConsoleName"`
	Flags        interface{} `json:"Flags"`
	ImageIcon    string      `json:"ImageIcon"`
	GameIcon     string      `json:"GameIcon"`
	ImageTitle   string      `json:"ImageTitle"`
	ImageIngame  string      `json:"ImageIngame"`
	ImageBoxArt  string      `json:"ImageBoxArt"`
	Publisher    string      `json:"Publisher"`
	Developer    string      `json:"Developer"`
	Genre        string      `json:"Genre"`
	Released     string      `json:"Released"`
	GameTitle    string      `json:"GameTitle"`
	Console      string      `json:"Console"`
}

type GameInfoExtendedResp struct {
	ID                         int           `json:"ID"`
	Title                      string        `json:"Title"`
	ConsoleID                  int           `json:"ConsoleID"`
	ForumTopicID               int           `json:"ForumTopicID"`
	Flags                      int           `json:"Flags"`
	ImageIcon                  string        `json:"ImageIcon"`
	ImageTitle                 string        `json:"ImageTitle"`
	ImageIngame                string        `json:"ImageIngame"`
	ImageBoxArt                string        `json:"ImageBoxArt"`
	Publisher                  string        `json:"Publisher"`
	Developer                  string        `json:"Developer"`
	Genre                      string        `json:"Genre"`
	Released                   string        `json:"Released"`
	IsFinal                    bool          `json:"IsFinal"`
	ConsoleName                string        `json:"ConsoleName"`
	RichPresencePatch          string        `json:"RichPresencePatch"`
	NumAchievements            int           `json:"NumAchievements"`
	NumDistinctPlayersCasual   int           `json:"NumDistinctPlayersCasual"`
	NumDistinctPlayersHardcore int           `json:"NumDistinctPlayersHardcore"`
	Achievements               []interface{} `json:"Achievements"`
}

type GameProgressResp struct {
	ID                         int           `json:"ID"`
	Title                      string        `json:"Title"`
	ConsoleID                  int           `json:"ConsoleID"`
	ForumTopicID               int           `json:"ForumTopicID"`
	Flags                      int           `json:"Flags"`
	ImageIcon                  string        `json:"ImageIcon"`
	ImageTitle                 string        `json:"ImageTitle"`
	ImageIngame                string        `json:"ImageIngame"`
	ImageBoxArt                string        `json:"ImageBoxArt"`
	Publisher                  string        `json:"Publisher"`
	Developer                  string        `json:"Developer"`
	Genre                      string        `json:"Genre"`
	Released                   string        `json:"Released"`
	IsFinal                    bool          `json:"IsFinal"`
	ConsoleName                string        `json:"ConsoleName"`
	RichPresencePatch          string        `json:"RichPresencePatch"`
	NumAchievements            int           `json:"NumAchievements"`
	NumDistinctPlayersCasual   int           `json:"NumDistinctPlayersCasual"`
	NumDistinctPlayersHardcore int           `json:"NumDistinctPlayersHardcore"`
	Achievements               []interface{} `json:"Achievements"`
	NumAwardedToUser           int           `json:"NumAwardedToUser"`
	NumAwardedToUserHardcore   int           `json:"NumAwardedToUserHardcore"`
	UserCompletion             int           `json:"UserCompletion"`
	UserCompletionHardcore     int           `json:"UserCompletionHardcore"`
}

type UserRankResp struct {
	Score int    `json:"Score"`
	Rank  string `json:"Rank"`
}

type UserRecentResp struct {
	Recent [][]struct {
		GameID                  string      `json:"GameID"`
		ConsoleID               string      `json:"ConsoleID"`
		ConsoleName             string      `json:"ConsoleName"`
		Title                   string      `json:"Title"`
		ImageIcon               string      `json:"ImageIcon"`
		LastPlayed              string      `json:"LastPlayed"`
		MyVote                  interface{} `json:"MyVote"`
		NumPossibleAchievements RetroInt    `json:"NumPossibleAchievements"`
		PossibleScore           RetroInt    `json:"PossibleScore"`
		NumAchieved             RetroInt    `json:"NumAchieved"`
		ScoreAchieved           RetroInt    `json:"ScoreAchieved"`
	} `json:"recent"`
}

type UserProgress struct {
	NumPossibleAchievements RetroInt `json:"NumPossibleAchievements"`
	PossibleScore           RetroInt `json:"PossibleScore"`
	NumAchieved             RetroInt `json:"NumAchieved"`
	ScoreAchieved           RetroInt `json:"ScoreAchieved"`
	NumAchievedHardcore     RetroInt `json:"NumAchievedHardcore"`
	ScoreAchievedHardcore   RetroInt `json:"ScoreAchievedHardcore"`
}

type UserProgressResp map[string]UserProgress

type Award struct {
	NumPossibleAchievements RetroInt `json:"NumPossibleAchievements"`
	PossibleScore           RetroInt `json:"PossibleScore"`
	NumAchieved             RetroInt `json:"NumAchieved"`
	ScoreAchieved           RetroInt `json:"ScoreAchieved"`
	NumAchievedHardcore     RetroInt `json:"NumAchievedHardcore"`
	ScoreAchievedHardcore   RetroInt `json:"ScoreAchievedHardcore"`
}

type RecentAchievement struct {
	ID               string `json:"ID"`
	GameID           string `json:"GameID"`
	GameTitle        string `json:"GameTitle"`
	Title            string `json:"Title"`
	Description      string `json:"Description"`
	Points           string `json:"Points"`
	BadgeName        string `json:"BadgeName"`
	IsAwarded        string `json:"IsAwarded"`
	DateAwarded      string `json:"DateAwarded"`
	HardcoreAchieved string `json:"HardcoreAchieved"`
}

type UserSummaryResp struct {
	RecentlyPlayedCount int `json:"RecentlyPlayedCount"`
	RecentlyPlayed      []struct {
		GameID      string      `json:"GameID"`
		ConsoleID   string      `json:"ConsoleID"`
		ConsoleName string      `json:"ConsoleName"`
		Title       string      `json:"Title"`
		ImageIcon   string      `json:"ImageIcon"`
		LastPlayed  string      `json:"LastPlayed"`
		MyVote      interface{} `json:"MyVote"`
	} `json:"RecentlyPlayed"`
	MemberSince  string `json:"MemberSince"`
	LastActivity struct {
		ID           string      `json:"ID"`
		Timestamp    string      `json:"timestamp"`
		Lastupdate   string      `json:"lastupdate"`
		Activitytype string      `json:"activitytype"`
		User         string      `json:"User"`
		Data         interface{} `json:"data"`
		Data2        interface{} `json:"data2"`
	} `json:"LastActivity"`
	RichPresenceMsg    string                                  `json:"RichPresenceMsg"`
	LastGameID         string                                  `json:"LastGameID"`
	ContribCount       string                                  `json:"ContribCount"`
	ContribYield       string                                  `json:"ContribYield"`
	TotalPoints        string                                  `json:"TotalPoints"`
	TotalTruePoints    string                                  `json:"TotalTruePoints"`
	Permissions        string                                  `json:"Permissions"`
	Untracked          string                                  `json:"Untracked"`
	ID                 string                                  `json:"ID"`
	UserWallActive     string                                  `json:"UserWallActive"`
	Motto              string                                  `json:"Motto"`
	Rank               string                                  `json:"Rank"`
	Awarded            map[string]Award                        `json:"Awarded"`
	RecentAchievements map[string]map[string]RecentAchievement `json:"RecentAchievements"`
	Points             string                                  `json:"Points"`
	UserPic            string                                  `json:"UserPic"`
	Status             string                                  `json:"Status"`
}

type UserByDateResp struct {
	Achievement [][]struct {
		Date          string `json:"Date"`
		HardcoreMode  string `json:"HardcoreMode"`
		AchievementID string `json:"AchievementID"`
		Title         string `json:"Title"`
		Description   string `json:"Description"`
		BadgeName     string `json:"BadgeName"`
		Points        string `json:"Points"`
		Author        string `json:"Author"`
		GameTitle     string `json:"GameTitle"`
		GameIcon      string `json:"GameIcon"`
		GameID        string `json:"GameID"`
		ConsoleName   string `json:"ConsoleName"`
		CumulScore    int    `json:"CumulScore"`
		BadgeURL      string `json:"BadgeURL"`
		GameURL       string `json:"GameURL"`
	} `json:"achievement"`
}
