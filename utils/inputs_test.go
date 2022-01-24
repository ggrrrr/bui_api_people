package utils_test

import (
	"testing"

	"github.com/ggrrrr/bui_api_people/utils"
)

var (
	TT = map[string]string{
		"oO": "00",
		"оO": "00",
	}
)

func TestInput(t *testing.T) {
	for k, v := range TT {
		r := utils.NormalizeNumber(k)
		t.Logf("from %v to %v ?? %v", k, v, r)
		if v != r {
			t.Errorf("ERRROR %v to %v != %v", k, v, r)
		}
	}

}
