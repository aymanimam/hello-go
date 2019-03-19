package server

import (
	"fmt"
	"github.com/aymanimam/hello-go/gopherdojo/slide331/omikuji"
	"net/http"
)

// All omikujis
var allOmikujis *omikuji.AllOmikujis

func initialize() {
	allOmikujis = omikuji.GetAllOmikujis()
	allOmikujis.Init()
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	randOmikuji := allOmikujis.GetRandom(omikuji.OMIKUJI_WITH_DAIKICHI_MIN, omikuji.OMIKUJI_MAX)
	fmt.Fprint(w, randOmikuji)
}

func StartServer() {
	// Initialize
	initialize()

	// Handle requests
	http.HandleFunc("/omikuji", omikujiHandler)
	http.ListenAndServe(":8080", nil)
}
