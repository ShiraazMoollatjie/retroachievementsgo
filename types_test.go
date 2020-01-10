package retroachievementsgo

import (
	"encoding/json"
	"testing"
)

type TestPlaceholder struct {
	Test RetroInt
}

func TestUnmarshalRetroInt(t *testing.T) {
	type RetroIntPlaceHolder struct {
		Value RetroInt
	}

	testCases := []struct {
		desc     string
		json     string
		expected RetroInt
	}{
		{"Basic json", `{"value":"123"}`, RetroInt(123)},
		{"Test zero", `{"value":0}`, RetroInt(0)},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var res *RetroIntPlaceHolder
			json.Unmarshal([]byte(tC.json), &res)
			if res.Value != tC.expected {
				t.Errorf("Incorrect RetroInt. \nExpected: %v\nActual: %v\n", tC.expected, res.Value)
			}
		})
	}
}
