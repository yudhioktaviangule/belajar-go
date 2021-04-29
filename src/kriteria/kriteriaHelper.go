package kriteria

import (
	"log"

	"github.com/denil/promethee/src/config/db"
)

func CastKriteriaToInterface(krit []Kriteria) []interface{} {
	var iface []interface{}
	for _, e := range krit {
		iface = append(iface, map[string]interface{}{
			"id":    e.ID,
			"name":  e.Name,
			"bobot": e.Bobot,
		})
	}
	return iface
}

func GetList() []interface{} {
	var Krit []Kriteria
	var iface []interface{}
	slackerSQL := "SELECT id,name,bobot FROM kriteria LIMIT 0,10"
	dbcon := db.Koneksi()
	stmt, _ := dbcon.Prepare(slackerSQL)
	result, _ := stmt.Query()
	for result.Next() {
		var c Kriteria
		err := result.Scan(&c.ID, &c.Name, &c.Bobot)
		if err != nil {
			return iface
		}
		Krit = append(Krit, c)
	}
	stmt.Close()
	defer dbcon.Close()
	iface = CastKriteriaToInterface(Krit)
	return iface
}

func GetSearchedList(queryString string) []interface{} {
	var Krit []Kriteria
	var iface []interface{}
	slackerSQL := "SELECT id,name,bobot FROM kriteria WHERE name LIKE ? OR bobot LIKE ? LIMIT 0,10"
	dbcon := db.Koneksi()
	stmt, _ := dbcon.Prepare(slackerSQL)
	result, _ := stmt.Query("%"+queryString+"%", "%"+queryString+"%")
	for result.Next() {
		var c Kriteria
		err := result.Scan(&c.ID, &c.Name, &c.Bobot)
		if err != nil {
			return iface
		}
		Krit = append(Krit, c)
	}
	defer stmt.Close()
	defer dbcon.Close()
	iface = CastKriteriaToInterface(Krit)
	return iface
}

func Simpan(kriteria PostKriteria) bool {
	slackerSQL := "INSERT INTO kriteria(name,bobot)VALUES(?,?)"
	koneksi := db.Koneksi()
	stmt, _ := koneksi.Prepare(slackerSQL)
	stmt.Exec(kriteria.Name, kriteria.Bobot)
	stmt.Close()
	defer koneksi.Close()
	return true
}

func edit(data Kriteria) {
	slackerSQL := "UPDATE kriteria SET name=?,bobot=? WHERE id=?"

	koneksi := db.Koneksi()
	stmt, _ := koneksi.Prepare(slackerSQL)
	stmt.Exec(data.Name, data.Bobot, data.ID)
	defer stmt.Close()
	defer koneksi.Close()
}

func IsCriteria(id int64) bool {

	slackerSQL := "SELECT id FROM kriteria WHERE id=?"
	koneksi := db.Koneksi()
	stmt, _ := koneksi.Prepare(slackerSQL)
	r, _ := stmt.Query(id)
	var isCriteria bool = false
	for r.Next() {
		isCriteria = true
	}
	defer stmt.Close()
	defer koneksi.Close()
	return isCriteria
}

func getCritByID(id int64) Kriteria {
	var kriteria Kriteria
	mysql := db.Koneksi()

	slackerSTMT, _ := mysql.Prepare("SELECT id,name,bobot FROM kriteria WHERE id=? LIMIT 1")
	result, err := slackerSTMT.Query(id)
	if err != nil {
		log.Println("NF Data")
	}
	for result.Next() {
		result.Scan(&kriteria.ID, &kriteria.Name, &kriteria.Bobot)
	}
	defer slackerSTMT.Close()
	defer mysql.Close()
	return kriteria
}

func Hapus(id int64) {
	mysql := db.Koneksi()
	slackerSTMT, _ := mysql.Prepare("DELETE FROM kriteria WHERE id=?")
	slackerSTMT.Exec(id)
	slackerNilaiSTMT, _ := mysql.Prepare("DELETE FROM nilai WHERE kriteria_id=?")
	slackerNilaiSTMT.Exec(id)
	defer slackerSTMT.Close()
	defer slackerNilaiSTMT.Close()
	defer mysql.Close()
}
