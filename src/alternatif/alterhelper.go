package alternatif

import "log"

func GetAlternative() []Alternatif {
	var alternatifDatas []Alternatif
	var sqlQuery string = "SELECT id,name FROM alternatif LIMIT 0,10"
	stmt, err := DBKU.Prepare(sqlQuery)
	var alternatifData Alternatif
	if err != nil {
		alternatifData.ID = 0
		alternatifData.Name = "NOuser"
		alternatifDatas = append(alternatifDatas, alternatifData)
		return alternatifDatas
	}
	result, err := stmt.Query()
	if err != nil {

		alternatifData.ID = 0
		alternatifData.Name = "NOuser"
		alternatifDatas = append(alternatifDatas, alternatifData)
		return alternatifDatas
	}

	for result.Next() {
		err = result.Scan(&alternatifData.ID, &alternatifData.Name)
		if err != nil {
			alternatifData.ID = 0
			alternatifData.Name = "NOuser"
			alternatifDatas = append(alternatifDatas, alternatifData)
			return alternatifDatas
		}
		alternatifDatas = append(alternatifDatas, alternatifData)
	}
	return alternatifDatas
}

func ShowAlter(id int64) Alternatif {
	var alternatif Alternatif
	var sqlStr string = "SELECT id,name FROM alternatif WHERE id=?"
	initDB()
	stmt, err := DBKU.Prepare(sqlStr)

	if err != nil {
		alternatif.Name = "NOT FOUND"
		alternatif.ID = 0
	}
	result, err := stmt.Query(id)
	if err != nil {
		alternatif.Name = "Error in Query"
		alternatif.ID = 0
	}
	for result.Next() {
		err = result.Scan(&alternatif.ID, &alternatif.Name)
		if err != nil {
			alternatif.Name = "Error Getting Result"
			alternatif.ID = 0
		}
	}

	return alternatif

}

func UpdateAlternatif(alter Alternatif) string {
	initDB()
	stmt, err := DBKU.Prepare("update alternatif SET name=? WHERE id=?")
	if err != nil {
		defer stmt.Close()
		defer DBKU.Close()
		return "SQL Error"
	}
	result, err := stmt.Query(alter.Name, alter.ID)
	if err != nil {
		defer stmt.Close()
		defer DBKU.Close()
		return "Failed Updated SQL"
	}
	log.Println(result)
	defer stmt.Close()
	defer DBKU.Close()
	return "Update Success"
}

func simpan(alternatif AlternatifInsert) string {
	sqlQuery := "INSERT INTO alternatif(name)VALUES(?)"
	initDB()
	stmt, err := DBKU.Prepare(sqlQuery)
	if err != nil {

		defer DBKU.Close()
		return "Gagal Simpan data"
	}

	res, err := stmt.Query(alternatif.Name)
	if err != nil {
		return "Gagal Simpan data"
	}
	log.Println(res)
	defer DBKU.Close()
	return "Save Success"
}
