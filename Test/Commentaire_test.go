package Test

import (
	code "main/code"
	"testing"
)

func TestZoneSaisi(t *testing.T) {
	var chaine string
	chaine = "test"
	got := code.ZoneSaisi(chaine)
	want := chaine

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}
