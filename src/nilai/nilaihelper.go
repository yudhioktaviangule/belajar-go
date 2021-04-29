package nilai

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/denil/promethee/src/config/db"
	"github.com/denil/promethee/src/config/helper"
)

func getList(w http.ResponseWriter, r *http.Request) []interface{} {
	myConect := db.Koneksi()
	query := `
	SELECT 
		N.id,
		N.nilai,
		A.name as alternatif_name,
		A.id AS alternatif_id,
		C.name AS kriteria_name, 
		C.id AS kriteria_id 
	FROM alternatif A
	INNER JOIN nilai N ON A.id=N.alternatif_id
    INNER JOIN kriteria C ON N.kriteria_id=C.id ORDER BY A.id ASC`
	result, err := myConect.Query(query)
	if err != nil {
		log.Fatal("QUERY ERROR", err)
	}

	var nilaiArr []NilaiResponse
	for result.Next() {
		var nresp NilaiResponse
		var nilai NilaiHeadResp
		err = result.Scan(&nilai.ID, &nilai.Nilai, &nilai.NamaAlternatif, &nilai.IDAlternatif, &nilai.NamaKriteria, &nilai.IDKriteria)
		if err != nil {
			log.Fatal("ERROR PARSING TO NILAI")
		}
		var appen bool = true
		var appIndeks int = 0
		var nmodel NilaiModel
		nmodel.ID = nilai.ID
		nmodel.AlternatifId = nilai.IDAlternatif
		nmodel.KriteriaId = nilai.IDKriteria
		nmodel.Nilai = nilai.Nilai
		if len(nilaiArr) > 0 {
			for index, e := range nilaiArr {
				if e.ID == nilai.IDAlternatif {
					appen = false
					appIndeks = index
				}
			}
			if !appen {
				nilaiArr[appIndeks].Nilai = append(nilaiArr[appIndeks].Nilai, nmodel)
			} else {
				nresp.ID = nilai.IDAlternatif
				nresp.Name = nilai.NamaAlternatif
				nresp.Nilai = append(nresp.Nilai, nmodel)
				nilaiArr = append(nilaiArr, nresp)

			}
		} else {
			nresp.ID = nilai.IDAlternatif
			nresp.Name = nilai.NamaAlternatif
			nresp.Nilai = append(nresp.Nilai, nmodel)
			nilaiArr = append(nilaiArr, nresp)
		}
	}
	//log.Println(nilaiArr)
	defer myConect.Close()
	var nilaiIface []interface{}
	for _, e := range nilaiArr {
		var ifaceHeader interface{}
		var ifaceDetail []interface{}
		for _, enil := range e.Nilai {
			ifaceDetail = append(ifaceDetail, map[string]interface{}{
				"id":    enil.KriteriaId,
				"nilai": enil.Nilai,
			})
		}
		ifaceHeader = map[string]interface{}{
			"id":    e.ID,
			"name":  e.Name,
			"nilai": ifaceDetail,
		}
		nilaiIface = append(nilaiIface, ifaceHeader)
	}
	return nilaiIface
}
func Edit(w http.ResponseWriter, r *http.Request) bool {
	var id string = r.URL.Query().Get("id")
	//var post InputanNilai = paramNilai(r)
	slackerSQL := "DELETE FROM nilai where alternatif_id=?"
	myDB := db.Koneksi()

	slackerSTMT, err := myDB.Prepare(slackerSQL)
	if err != nil {
		return false
	}
	_, err = slackerSTMT.Exec(id)
	if err != nil {
		return false
	}
	defer slackerSTMT.Close()
	defer myDB.Close()
	return Simpan(w, r)
}
func DeleteByAlter(alternatif_id int64) {
	slackerSQL := "DELETE from nilai WHERE alternatif_id=?"
	koneksi := db.Koneksi()
	stmt, _ := koneksi.Prepare(slackerSQL)
	stmt.Exec(alternatif_id)
	defer stmt.Close()
	defer koneksi.Close()
}
func DeleteByCriteria(kriteria_id int64) {
	slackerSQL := "DELETE from nilai WHERE kriteria_id=?"
	koneksi := db.Koneksi()
	stmt, _ := koneksi.Prepare(slackerSQL)
	stmt.Exec(kriteria_id)
	defer stmt.Close()
	defer koneksi.Close()
}

func DeleteByID(id int64) {
	slackerSQL := "DELETE from nilai WHERE id=?"
	koneksi := db.Koneksi()
	stmt, _ := koneksi.Prepare(slackerSQL)
	stmt.Exec(id)
	defer stmt.Close()
	defer koneksi.Close()
}

func ParamNilai(r *http.Request) InputanNilai {
	slackerValue := helper.GetParam(r)
	var iNilai InputanNilai

	slackerMarshall, err := json.Marshal(slackerValue)
	if err != nil {
		log.Println("ERROR CONVERTING JSON")
		return iNilai
	}
	json.Unmarshal(slackerMarshall, &iNilai)
	return iNilai
}

func Simpan(w http.ResponseWriter, r *http.Request) bool {
	var saving bool = true
	slackerSQL := "INSERT INTO nilai(alternatif_id,kriteria_id,nilai)VALUES(?,?,?)"
	mydb := db.Koneksi()
	statement, err := mydb.Prepare(slackerSQL)

	if err != nil {
		return false
	}
	post := ParamNilai(r)
	for _, e := range post.Nilai {
		_, err = statement.Exec(post.Alternatif_Id, e.Kriteria_Id, e.Nilai)
		if err != nil {
			return false
		}
	}
	defer statement.Close()
	defer mydb.Close()
	return saving
}
