package nilai

import (
	"net/http"

	"github.com/denil/promethee/src/config/helper"
)

func EditNilai(w http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		helper.JsonResponse(w, map[string]interface{}{
			"message": "Method " + request.Method + " Not supported",
		})
		return
	}
	update := Edit(w, request)
	if update {
		helper.JsonResponse(w, map[string]interface{}{"message": "OK"})
	} else {
		helper.JsonResponse(w, map[string]interface{}{"message": "FAILED UPDATE DATA", "code": 500})

	}
}
func Create(w http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		helper.JsonResponse(w, map[string]interface{}{
			"message": "Method " + request.Method + " Not supported",
		})
		return
	}
	arrNilai := Simpan(w, request)
	if arrNilai {
		helper.JsonResponse(w, map[string]interface{}{`message`: `Berhasil Menyimpan Nilai`})

	}
}

func Index(w http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		helper.JsonResponse(w, map[string]interface{}{
			"message": "Method " + request.Method + " Not supported",
		})
		return
	}
	arrNilai := getList(w, request)
	helper.JsonResponse(w, arrNilai)
}
