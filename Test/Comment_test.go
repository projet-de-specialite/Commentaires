package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/projet-de-specialite/Commentaires/api"
	"github.com/projet-de-specialite/Commentaires/config"
)

func TestGetNombreAvis(t *testing.T) {
	config.DatabaseInit()
	router := gin.Default()

	router.GET("/api/v1/comment/nombreAvis/:Id_Commentaire", api.GetNombreAvis)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/comment/nombreAvis/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Code de status attendu %d; obtenu %d", http.StatusOK, w.Code)
	}
}

func TestFindCommentById(t *testing.T) {
	config.DatabaseInit()
	router := gin.Default()

	router.GET("/api/v1/comment/findCommentById/:Id_Commentaire", api.FindCommentById)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/comment/findCommentById/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if (w.Code != http.StatusOK) && (w.Code != 201) {
		t.Errorf("Code de status attendu %d; obtenu %d", http.StatusOK, w.Code)
	}
}

func TestPrintCommentByPost(t *testing.T) {
	config.DatabaseInit()
	router := gin.Default()

	router.GET("/api/v1/comment/printComment/:Id_Post", api.PrintCommentByPost)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/comment/printComment/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if (w.Code != http.StatusOK) && (w.Code != 201) {
		t.Errorf("Code de status attendu %d; obtenu %d", http.StatusOK, w.Code)
	}
}

func TestDeleteComment(t *testing.T) {
	config.DatabaseInit()
	router := gin.Default()

	router.DELETE("/api/v1/comment/deleteComment/:Id_Commentaire", api.DeleteComment)

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/comment/deleteComment/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if (w.Code != http.StatusOK) && (w.Code != 201) {
		t.Errorf("Code de status attendu %d; obtenu %d", http.StatusOK, w.Code)
	}
}

func TestPostComment(t *testing.T) {
	config.DatabaseInit()
	router := gin.Default()

	router.POST("/api/v1/comment/newComment", api.PostComment)

	input := struct {
		Id_User int    `json:"Id_User"`
		Texte   string `json:"Texte"`
		Id_Post string `json:"Id_Post"`
	}{
		Id_User: 404,
		Texte:   "Test Unitaire",
		Id_Post: "158c9999-18da-4b5f-a8db-4891b1736574",
	}

	bodyReq, _ := json.Marshal(input)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/comment/newComment", bytes.NewBuffer(bodyReq))

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if (w.Code != http.StatusOK) && (w.Code != 201) {
		t.Errorf("Code de status attendu %d; obtenu %d", http.StatusOK, w.Code)
	}
}

func TestPostAvis(t *testing.T) {
	config.DatabaseInit()
	router := gin.Default()

	router.POST("/api/v1/comment/newAvis", api.PostAvis)

	input := struct {
		Id_Commentaire int  `json:"Id_Commentaire"`
		Id_User        int  `json:"Id_User"`
		Valeur         bool `json:"Valeur"`
	}{
		Id_Commentaire: 4,
		Id_User:        3,
		Valeur:         false,
	}

	bodyReq, _ := json.Marshal(input)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/comment/newAvis", strings.NewReader(string(bodyReq)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if (w.Code != http.StatusOK) && (w.Code != 201) {
		t.Errorf("Code de status attendu %d; obtenu %d", http.StatusOK, w.Code)
	}
}
