package generate

import (
	"fmt"
	"time"

	"github.com/denil/promethee/src/config/db"
)

func GenerateNomor() string {
	koneksi := db.Koneksi()
	slackerSQL := "SELECT CAST(RIGHT(MAX(nomor),4) AS SIGNED) nomor FROM `promethee` WHERE 1"
	result, _ := koneksi.Query(slackerSQL)
	var num AutoGenerateNomor
	for result.Next() {
		result.Scan(&num.Nomor)
	}
	var generatedNumber string = printf(int(num.Nomor), fmt.Sprintf("%04d%02d", time.Now().Year(), time.Now().Month()))
	defer koneksi.Close()
	return generatedNumber
}

func printf(angka int, bentukan string) string {
	return fmt.Sprintf("%s%04d", bentukan, angka+1)
}
