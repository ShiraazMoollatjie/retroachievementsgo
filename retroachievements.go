package retroachievementsgo

import (
	"fmt"
)

func (c *Client) GetTopTenUsers() (*TopTenUsersResp, error) {
	var res TopTenUsersResp
	err := c.get(c.getURL("top_ten"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetConsoleIDs() (*ConsoleIDsResp, error) {
	var res ConsoleIDsResp
	err := c.get(c.getURL("console_id"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// TODO(shiraaz): Add consoleID parameter
func (c *Client) GetGameList() (*GameListResp, error) {
	var res GameListResp
	err := c.get(c.getURL("game_list"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// TODO(shiraaz): Add gameID parameter
func (c *Client) GetGameInfo() (*GameInfoResp, error) {
	var res GameInfoResp
	err := c.get(c.getURL("game_info"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetGameInfoExtended() (*GameInfoExtendedResp, error) {
	var res GameInfoExtendedResp
	err := c.get(c.getURL("game_info_extended"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetGameProgress() string {
	panic("implement me!")
}

func (c *Client) GetUserRankAndScore() string {
	panic("implement me!")
}

func (c *Client) GetUserSummary() string {
	panic("implement me!")
}

func (c *Client) GetUserFeed() string {
	panic("implement me!")
}

func (c *Client) ListUserAchievementsByDate() string {
	panic("implement me!")
}

func (c *Client) ListUserAchievementsByDateRange() string {
	panic("implement me!")
}

func (c *Client) getURL(endpoint string) string {
	return fmt.Sprintf("%s/%s.php?%s", c.baseURL, endpoint, c.buildReqQueryParams().Encode())
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
