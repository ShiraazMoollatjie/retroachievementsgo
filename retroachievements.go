package retroachievementsgo

import (
	"fmt"
	"strconv"
	"time"
)

func (c *Client) GetTopTenUsers() (*TopTenUsersResp, error) {
	var res TopTenUsersResp
	err := c.get(c.getURL("top_ten", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetConsoleIDs() (*ConsoleIDsResp, error) {
	var res ConsoleIDsResp
	err := c.get(c.getURL("console_id", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetGameList(consoleID string) (*GameListResp, error) {
	var res GameListResp
	var p = make(map[string]string)
	if consoleID != "" {
		p["console"] = consoleID
	}

	err := c.get(c.getURL("game_list", p), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetGameInfo(gameID string) (*GameInfoResp, error) {
	var res GameInfoResp
	var p = make(map[string]string)
	if gameID != "" {
		p["game"] = gameID
	}

	err := c.get(c.getURL("game_info", p), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetGameInfoExtended(gameID string) (*GameInfoExtendedResp, error) {
	var res GameInfoExtendedResp
	var p = make(map[string]string)
	if gameID != "" {
		p["game"] = gameID
	}

	err := c.get(c.getURL("game_info_extended", p), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetGameProgress(gameID string) (*GameProgressResp, error) {
	var res GameProgressResp
	var p = make(map[string]string)
	if gameID != "" {
		p["game"] = gameID
	}

	err := c.get(c.getURL("game_progress", p), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUserRank(user, gameID string) (*UserRankResp, error) {
	var res UserRankResp
	var p = make(map[string]string)
	if gameID != "" {
		p["game"] = gameID
	}
	if user != "" {
		p["member"] = user
	}

	err := c.get(c.getURL("user_rank", p), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUserRecent(user string) (*UserRecentResp, error) {
	var res UserRecentResp
	var p = make(map[string]string)
	if user != "" {
		p["member"] = user
	}

	err := c.get(c.getURL("user_recent", p), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUserProgress(user, gameID string) (*UserProgressResp, error) {
	var res UserProgressResp
	var p = make(map[string]string)
	if gameID != "" {
		p["game"] = gameID
	}
	if user != "" {
		p["member"] = user
	}

	err := c.get(c.getURL("user_progress", p), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUserSummary(user string) (*UserSummaryResp, error) {
	var res UserSummaryResp
	var p = make(map[string]string)
	if user != "" {
		p["member"] = user
	}

	err := c.get(c.getURL("user_summary", p), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListUserAchievementsByDate(user string, start time.Time) (*UserByDateResp, error) {
	var res UserByDateResp
	var p = make(map[string]string)
	if user != "" {
		p["member"] = user
	}

	if !start.IsZero() {
		p["date"] = strconv.FormatInt(start.Unix(), 10)
	}

	err := c.get(c.getURL("user_by_date", p), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListUserAchievementsByDateRange(user string, start, end time.Time) (*UserByDateResp, error) {
	var res UserByDateResp
	var p = make(map[string]string)
	if user != "" {
		p["member"] = user
	}

	if !start.IsZero() {
		p["startdate"] = strconv.FormatInt(start.Unix(), 10)
	}

	if !end.IsZero() {
		p["enddate"] = strconv.FormatInt(end.Unix(), 10)
	}

	err := c.get(c.getURL("user_by_date", p), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) getURL(endpoint string, params map[string]string) string {
	qp := c.buildReqQueryParams()
	for k, v := range params {
		qp.Add(k, v)
	}
	return fmt.Sprintf("%s/%s.php?%s", c.baseURL, endpoint, qp.Encode())
}
