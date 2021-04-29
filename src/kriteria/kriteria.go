package kriteria

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/denil/promethee/src/config/helper"
)

func isMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method == method {
		return true
	} else {
		helper.ThrowError("METHOD "+r.Method+" NOT SUPPORTED", 401, w)
		return false
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	validasi := isMethod("GET", w, r)
	if !validasi {
		return
	}
	var interf []interface{} = GetList()
	helper.JsonResponse(w, interf)
}

func IndexSearcher(w http.ResponseWriter, r *http.Request) {
	validasi := isMethod("GET", w, r)
	if !validasi {
		return
	}
	q := r.URL.Query().Get("q")
	if q == "" {
		helper.ThrowError("No QueryString", 500, w)
		return
	}
	var interf []interface{} = GetSearchedList(q)
	helper.JsonResponse(w, interf)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var post PostKriteria
	if !isMethod("POST", w, r) {
		return
	}
	param := helper.GetParam(r)

	slackerMarshall, _ := json.Marshal(param)
	json.Unmarshal(slackerMarshall, &post)

	Simpan(post)
	var interf interface{} = map[string]interface{}{
		"message": "OK",
		"code":    200,
		"PARAM":   param,
	}
	helper.JsonResponse(w, interf)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if !isMethod("POST", w, r) {
		return
	}
	var post Kriteria
	q := r.URL.Query().Get("id")
	if q == "" {
		helper.ThrowError("No ID as QueryString", 500, w)
		return
	}
	umarshal := helper.GetParam(r)
	marsh, _ := json.Marshal(umarshal)
	json.Unmarshal(marsh, &post)
	post.ID, _ = strconv.ParseInt(q, 10, 64)
	if !IsCriteria(post.ID) {
		helper.ThrowError(post.Name+" is not a criteria", 404, w)
		return
	}
	edit(post)
	var response interface{} = map[string]interface{}{
		"message": "Berhasil Menyimpan Kriteria",
		"code":    200,
		"updated_criteria": map[string]interface{}{
			"id":    post.ID,
			"name":  post.Name,
			"bobot": post.Bobot,
		},
	}
	helper.JsonResponse(w, response)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	if !isMethod("POST", w, r) {
		return
	}
	var post Kriteria
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		helper.ThrowError("Expected QueryString", 404, w)
		return
	}
	post.ID = id
	post = getCritByID(post.ID)
	if !IsCriteria(post.ID) {
		helper.ThrowError("NOT A CRITERIA", 404, w)
		return
	}
	Hapus(post.ID)
	respons := map[string]interface{}{
		"message": "Delete Berhasil",
		"code":    200,
		"_deleted": map[string]interface{}{
			"id":    post.ID,
			"name":  post.Name,
			"bobot": post.Bobot,
		},
	}
	helper.JsonResponse(w, respons)
}
