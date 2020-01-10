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
