package Test

import (
	"main/models"
	"testing"
)

func TestToStringValide(t *testing.T) {
	var c models.Commentaire

	got := models.ToString(c)
	want := ""

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestToStringVide(t *testing.T) {
	var c models.Commentaire

	got := models.ToString(c)
	want := ""

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
