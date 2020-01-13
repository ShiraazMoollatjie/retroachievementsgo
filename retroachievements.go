package retroachievementsgo

import (
	"fmt"
	"strconv"
	"time"
)

// GetTopTenUsers returns a current top ten list of user names and their scores.
func (c *Client) GetTopTenUsers() (*TopTenUsersResp, error) {
	var res TopTenUsersResp
	err := c.get(c.getURL("top_ten", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetConsoleIDs returns a list of supported consoles and their ID numbers.
func (c *Client) GetConsoleIDs() (*ConsoleIDsResp, error) {
	var res ConsoleIDsResp
	err := c.get(c.getURL("console_id", make(map[string]string)), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetGameList returns a list of supported games and Game IDs for the provided Console ID.
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

// GetGameInfo returns the game information for the specified Game ID.
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

// GetGameInfoExtended returns the game information for the specified Game ID with some extra information, including available achievements.
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

// GetGameProgress returns the game information for the specified Game ID with achievement progress information.
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

// GetUserRank returns the specified user's rank and overall score for a specified Game ID.
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

// GetUserRecent returns the specified user's recently played game list.
// The default number of results is 10.
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

// GetUserProgress returns the progress of a user based on the specified Game ID
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

// GetUserSummary returns a summary of recent results for the specified user.
// The default number of results is 10.
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

// ListUserAchievementsByDate returns a summary of last played games for the specified user on the specified date. Dates must be in UNIX format
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

// ListUserAchievementsByDateRange returns a summary of last played games for the specified user on the specified date range.
// Dates must be in UNIX format.
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
