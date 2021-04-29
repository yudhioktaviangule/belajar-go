package main

import (
	"github.com/denil/promethee/src/alternatif"
	"github.com/denil/promethee/src/kriteria"
	"github.com/denil/promethee/src/login"
	"github.com/denil/promethee/src/nilai"
)

func RouteKriteria() {
	crtieriaIndex := kriteria.Index
	criteriaSearch := kriteria.IndexSearcher
	criteriaCreate := kriteria.Create
	criteriaUpdate := kriteria.Update
	criteriaDelete := kriteria.Delete
	mux.HandleFunc("/api/criteria", crtieriaIndex)
	mux.HandleFunc("/api/criteria/s", criteriaSearch)
	mux.HandleFunc("/api/criteria/save", criteriaCreate)
	mux.HandleFunc("/api/criteria/update", criteriaUpdate)
	mux.HandleFunc("/api/criteria/delete", criteriaDelete)
}

func RouteAlternatif() {
	alterHandlerIndex := alternatif.Index
	alterHandlerShow := alternatif.Show
	alterHandlerUpdate := alternatif.Update
	alterHandlerDelete := alternatif.Delete
	alterHandlerCreate := alternatif.Save
	mux.HandleFunc("/api/alternatif", alterHandlerIndex)
	mux.HandleFunc("/api/alternatif/create", alterHandlerCreate)
	mux.HandleFunc("/api/alternatif/show", alterHandlerShow)
	mux.HandleFunc("/api/alternatif/update", alterHandlerUpdate)
	mux.HandleFunc("/api/alternatif/delete", alterHandlerDelete)
}

func RouteNilai() {

	nilaiIndex := nilai.Index
	nilaiCreate := nilai.Create
	nilaiUpdate := nilai.EditNilai

	mux.HandleFunc("/api/nilai", nilaiIndex)
	mux.HandleFunc("/api/nilai/save", nilaiCreate)
	mux.HandleFunc("/api/nilai/update", nilaiUpdate)
}

func RouteLogin() {
	loginAlt := login.DoLogin
	lToken := login.Tokenku
	mux.HandleFunc("/login", loginAlt)
	mux.HandleFunc("/login/token", lToken)
}
