package alternatif

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/denil/promethee/src/config/db"
	"github.com/denil/promethee/src/config/helper"
	"github.com/denil/promethee/src/nilai"
)

var DBKU *sql.DB

func initDB() {
	DBKU = db.Koneksi()
}

func Index(w http.ResponseWriter, request *http.Request) {

	initDB()
	var alter []Alternatif = GetAlternative()
	var alters []interface{}
	for i, e := range alter {
		alters = append(alters, map[string]interface{}{
			"id":   e.ID,
			"name": e.Name,
		})
		log.Println(i)
	}
	defer DBKU.Close()
	helper.JsonResponse(w, alters)
}

func Show(w http.ResponseWriter, request *http.Request) {

	initDB()
	var get AlterGet
	decode := json.NewDecoder(request.Body)
	err := decode.Decode(&get)
	var response interface{}
	if err != nil {
		response = map[string]interface{}{
			"message": "Error Getting PARAM DATA",
		}
	} else {
		var alter Alternatif = ShowAlter(get.ID)
		response = map[string]interface{}{
			"result": map[string]interface{}{
				"message": "OK",
				"data": map[string]interface{}{
					"id":   alter.ID,
					"name": alter.Name,
				},
			},
		}

	}
	helper.JsonResponse(w, response)

}

func Update(w http.ResponseWriter, request *http.Request) {

	method := request.Method
	var post Alternatif
	decoder := json.NewDecoder(request.Body)
	var ids string = request.URL.Query().Get("id")
	xid, err := strconv.ParseInt(ids, 10, 64)
	var validasi bool = true
	if err != nil {
		validasi = false
	}
	post.ID = xid
	if method != http.MethodPost {
		validasi = false
	}
	err = decoder.Decode(&post)
	if err != nil {
		validasi = false
	}
	message := "Update Data Berhasil"
	status := 200
	if !validasi {
		message = "Update Data Gagal"
		status = 500
	}
	message = UpdateAlternatif(post)
	helper.JsonResponse(w, map[string]interface{}{
		"error":   (status == 500),
		"message": message,
		"_post": map[string]interface{}{
			"id":   post.ID,
			"name": post.Name,
		},
	})
}

func Delete(w http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var id string = request.URL.Query().Get("id")
		mconnect := db.Koneksi()
		mysql := "DELETE FROM alternatif WHERE id=?"
		stmt, err := mconnect.Prepare(mysql)
		if err != nil {
			log.Fatal("ERROR QUERY")
		}
		stmt.Query(id)
		defer stmt.Close()
		defer mconnect.Close()
		intx, _ := strconv.ParseInt(id, 10, 64)
		nilai.DeleteByAlter(intx)
		helper.JsonResponse(w, map[string]interface{}{
			"code":    200,
			"message": "Done Response",
		})
	} else {
		helper.JsonResponse(w, map[string]interface{}{
			"code":    401,
			"message": "Tidak menerima method POST",
		})
	}
}

func Save(w http.ResponseWriter, request *http.Request) {

	var post AlternatifInsert
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&post)
	message := "Success"
	if err != nil {
		message = "Failed Get Param Data"
	} else {
		message = simpan(post)
	}
	helper.JsonResponse(w, map[string]interface{}{
		"message": message,
	})
}
