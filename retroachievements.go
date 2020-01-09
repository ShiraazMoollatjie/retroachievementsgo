package retroachievementsgo

import (
	"fmt"
)

func (c *Client) GetTopTenUsers() (*TopTenUsersResp, error) {
	u := fmt.Sprintf("%s/top_ten.php?%s", c.baseURL, c.buildReqQueryParams().Encode())

	var res TopTenUsersResp
	err := c.get(u, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetConsoleIDs() string {
	panic("implement me!")
}

func (c *Client) GetGameList() string {
	panic("implement me!")
}

func (c *Client) GetGameInfo() string {
	panic("implement me!")
}

func (c *Client) GetEnhancedGameInfo() string {
	panic("implement me!")
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
