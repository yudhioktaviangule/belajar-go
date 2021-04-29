package main

import (
	"log"
	"net/http"

	"github.com/denil/promethee/src/config/auth"
	"github.com/denil/promethee/src/config/helper"
)

var mux *http.ServeMux

func Index(w http.ResponseWriter, req *http.Request) {
	helper.SetHeader(w)
	helper.JsonResponse(w, map[string]interface{}{
		"title":       "SPK Promethee",
		"founder":     "Denil-Novy",
		"description": "Sistem Penunjang Keputusan Promethee",
	})
}
func routeMe() {
	mux = http.DefaultServeMux
	mux.HandleFunc("/", Index)

	RouteLogin()
	RouteKriteria()
	RouteAlternatif()
	RouteNilai()

	var handler http.Handler = mux
	handler = auth.Auth(handler)
	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = handler
	log.Println("SERVER BERJALAN DI PORT 8080")
	server.ListenAndServe()

}
func main() {
	routeMe()
}
