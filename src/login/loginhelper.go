package login

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/denil/promethee/src/config/db"
	"golang.org/x/crypto/bcrypt"
)

var Mydb *sql.DB

func initDB() {
	Mydb = db.Koneksi()
}
func Attempt(email string, password string) bool {
	initDB()
	query := "SELECT id,email,password,token,name FROM users WHERE email = ?"

	stmt, err := Mydb.Prepare(query)
	if err != nil {
		log.Println("SQL ERROR", email)
		return false
	}
	result, err := stmt.Query(email)
	if err != nil {
		log.Println("QUERY ERROR")
		return false
	}
	var user UserLogin
	for result.Next() {
		err = result.Scan(&user.ID, &user.Email, &user.Password, &user.Token, &user.Name)
		if err != nil {
			log.Println("NO RESULT")
			return false
		}

	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		UpdateToken(user)
		defer Mydb.Close()
		return true
	} else {
		return false
	}
}

func hashToken(ulog UserLogin) string {
	var hash bool = false
	var hashed string = ""
	myConnect := db.Koneksi()
	sqlku := "SELECT id FROM users WHERE token=?"
	res, err := myConnect.Prepare(sqlku)
	if err != nil {
		log.Fatal("ERR SQL")
	}

	for i := 0; !hash; i++ {

		start := time.Now()
		token := ulog.Email + "" + string(rune(ulog.ID)) + string(rune(start.Year())) + string(rune(start.Month())) + string(rune(start.Day())) + string(rune(start.Hour())) + string(rune(start.Minute())) + string(rune(start.Second()))
		byteToken := []byte(token)
		byteHashed, err := bcrypt.GenerateFromPassword(byteToken, bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("CANNOT ENCRYPT")
		}
		hashed = string(byteHashed) + ""

		type defx struct {
			ID int
		}
		var defeks defx
		defeks.ID = 0
		result, err := res.Query(hashed)
		if err != nil {
			log.Println("")
		}
		for result.Next() {
			err = result.Scan(&defeks.ID)
			if err != nil {
				log.Println("E")
			}
		}
		if defeks.ID == 0 {
			hash = true
		}
	}

	defer myConnect.Close()
	return hashed
}

func UpdateToken(ulog UserLogin) {
	myConnect := db.Koneksi()
	hashed := hashToken(ulog)
	sqlku := "UPDATE users SET token=? WHERE id=?"
	stmt, err := myConnect.Prepare(sqlku)
	if err != nil {
		log.Fatal("FAILED QUERY")
	}
	var uelog UserLogin = ulog
	stmt.Query(hashed, ulog.ID)
	defer myConnect.Close()
	uelog.Token = hashed
	userActive.Email = uelog.Email
	userActive.ID = uelog.ID
	userActive.Token = uelog.Token
	userActive.Name = uelog.Name
}

func LoginWithToken(w http.ResponseWriter, req *http.Request) bool {
	var loginState bool = false
	conn := db.Koneksi()
	slackerSQL := "SELECT id,name,email,token FROM users WHERE token=?"
	stmt, _ := conn.Prepare(slackerSQL)
	var token string = w.Header().Get("Authorization")
	result, _ := stmt.Query(token)
	for result.Next() {
		loginState = true
		result.Scan(&userActive.ID, &userActive.Name, &userActive.Email, &userActive.Token)
	}
	defer conn.Close()
	return loginState
}
