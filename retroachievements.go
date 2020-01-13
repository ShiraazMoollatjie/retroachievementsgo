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

// TODO(shiraaz): Add gameID parameter
func (c *Client) GetGameInfoExtended() (*GameInfoExtendedResp, error) {
	var res GameInfoExtendedResp
	err := c.get(c.getURL("game_info_extended"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// TODO(shiraaz): Add gameID parameter
func (c *Client) GetGameProgress() (*GameProgressResp, error) {
	var res GameProgressResp
	err := c.get(c.getURL("game_progress"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// TODO(shiraaz): Add gameID and member parameter
func (c *Client) GetUserRank() (*UserRankResp, error) {
	var res UserRankResp
	err := c.get(c.getURL("user_rank"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUserRecent() (*UserRecentResp, error) {
	var res UserRecentResp
	err := c.get(c.getURL("user_recent"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUserProgress() (*UserProgressResp, error) {
	var res UserProgressResp
	err := c.get(c.getURL("user_progress"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUserSummary() (*UserSummaryResp, error) {
	var res UserSummaryResp
	err := c.get(c.getURL("user_summary"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListUserAchievementsByDate() (*UserByDateResp, error) {
	var res UserByDateResp
	err := c.get(c.getURL("user_by_date"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListUserAchievementsByDateRange() (*UserByDateResp, error) {
	var res UserByDateResp
	err := c.get(c.getURL("user_by_date"), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) getURL(endpoint string) string {
	return fmt.Sprintf("%s/%s.php?%s", c.baseURL, endpoint, c.buildReqQueryParams().Encode())
}
