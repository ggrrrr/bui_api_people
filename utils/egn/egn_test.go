package egn_test

import (
	"testing"

	"github.com/ggrrrr/bui_api_people/utils/egn"
)

func TestEgn(t *testing.T) {
	egn1 := "o645196854"
	t1 := egn.Parse(egn1)
	t.Logf("t1: %+v", t1)
	if !t1.Ok {
		t.Error("asd")
	}

	egn2 := "4202136264"
	t2 := egn.Parse(egn2)
	t.Logf("%+v", t2)
	if !t2.Ok {
		t.Error("asd")
	}
	if t2.Gender != "male" {
		t.Error("asd")
	}

}
