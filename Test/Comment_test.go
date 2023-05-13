package Test

import (
	"testing"

	"github.com/projet-de-specialite/Commentaires/models"
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
