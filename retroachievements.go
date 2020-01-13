package retroachievementsgo

import (
	"fmt"
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

// TODO(shiraaz): Add gameID parameter
func (c *Client) GetGameInfoExtended() (*GameInfoExtendedResp, error) {
	var res GameInfoExtendedResp
	err := c.get(c.getURL("game_info_extended", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// TODO(shiraaz): Add gameID parameter
func (c *Client) GetGameProgress() (*GameProgressResp, error) {
	var res GameProgressResp
	err := c.get(c.getURL("game_progress", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// TODO(shiraaz): Add gameID and member parameter
func (c *Client) GetUserRank() (*UserRankResp, error) {
	var res UserRankResp
	err := c.get(c.getURL("user_rank", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUserRecent() (*UserRecentResp, error) {
	var res UserRecentResp
	err := c.get(c.getURL("user_recent", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUserProgress() (*UserProgressResp, error) {
	var res UserProgressResp
	err := c.get(c.getURL("user_progress", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUserSummary() (*UserSummaryResp, error) {
	var res UserSummaryResp
	err := c.get(c.getURL("user_summary", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListUserAchievementsByDate() (*UserByDateResp, error) {
	var res UserByDateResp
	err := c.get(c.getURL("user_by_date", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListUserAchievementsByDateRange() (*UserByDateResp, error) {
	var res UserByDateResp
	err := c.get(c.getURL("user_by_date", make(map[string]string)), &res)
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
